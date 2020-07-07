package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ProcessStatus struct {
	PID          int32   `table:"PID"`
	Name         string  `table:"Name"`
	Username     string  `table:"Username"`
	Exe          string  `table:"Exe"`
	CPU          float64 `table:"CPU"`
	Mem          float32 `table:"Mem"`
	Connections  int     `table:"Connections"`
	OpenFiles    int     `table:"Open Files"`
	MaxOpenFiles int     `table:"Max Open Files"`
}

type ProcessesDisplay []*ProcessStatusDisplay

type ProcessStatusDisplay struct {
	PID          int32  `table:"PID"`
	Name         string `table:"Name"`
	Username     string `table:"Username"`
	Exe          string `table:"Exe"`
	CPU          string `table:"CPU"`
	Mem          string `table:"Mem"`
	Connections  int    `table:"Connections"`
	OpenFiles    int    `table:"Open Files"`
	MaxOpenFiles int    `table:"Max Open Files"`
}

type Limit struct {
	HLimit int // 硬限制
	SLimit int // 软限制
}

// 获取进程最大打开的文件数
func GetProcessMaxOpenFiles(pid int32) (*Limit, error) {
	var (
		err       error
		limitFile = fmt.Sprintf("/proc/%d/limits", pid)
		file      *os.File
		limit     *Limit
	)

	if file, err = os.Open(limitFile); err != nil {
		return nil, err
	}

	defer file.Close()

	space := regexp.MustCompile(`\s+`)
	br := bufio.NewReader(file)
	limit = new(Limit)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		s := string(a)
		if strings.HasPrefix(s, "Max open files") {
			openfiles := space.Split(s, -1)
			slimit, _ := strconv.Atoi(openfiles[3])
			hlimit, _ := strconv.Atoi(openfiles[4])
			limit.SLimit = slimit
			limit.HLimit = hlimit
		}
	}

	return limit, nil
}

// 获取进程一些状态
func GetProcessStatus(pid int32) (*ProcessStatus, error) {
	var (
		err           error
		p             *process.Process
		username      string
		cmd           string
		cpu           float64
		mem           float32
		name          string
		openFile      []process.OpenFilesStat
		processStatus *ProcessStatus
		limit         *Limit
		connections   []net.ConnectionStat
	)

	p, err = process.NewProcess(pid)
	if err != nil {
		return nil, fmt.Errorf("process does not exist, pid: %d", pid)
	}

	processStatus = &ProcessStatus{PID: pid}

	// 获取进程名
	name, err = p.Name()
	if err != nil {
		return nil, fmt.Errorf("get process name error: %s", err)
	}
	processStatus.Name = name

	// 获取进程的用户名
	username, err = p.Username()
	if err != nil {
		return nil, fmt.Errorf("get process username error: %s", err)
	}
	processStatus.Username = username

	// 获取程序的路径
	cmd, err = p.Exe()
	if err != nil {
		if cmd, err = p.Cmdline(); err != nil {
			return nil, fmt.Errorf("get process cmdline error: %s", err)
		}
		processStatus.Exe = strings.Split(cmd, " ")[0]
	} else {
		if len(cmd) > 20 {
			processStatus.Exe = cmd[0:20] + "..."
		} else {
			processStatus.Exe = cmd
		}
	}

	// 获取进程的CPU使用率
	cpu, err = p.CPUPercent()
	if err != nil {
		return nil, fmt.Errorf("get process cpu percent error: %s", err)
	}
	processStatus.CPU = cpu

	// 获取内存使用率
	mem, err = p.MemoryPercent()
	if err != nil {
		return nil, fmt.Errorf("get process mem percent error: %s", err)
	}
	processStatus.Mem = mem

	// 获取进程打开文件数
	openFile, err = p.OpenFiles()
	if err != nil {
		return nil, fmt.Errorf("get process openfiles error: %s", err)
	}
	processStatus.OpenFiles = len(openFile)

	// 获取系统打开文件总数
	limit, err = GetProcessMaxOpenFiles(pid)
	if err != nil {
		return nil, fmt.Errorf("get process MaxOpenFiles error: %s", err)
	}
	processStatus.MaxOpenFiles = limit.SLimit

	// 获取连接数
	if connections, err = p.Connections(); err != nil {
		return nil, fmt.Errorf("get process connections error: %s", err)
	}
	processStatus.Connections = len(connections)

	return processStatus, nil
}

// 获取所有的进程
func GetAllProcess() ([]int32, error) {
	var (
		err      error
		proc     = "/proc"
		fileInfo []os.FileInfo
		pattern  = regexp.MustCompile(`\d+`)
		pids     []int32
	)

	if fileInfo, err = ioutil.ReadDir(proc); err != nil {
		return nil, err
	}

	for _, f := range fileInfo {
		name := f.Name()
		if pattern.MatchString(name) {
			pid, _ := strconv.Atoi(name)
			pids = append(pids, int32(pid))
		}
	}

	return pids, nil
}

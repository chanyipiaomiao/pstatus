package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SystemInfo struct {
	CPU                int     `table:"CPU"`
	CPUPercent         string  `table:"CPU Used Percent"`
	Memory             float64 `table:"Memory"`
	MemoryUsedPercent  string  `table:"Memory Used Percent"`
	ProcessCount       int     `table:"Process Number"`
	SystemOpenFiles    int     `table:"System Open Files"`
	MaxSystemOpenFiles int     `table:"Max System Open Files"`
}

func GetSystemInfo() ([]*SystemInfo, error) {
	var (
		err            error
		max            = "/proc/sys/fs/file-max"
		used           = "/proc/sys/fs/file-nr"
		fileMax        []byte
		fileUsed       []byte
		maxNum         int
		usedNum        int
		usedArr        []string
		usedRex        = regexp.MustCompile(`\s+`)
		cpuCount       int
		memInfo        *mem.VirtualMemoryStat
		memTotal       float64
		memUsedPercent string
		processNum     []*process.Process
		cpuUsedPercent []float64
	)

	if fileMax, err = ioutil.ReadFile(max); err != nil {
		return nil, err
	}

	if maxNum, err = strconv.Atoi(strings.Trim(string(fileMax), "\n")); err != nil {
		return nil, err
	}

	if fileUsed, err = ioutil.ReadFile(used); err != nil {
		return nil, err
	}

	usedArr = usedRex.Split(string(fileUsed), -1)
	if usedNum, err = strconv.Atoi(usedArr[0]); err != nil {
		return nil, err
	}

	// 获取CPU核数
	if cpuCount, err = cpu.Counts(true); err != nil {
		return nil, err
	}

	// 获取CPU使用率
	if cpuUsedPercent, err = cpu.Percent(1*time.Second, false); err != nil {
		return nil, err
	}

	// 获取内存
	if memInfo, err = mem.VirtualMemory(); err != nil {
		return nil, err
	}

	memTotal = math.Round(float64(memInfo.Total) / 1024 / 1024 / 1024)
	memUsedPercent = fmt.Sprintf("%.2f%%", memInfo.UsedPercent)

	// 获取进程总数量
	if processNum, err = process.Processes(); err != nil {
		return nil, err
	}

	return []*SystemInfo{
		{
			CPU:                cpuCount,
			CPUPercent:         fmt.Sprintf("%.2f%%", cpuUsedPercent[0]),
			Memory:             memTotal,
			MemoryUsedPercent:  memUsedPercent,
			ProcessCount:       len(processNum),
			SystemOpenFiles:    usedNum,
			MaxSystemOpenFiles: maxNum,
		},
	}, nil
}

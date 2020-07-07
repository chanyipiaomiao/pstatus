package main

import (
	"flag"
	"fmt"
	"github.com/modood/table"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var (
		err            error
		pid            *string
		top            *int
		sorted         *string
		pids           []int32
		processes      Processes
		systemInfo     []*SystemInfo
		processDisplay ProcessesDisplay
	)

	pid = flag.String("p", "", "process pid, like: 1234 or 123,456")
	top = flag.Int("t", 0, "display top x")
	sorted = flag.String("s", "openfile", "sort field, value: openfile|cpu|mem|conn")
	flag.Parse()

	if *pid == "" {
		pids, err = GetAllProcess()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		for _, p := range strings.Split(*pid, ",") {
			pi, _ := strconv.Atoi(p)
			pids = append(pids, int32(pi))
		}
	}

	if systemInfo, err = GetSystemInfo(); err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range pids {
		status, err := GetProcessStatus(p)
		if err != nil {
			continue
		}
		if status.OpenFiles > 1 {
			processes = append(processes, status)
		}
	}
	switch *sorted {
	case "openfile":
		sort.Sort(processes)
	case "cpu":
		sort.Sort(ProcessSortByCPU{processes})
	case "mem":
		sort.Sort(ProcessSortByMem{processes})
	case "conn":
		sort.Sort(ProcessSortByConnections{processes})
	}

	for _, p := range processes {
		processDisplay = append(processDisplay, &ProcessStatusDisplay{
			PID:          p.PID,
			Name:         p.Name,
			Username:     p.Username,
			Exe:          p.Exe,
			CPU:          fmt.Sprintf("%.2f%%", p.CPU),
			Mem:          fmt.Sprintf("%.2f%%", p.Mem),
			Connections:  p.Connections,
			OpenFiles:    p.OpenFiles,
			MaxOpenFiles: p.MaxOpenFiles,
		})
	}

	table.Output(systemInfo)

	if *top == 0 {
		table.Output(processDisplay)
	} else {
		length := len(processDisplay)
		if length > *top {
			table.Output(processDisplay[0:*top])
		} else {
			table.Output(processDisplay)
		}
	}
}

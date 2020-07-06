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
		pids           []int32
		processes      Processes
		systemOpenFile []*SystemOpenFiles
	)

	pid = flag.String("p", "", "process pid, like: 1234 or 123,456")
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

	if systemOpenFile, err = GetSystemOpenFiles(); err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range pids {
		status, err := GetProcessStatus(p)
		if err != nil {
			fmt.Println(err)
			continue
		}
		processes = append(processes, status)
	}

	sort.Sort(processes)

	table.Output(systemOpenFile)
	table.Output(processes)

}

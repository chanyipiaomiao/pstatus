package main

import (
	"flag"
	"fmt"
	"github.com/modood/table"
	"strconv"
	"strings"
)

func main() {

	var (
		err            error
		pid            *string
		pids           []int32
		processes      []*ProcessStatus
		systemOpenFile []*SystemOpenFiles
	)

	pid = flag.String("p", "", "process pid, like: 1234 or 123,456")
	flag.Parse()

	if *pid == "" {
		fmt.Println("need process pid, like: 1234 or 123,456")
		return
	}

	if systemOpenFile, err = GetSystemOpenFiles(); err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range strings.Split(*pid, ",") {
		pi, _ := strconv.Atoi(p)
		pids = append(pids, int32(pi))
	}

	for _, p := range pids {
		status, err := GetProcessStatus(p)
		if err != nil {
			fmt.Println(err)
			continue
		}
		processes = append(processes, status)
	}

	table.Output(systemOpenFile)
	table.Output(processes)

}

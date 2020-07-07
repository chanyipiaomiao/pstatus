package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"math"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "/data/app/dolphin/monitor/monitor"
	fmt.Printf("%.40s\n", s)

	s1 := strings.Split(s, " ")
	fmt.Println(s1)

	s2 := "/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.242.b07-1.el6_10.x86_64/jre/bin/java"
	fmt.Println(s2[0:50] + "...")
}

func TestGetProcessMaxOpenFiles(t *testing.T) {
	limit, err := GetProcessMaxOpenFiles(359)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(limit)
}

func TestGetSystemOpenFiles(t *testing.T) {
	f, err := GetSystemInfo()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(f)
}

func TestGetAllProcess(t *testing.T) {
	pids, err := GetAllProcess()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(pids)
}

func TestConvertMem(t *testing.T) {
	fmt.Println(float64(8061128) / 1000 / 1000)
	s := math.Round(float64(8061128) / 1000 / 1000)
	fmt.Println(s)
}

func TestGetProcessNum(t *testing.T) {
	p, err := process.Processes()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(p)
}

package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "/data/app/dolphin/monitor/monitor"
	fmt.Printf("%.40s\n", s)

	s1 := strings.Split(s, " ")
	fmt.Println(s1)
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
	f, err := GetSystemOpenFiles()
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

package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type SystemOpenFiles struct {
	SystemOpenFiles    int `table:"System Open Files"`
	MaxSystemOpenFiles int `table:"Max System Open Files"`
}

func GetSystemOpenFiles() ([]*SystemOpenFiles, error) {
	var (
		err      error
		max      = "/proc/sys/fs/file-max"
		used     = "/proc/sys/fs/file-nr"
		fileMax  []byte
		fileUsed []byte
		maxNum   int
		usedNum  int
		usedArr  []string
		usedRex  = regexp.MustCompile(`\s+`)
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

	return []*SystemOpenFiles{
		{
			SystemOpenFiles:    usedNum,
			MaxSystemOpenFiles: maxNum,
		},
	}, nil
}

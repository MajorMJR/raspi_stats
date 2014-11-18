package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Mounts struct {
	Mount []Mount
}

type Mount struct {
	Device string
	Mount  string
	Fs     string
}

func getMounts(path string) (*Mounts, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	mounts := Mounts{}
	mount := &Mount{}

	content := string(b)
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		if i == len(lines)-1 {
			continue
		}
		mountLine := strings.Fields(line)
		mount.Device = mountLine[0]
		mount.Mount = mountLine[1]
		mount.Fs = mountLine[2]
		mounts.Mount = append(mounts.Mount, *mount)
	}

	return &mounts, nil
}

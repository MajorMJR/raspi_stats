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

		fmt.Println(i, len(line), len(lines))
		mountLine := strings.Fields(line)
		fmt.Println(mountLine[0], mountLine[1], mountLine[2])
		mount.Device = mountLine[0]
		mount.Mount = mountLine[1]
		mount.Fs = mountLine[2]
		mounts.Mount = append(mounts.Mount, *mount)
	}
	fmt.Println(mounts)
	return &mounts, nil
}

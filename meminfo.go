package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Mem struct {
	Total     float64
	Free      float64
	Available float64
	Cached    float64
}

var meminfoRegExp = regexp.MustCompile("([^:]*?)\\s*:\\s*(.*)$")

func getMemInfo(path string) (*Mem, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	memory := &Mem{}

	content := string(b)
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		if i == len(lines)-1 {
			continue
		}

		submatches := meminfoRegExp.FindStringSubmatch(line)
		key := submatches[1]
		value := submatches[2]

		// If value has "kB" then trim it
		if value == "0" {
			value = value
		} else {
			value = value[:len(value)-3]
		}

		switch key {
		case "MemTotal":
			memory.Total, _ = strconv.ParseFloat(value, 64)
			memory.Total = memory.Total / 1000
		case "MemFree":
			memory.Free, _ = strconv.ParseFloat(value, 64)
			memory.Free = memory.Free / 1000
		case "MemAvailable":
			memory.Available, _ = strconv.ParseFloat(value, 64)
			memory.Available = memory.Available / 1000
		case "Cached":
			memory.Cached, _ = strconv.ParseFloat(value, 64)
			memory.Cached = memory.Cached / 1000
		}
	}
	return memory, nil
}

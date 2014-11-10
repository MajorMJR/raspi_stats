// Adapted from

package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type CPUinfo struct {
	Processors []Processor `json:"processors"`
}

type Processor struct {
	Id        int64    `json:"id"`
	VendorId  string   `json:"vendor_id"`
	Model     int64    `json:"model"`
	ModelName string   `json:"model_name"`
	Flags     []string `json:"flags"`
	Cores     int64    `json:"cores"`
	MHz       float64  `json:"mhz"`
}

type CPUload struct {
	OneMin     float64
	FiveMin    float64
	FifteenMin float64
}

var cpuinfoRegExp = regexp.MustCompile("([^:]*?)\\s*:\\s*(.*)$")

func getCPUInfo(path string) (*CPUinfo, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cpuinfo := CPUinfo{}
	processor := &Processor{}

	content := string(b)
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		var key string
		var value string

		if len(line) == 0 && i != len(lines)-1 {
			// end of a processor
			cpuinfo.Processors = append(cpuinfo.Processors, *processor)
			processor = &Processor{}
			continue
		} else if i == len(lines)-1 {
			continue
		}

		submatches := cpuinfoRegExp.FindStringSubmatch(line)
		key = submatches[1]
		value = submatches[2]

		switch key {
		case "processor":
			processor.Id, _ = strconv.ParseInt(value, 10, 32)
		case "vendor_id":
			processor.VendorId = value
		case "model":
			processor.Model, _ = strconv.ParseInt(value, 10, 32)
		case "model name":
			if value == "ARMv6-compatible processor rev 7 (v6l)" {
				processor.MHz, _ = getCPUMHzoRaspi()
			}
			processor.ModelName = value
		case "flags":
			processor.Flags = strings.Fields(value)
		case "cpu cores":
			processor.Cores, _ = strconv.ParseInt(value, 10, 32)
		case "cpu MHz":
			processor.MHz, _ = strconv.ParseFloat(value, 64)
		}
	}

	return &cpuinfo, nil
}

func getCPUMHzRaspi() (float64, error) {
	raspi_mhz, err := ioutil.ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq")
	if err != nil {
		return 0, err
	}
	mHz_string := string(raspi_mhz)
	mHz, err := strconv.ParseFloat(mHz_string[:len(mHz_string)-1], 64)
	if err != nil {
		return 0, err
	}

	return mHz, nil
}

func getCpuLoad(path string) (*CPUload, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	load := &CPUload{}
	content := string(b)
	lines := strings.Split(content, " ")

	load.OneMin, _ = strconv.ParseFloat(lines[0], 64)
	load.FiveMin, _ = strconv.ParseFloat(lines[1], 64)
	load.FifteenMin, _ = strconv.ParseFloat(lines[2], 64)

	return load, nil
}

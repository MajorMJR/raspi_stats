package main

import (
	"io/ioutil"
	"strconv"
)

type System struct {
	Hostname string
	Ip_addr  string
	Temp     float64
}

func getTemp() (float64, error) {
	b, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}

	tempStr := string(b)[:5]

	temp, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		return 0, err
	}
	temp = temp / 1000
	return temp, nil
}

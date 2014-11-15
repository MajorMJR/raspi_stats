package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

type System struct {
	Hostname string
	Ip_addr  string
	Temp     float64
	OS       string
}

func getSysinfo() (*System, error) {
	hostname, _ := os.Hostname()
	temp, _ := getTemp()
	system := &System{Hostname: hostname, Temp: temp}

	return system, nil
}

func getOS() {

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

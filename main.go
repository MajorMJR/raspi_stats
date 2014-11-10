package main

import (
	"net/http"
	"os"
	"text/template"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	temp, _ := getTemp()
	system := &System{Hostname: hostname, Temp: temp}
	cpu, _ := getCPUInfo("/proc/cpuinfo")
	cpuload, _ := getCpuLoad("/proc/loadavg")

	data := struct {
		System  *System
		CPU     *CPUinfo
		CPUload *CPUload
		Env     string
	}{
		system,
		cpu,
		cpuload,
		os.Getenv("GOPATH"),
	}
	//fmt.Println(getCPUInfo("/proc/cpuinfo"))
	t, _ := template.ParseFiles("tmpl/index.html")
	t.Execute(w, data)
}

func main() {

	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)

}

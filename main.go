package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	system := &System{Hostname: hostname}
	cpu, _ := readCPUInfo("/proc/cpuinfo")
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
	fmt.Println(getCpuLoad("/proc/loadavg"))
	t, _ := template.ParseFiles("tmpl/index.html")
	t.Execute(w, data)
}

func main() {

	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)

}

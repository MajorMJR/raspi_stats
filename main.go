package main

import (
	"net/http"
	"text/template"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	system, _ := getSysinfo()
	cpu, _ := getCPUInfo("/proc/cpuinfo")
	cpuload, _ := getCpuLoad("/proc/loadavg")
	memory, _ := getMemInfo("/proc/meminfo")

	data := struct {
		System  *System
		CPU     *CPUinfo
		CPUload *CPUload
		Memory  *Mem
	}{
		system,
		cpu,
		cpuload,
		memory,
	}

	t, _ := template.ParseFiles("tmpl/index.html")
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)

}

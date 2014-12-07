package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	system, _ := getSysinfo()
	cpu, _ := getCPUInfo("/proc/cpuinfo")
	cpuload, _ := getCpuLoad("/proc/loadavg")
	memory, _ := getMemInfo("/proc/meminfo")
	mounts, _ := getMounts("/proc/mounts")

	data := struct {
		System  *System
		CPU     *CPUinfo
		CPUload *CPUload
		Memory  *Mem
		Mounts  *Mounts
	}{
		system,
		cpu,
		cpuload,
		memory,
		mounts,
	}

	t, _ := template.ParseFiles("tmpl/index.html")
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8888", nil)
	fmt.Println("Server running on :8080")

}

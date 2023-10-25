package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/procfs"
)

var (
	ProcFS, _ = procfs.NewDefaultFS()
	memFree   = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "mem_free",
			Help: "Free memory",
		},
	)
	proc_count = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "proc_count",
			Help: "Total count of running processes",
		},
	)
)

func init() {
	prometheus.MustRegister(memFree)
	prometheus.MustRegister(proc_count)
}

func main() {
	go collectMemInfo()
	go collectProcStat()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start HTTP server: %s\n", err.Error())
		os.Exit(1)
	}
}

func collectMemInfo() {
	for {
		memStat, err := ProcFS.Meminfo()
		if err != nil {
			fmt.Printf("Failed to retrieve memory stats: %s\n", err.Error())
			continue
		}

		memFree.Set(float64(*memStat.MemFree))

		time.Sleep(10 * time.Second)
	}
}

func collectProcStat() {
	for {
		procStat, err := ProcFS.AllProcs()
		if err != nil {
			fmt.Printf("Failed to retrieve processes stats: %s\n", err.Error())
			continue
		}

		proc_count.Set(float64(len(procStat)))

		time.Sleep(10 * time.Second)
	}

}

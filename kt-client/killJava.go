package main

import (
	"log"
	"os"
	"regexp"
	"sort"

	"github.com/shirou/gopsutil/process"
)

var javaRegexp = regexp.MustCompile(`javaw?(\.exe)?`)

type byMemoryUsage []*process.Process

func (a byMemoryUsage) Len() int { return len(a) }

func (a byMemoryUsage) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a byMemoryUsage) Less(i, j int) bool {
	iMU, err := a[i].MemoryPercent()
	if err != nil {
		log.Panic(err)
	}

	jMU, err := a[j].MemoryPercent()
	if err != nil {
		log.Panic(err)
	}

	return iMU < jMU
}

func killJava() {
	log.Println("=== killJava ===")
	defer log.Println("================")

	pids, err := process.Pids()
	if err != nil {
		log.Panic(err)
	}

	var procs []*process.Process
	for _, pid := range pids {
		proc, err := process.NewProcess(pid)
		if err != nil {
			log.Panic(err)
		}

		name, err := proc.Name()
		if err != nil {
			log.Println(err)
			continue
		}

		if javaRegexp.MatchString(name) {
			procs = append(procs, proc)
		}
	}

	sort.Sort(byMemoryUsage(procs))
	log.Println("Found", len(procs), "Java processes.")
	if len(procs) == 0 {
		return
	}

	log.Println("procs:", procs)
	log.Println("Killing top memory user:", procs[0])
	osProc, err := os.FindProcess(int(procs[0].Pid))
	if err != nil {
		log.Panic(err)
	}
	if err = osProc.Kill(); err != nil {
		log.Panic(err)
	}
}

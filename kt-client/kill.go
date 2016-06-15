package main

import (
	"log"
	"os"
	"regexp"

	"github.com/shirou/gopsutil/process"
)

func kill(m map[string]string) error {
	re, err := regexp.Compile(m["regex"])
	if err != nil {
		return err
	}

	pids, err := process.Pids()
	if err != nil {
		return err
	}

	var procs []int
	for _, pid := range pids {
		proc, err := process.NewProcess(pid)
		if err != nil {
			return err
		}

		name, err := proc.Name()
		if err != nil {
			log.Println(err)
			continue
		}

		if re.MatchString(name) {
			procs = append(procs, int(proc.Pid))
		}
	}

	log.Println("Found", len(procs), "Java processes.")
	if len(procs) == 0 {
		return nil
	}
	log.Println("procs:", procs)

	for _, pid := range procs {
		proc, err := os.FindProcess(pid)
		if err != nil {
			return err
		}
		if err = proc.Kill(); err != nil {
			return err
		}
	}
	return nil
}

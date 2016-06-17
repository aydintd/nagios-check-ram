package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	mem_total int
	mem_free  int
	mem_used  int
	mem_cache int
	mem_perc  int
	err       error
)

const (
	_version = "1.0"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func memory_usage() (mem_used, mem_perc int) {
	data, err := os.Open("/proc/meminfo")
	check(err)
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		switch parts[0] {
		case "MemTotal:":
			mem_total, err = strconv.Atoi(parts[1])
			check(err)
		case "MemFree:":
			mem_free, err = strconv.Atoi(parts[1])
			check(err)
		case "Cached:":
			mem_cache, err = strconv.Atoi(parts[1])
			check(err)
		}

	}

	mem_used = mem_total - (mem_free + mem_cache)
	mem_perc = mem_used * 100 / mem_total

	return mem_used, mem_perc
}

func help() {
	fmt.Printf("%s v%s\n", os.Args[0], _version)
	fmt.Println()
	fmt.Printf("Usage : ./%s -w %%WARNING -c %%CRITICAL\n", os.Args[0])
	fmt.Println()
	fmt.Println("WARNING and CRITICAL values are percentage values without %")
	fmt.Println()
	fmt.Println("2016 - Aydin Doyak <aydintd@gmail.com>")
	os.Exit(5)
}

func main() {
	mem_used, percentage := memory_usage()
	var warn, crit int

	if len(os.Args) != 5 {
		help()
		os.Exit(5)
	}

	if os.Args[1] == "-w" && os.Args[3] == "-c" {
		warn, err = strconv.Atoi(os.Args[2])
		check(err)
		crit, err = strconv.Atoi(os.Args[4])
		check(err)
		if warn >= crit {
			fmt.Println("WARNING value can not be bigger than CRITICAL value")
			os.Exit(5)
		} else if crit > 100 {
			fmt.Println("%CRITICAL value can not be bigger than %100")
			os.Exit(5)
		}
	} else {
		help()
	}

	switch {
	case 0 <= percentage && percentage < warn:
		fmt.Printf("Memory: OK - Total: %d KB - Used: %d KB - Memory Usage: %%%d|TOTAL=%d;;;; USED=%d;;;; FREE=%d;;;; CACHE=%d;;;; \n", mem_total, mem_free, percentage, mem_total, mem_used, mem_free, mem_cache)
		os.Exit(0)
	case warn <= percentage && percentage < crit:
		fmt.Printf("Memory: WARNING - Total: %d KB - Used: %d KB - Memory Usage: %%%d|TOTAL=%d;;;; USED=%d;;;; FREE=%d;;;; CACHE=%d;;;; \n", mem_total, mem_free, percentage, mem_total, mem_used, mem_free, mem_cache)
		os.Exit(1)
	case crit <= percentage:
		fmt.Printf("Memory: CRITICAL - Total: %d KB - Used: %d KB - Memory Usage: %%%d|TOTAL=%d;;;; USED=%d;;;; FREE=%d;;;; CACHE=%d;;;; \n", mem_total, mem_free, percentage, mem_total, mem_used, mem_free, mem_cache)
		os.Exit(2)
	default:
		fmt.Printf("UNKNOWN value.\n")
		os.Exit(3)
	}
}

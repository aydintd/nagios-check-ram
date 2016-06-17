package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	mem_total int
	mem_free  int
	mem_used  int
	mem_cache int
	mem_perc  int
)

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

func main() {
	_, percentage := memory_usage()
	warn, err := strconv.Atoi(os.Args[1])
	check(err)
	crit, err := strconv.Atoi(os.Args[2])
	check(err)

	switch {
	case 0 <= percentage && percentage < warn:
		fmt.Printf("OK - %%%d memory usage.\n", percentage)
		os.Exit(0)
	case warn <= percentage && percentage < crit:
		fmt.Printf("WARNING - %%%d memory usage.\n", percentage)
		os.Exit(1)
	case crit <= percentage:
		fmt.Printf("CRITICAL - %%%d memory usage.\n", percentage)
		os.Exit(2)
	default:
		fmt.Printf("UNKNOWN value.\n", percentage)
		os.Exit(3)
	}
}

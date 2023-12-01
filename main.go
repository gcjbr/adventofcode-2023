package main

import (
	"flag"

	"github.com/gcjbr/adventofcode-2023/day01"
)

func main() {
	day := flag.Int("day", 1, "Day to run")
	part := flag.Int("part", 1, "Part to run")

	flag.Parse()

	switch *day {
	case 1:
		switch *part {
		case 1:
			day01.Part1()
		default:
			println("Invalid part")
		}

	default:
		println("Invalid day")
	}

}

package main

import (
	"flag"
	"fmt"

	"github.com/gcjbr/adventofcode-2023/day01"
)

type PartFunction func()

func main() {
	day := flag.Int("day", 1, "Day to run")
	part := flag.Int("part", 1, "Part to run")

	flag.Parse()

	functionMap := map[string]PartFunction{
		"1_1": day01.Part1,
	}

	key := fmt.Sprintf("%d_%d", *day, *part)

	if fn, ok := functionMap[key]; ok {
		fn()
	} else {
		fmt.Println("Invalid day or part")
	}

}

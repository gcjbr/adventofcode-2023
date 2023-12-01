package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Part1() {
	f, err := os.Open("day01/day01-01.txt")

	var calibrations []int

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		calibrations = append(calibrations, computeCalibration(scanner.Text()))
	}

	println("Result: ", sumSlice(calibrations))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeCalibration(line string) int {
	var firstNumber string
	var lastNumber string

	// Find first number
	for i := 0; i < len(line); i++ {

		if _, err := strconv.Atoi(string(line[i])); err == nil {

			firstNumber = string(line[i])

			break
		}
	}

	// Find last number
	for i := len(line) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(line[i])); err == nil {

			lastNumber = string(line[i])

			break
		}
	}

	result, _ := strconv.Atoi(firstNumber + lastNumber)

	return result
}

func sumSlice(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

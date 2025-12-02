package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"aoc2025/internal/days"
)

func main() {
	day := flag.Int("day", 0, "Day to run (1-25)")
	part := flag.Int("part", 0, "Part to run (1 or 2, 0 for both)")
	flag.Parse()

	if *day < 1 || *day > 25 {
		fmt.Println("Please specify a day between 1 and 25")
		os.Exit(1)
	}

	solver := days.GetSolver(*day)
	if solver == nil {
		fmt.Printf("Day %d not implemented yet\n", *day)
		os.Exit(1)
	}

	input, err := os.ReadFile(fmt.Sprintf("inputs/day%02d.txt", *day))
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	if *part == 1 || *part == 0 {
		start := time.Now()
		result := solver.Part1(string(input))
		duration := time.Since(start)
		fmt.Printf("Day %d Part 1: %v (took %v)\n", *day, result, duration)
	}

	if *part == 2 || *part == 0 {
		start := time.Now()
		result := solver.Part2(string(input))
		duration := time.Since(start)
		fmt.Printf("Day %d Part 2: %v (took %v)\n", *day, result, duration)
	}
}

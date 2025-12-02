package days

import (
	"aoc2025/internal/util"
	"fmt"
	"strconv"
	"strings"
)

type Day01 struct{}

func (d *Day01) Part1(input string) any {
	// 	test_input := `L68
	// L30
	// R48
	// L5
	// R60
	// L55
	// L1
	// L99
	// R14
	// L82
	// `
	position := 50
	answer_count := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		operation := line[0]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}
		switch operation {
		case 'L':
			{
				position = util.WrapValue((position - num), 100)
			}
		case 'R':
			{
				position = util.WrapValue((position + num), 100)
			}
		default:
			{
				fmt.Printf("found bad operation: %d", operation)
				continue
			}
		}
		if position == 0 {
			answer_count++
		}
	}
	return answer_count
}

func (d *Day01) Part2(input string) any {
	position := 50
	zeroCount := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		operation := line[0]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}

		for range num {
			switch operation {
			case 'L':
				position--
			case 'R':
				position++
			}

			if position < 0 {
				position = 99
			}
			if position > 99 {
				position = 0
			}

			if position == 0 {
				zeroCount++
			}
		}
	}

	return zeroCount
}

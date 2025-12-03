package days

import (
	"strings"
)

type Day03 struct{}

func (d *Day03) Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	battery_sum := 0
	for _, line := range lines {
		// go has no builtins for just converting slices???
		numbers := make([]int, len(line))
		for i, ch := range line {
			// we're promised by the input this is all valid !!
			numbers[i] = int(ch - '0')
		}

		battery := 0
		// bruteforce
		for i := range numbers {
			for j := i + 1; j < len(numbers); j++ {
				pair := (numbers[i] * 10) + numbers[j]
				if pair > battery {
					battery = pair
				}
			}
		}

		battery_sum += battery
	}

	return battery_sum
}

func (d *Day03) Part2(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	battery_sum := 0
	for _, line := range lines {
		// go has no builtins for just converting slices???
		numbers := make([]int, len(line))
		for i, ch := range line {
			// we're promised by the input this is all valid !!
			numbers[i] = int(ch - '0')
		}

		toRemove := len(numbers) - 12
		result := []int{}
		// greed
		for i := range numbers {
			for len(result) > 0 &&
				result[len(result)-1] < numbers[i] &&
				toRemove > 0 {
				result = result[:len(result)-1]
				toRemove--
			}
			result = append(result, numbers[i])
		}
		result = result[:12]
		sum := 0
		for _, n := range result {
			sum = sum*10 + n
		}

		battery_sum += sum
	}

	return battery_sum
}

package days

import (
	"sort"
	"strconv"
	"strings"
)

type Day05 struct{}

func (d *Day05) Part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	ranges := [][2]int{}
	products := []int{}

outer:
	for _, v := range lines {
		for _, ch := range v {
			if ch == '-' {
				range_str := strings.Split(v, "-")
				// we're promised by the input this is all valid !!
				lr, _ := strconv.Atoi(range_str[0])
				rr, _ := strconv.Atoi(range_str[1])
				sl := [2]int{lr, rr}
				ranges = append(ranges, sl)
				continue outer
			}
		}
		if v == "" {
			continue
		}
		product, _ := strconv.Atoi(v)
		products = append(products, product)
	}

	sum := 0
	for _, p := range products {
		for _, r := range ranges {
			if p >= r[0] && p <= r[1] {
				sum += 1
				break
			}
		}
	}

	return sum
}

func (d *Day05) Part2(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ranges := [][2]int{}
	for _, v := range lines {
		for _, ch := range v {
			if ch == '-' {
				range_str := strings.Split(v, "-")
				// we're promised by the input this is all valid !!
				lr, _ := strconv.Atoi(range_str[0])
				rr, _ := strconv.Atoi(range_str[1])
				sl := [2]int{lr, rr}
				ranges = append(ranges, sl)
				break
			}
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := [][2]int{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		curr := ranges[i]
		if curr[0] <= last[1]+1 {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			merged = append(merged, curr)
		}
	}

	sum := 0
	for _, r := range merged {
		sum += r[1] - r[0] + 1
	}

	return sum
}

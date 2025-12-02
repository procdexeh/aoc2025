package days

import (
	"strconv"
	"strings"
)

type Day02 struct{}

func (d *Day02) Part1(input string) any {
	id_list := strings.Split(strings.TrimSpace(input), ",")
	bad_id_count := 0
	for _, v := range id_list {
		id_pair := strings.Split(v, "-")
		left_num, err := strconv.Atoi(id_pair[0])
		if err != nil {
			return err
		}
		right_num, err := strconv.Atoi(id_pair[1])
		if err != nil {
			return err
		}

		for id := left_num; id <= right_num; id++ {
			id_str := strconv.Itoa(id)
			if len(id_str)%2 != 0 {
				continue
			}
			middle := len(id_str) / 2

			if id_str[:middle] == id_str[middle:] {
				bad_id_count += id
			}

		}
	}

	return bad_id_count
}

func (d *Day02) Part2(input string) any {
	id_list := strings.Split(strings.TrimSpace(input), ",")
	bad_id_count := 0
	for _, v := range id_list {
		id_pair := strings.Split(v, "-")
		left_num, err := strconv.Atoi(id_pair[0])
		if err != nil {
			return err
		}
		right_num, err := strconv.Atoi(id_pair[1])
		if err != nil {
			return err
		}

		for id := left_num; id <= right_num; id++ {
			id_str := strconv.Itoa(id)

			// we have to check if there is a repeat at least twice
			if strings.Contains((id_str + id_str)[1:len(id_str+id_str)-1], id_str) {
				bad_id_count += id
			}
		}
	}

	return bad_id_count
}

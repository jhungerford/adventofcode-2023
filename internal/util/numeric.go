package util

import (
	"strconv"
	"strings"
)

// MustAtoi runs strconv.Atoi(s), panicing if the string isn't a number.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

// IntList converts a space-separated list of numbers like '1 2 3' into a list of ints.
func IntList(s string) ([]int, error) {
	strs := strings.Split(s, " ")

	list := make([]int, 0, len(strs))

	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		list = append(list, num)
	}

	return list, nil
}

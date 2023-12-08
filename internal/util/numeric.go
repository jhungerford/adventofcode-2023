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

// GCD returns the greatest common divisor of the two numbers.
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// LCM returns the least common multiple of the given numbers.
func LCM(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	result := nums[0] * nums[1] / GCD(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		result = LCM(result, nums[i])
	}

	return result
}

package main

import "fmt"

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2

		if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
}

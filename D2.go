package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input2.txt")
	if err != nil {
		fmt.Println(err)
	}

	counter := 0
	lines := strings.Split(string(content), "\n")
	for _, line := range(lines) {
		numsStrings := strings.Split(line, " ")
		nums := []int{}
		for i := range(numsStrings) {
			num, err := strconv.Atoi(numsStrings[i])
			if err != nil {
				fmt.Println(err)
			} 
			nums = append(nums, num)
		}

		if isSafe(nums) {
			counter++
		} else {
			safe := false
			for i := range nums {
				numscopy := append([]int{}, nums[:i]...)
				numscopy = append(numscopy, nums[i+1:]...)
				if isSafe(numscopy) {
					safe = true
				}
			}

			if safe {
				counter++
			}
		}
	}

	fmt.Println(counter)
}

func isSafe(nums []int) bool {
	safe := true
	ascending := true
	if nums[0] > nums[1] {
		ascending = false
	}
	if nums[0] == nums[1] {
		safe = false
	}

	for i := 1; i < len(nums); i++ {
		if ascending {
			if nums[i] - nums[i-1] > 3 || nums[i] - nums[i-1] < 1 {
				safe = false
			}
		}
		if !ascending {
			if nums[i-1] - nums[i] > 3 || nums[i-1] - nums[i] < 1{
				safe = false
			}
		}
	}

	return safe
}
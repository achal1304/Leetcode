package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	} else if len(nums) == 1 {
		return []string{fmt.Sprint(nums[0])}
	}
	ans := []string{}

	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if start == i-1 {
				ans = append(ans, fmt.Sprint(nums[i-1]))
				start = i
			} else {
				ans = append(ans, fmt.Sprint(nums[start])+"->"+fmt.Sprint(nums[i-1]))
				start = i
			}
		}
	}

	if start == len(nums)-1 {
		ans = append(ans, fmt.Sprint(nums[len(nums)-1]))
	} else {
		ans = append(ans, fmt.Sprint(nums[start])+"->"+fmt.Sprint(nums[len(nums)-1]))
	}

	return ans
}

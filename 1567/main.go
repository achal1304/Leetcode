package main

import "fmt"

// REVISIT
func main() {
	fmt.Println(getMaxLen([]int{-1, -2, -3, 1, 1}))
}

func getMaxLen(nums []int) int {
	maxLen := 0

	mostLeftPositiveIdx := -1
	mostLeftNegativeIdx := -1
	// Since we are not interested in real product value, only in it's sign (positive or negative)
	// Sounds logically to use booled instead of number
	isProductPositive := true
	for i, num := range nums {
		// Zero — is a very special case. It's like a border,
		// since we can't include it into subarray with positive product,
		// because zero will turn anything into zero. Each zero — is a fresh new start =)
		if num == 0 {
			mostLeftPositiveIdx = i
			mostLeftNegativeIdx = -1
			isProductPositive = true
			continue
		}

		if num < 0 {
			// While all nums is positive, nothing to do with isProductPositive variable
			// The only case when it's changing sign — is negative values
			isProductPositive = !isProductPositive
		}

		if isProductPositive {
			if maxLen < i-mostLeftPositiveIdx {
				maxLen = i - mostLeftPositiveIdx
			}
		} else {
			// To turn negative product into positive, required another negative to the left on the current one
			if mostLeftNegativeIdx == -1 {
				// -1 means that there was no negative products to the left of current one
				// If there was no negative value to the left of current one, this one will be that value for further ones
				mostLeftNegativeIdx = i
				continue
			}
			if maxLen < i-mostLeftNegativeIdx {
				maxLen = i - mostLeftNegativeIdx
			}
		}
	}
	return maxLen
}

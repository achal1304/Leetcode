package main

func combine(n int, k int) [][]int {
	result := [][]int{}
	combination := []int{}
	var backtrack func(index int)

	backtrack = func(index int) {
		if len(combination) == k {
			temp := make([]int, len(combination))
			copy(temp, combination)
			result = append(result, temp)
			return
		}

		for i := index; i < n+1; i++ {
			combination = append(combination, i)
			backtrack(i + 1)
			combination = combination[:len(combination)-1]
		}
	}

	backtrack(1)

	return result

}

package main

func permute(nums []int) [][]int {

	result := [][]int{}
	combination := []int{}

	var backtrack func(arr []int, index int)

	backtrack = func(arr []int, index int) {
		if len(combination) == len(nums) {
			temp := make([]int, len(combination))
			copy(temp, combination)
			result = append(result, temp)
			return
		}

		for i := index; i < len(nums); i++ {
			arr[i], arr[index] = arr[index], arr[i]
			combination = append(combination, arr[index])
			backtrack(arr, index+1)
			combination = combination[:len(combination)-1]
			arr[i], arr[index] = arr[index], arr[i]
		}
	}

	backtrack(nums, 0)
	return result

}

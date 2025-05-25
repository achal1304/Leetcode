package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	ques := []int{5, 5, 10, 9, 3, 6, 4, 2, 9, 10, 5}

	result := []int{1}
	resIndex := 0
	minNum := ques[0]
	for i := 1; i < len(ques); i++ {
		fmt.Println("resindex", resIndex, "prevnum", minNum, "currnum", ques[i], "res", result)

		if ques[i] < minNum {
			gcdAns := gcd(minNum, ques[i])
			fmt.Println("gcd ", gcdAns)
			if gcdAns > 1 {
				if resIndex == len(result) {
					result = append(result, 1)
				} else {
					result[resIndex]++
				}
			} else {
				result = append(result, 1)
				resIndex++
				minNum = ques[i]
			}
			minNum = ques[i]
		} else {
			gcdAns := gcd(minNum, ques[i])
			fmt.Println("gcd else ", gcdAns)
			if gcdAns > 1 {
				if resIndex == len(result) {
					result = append(result, 1)
				} else {
					result[resIndex]++
				}
			} else {
				result = append(result, 1)
				resIndex++
				minNum = ques[i]
			}
		}
	}

	fmt.Println(result)
}

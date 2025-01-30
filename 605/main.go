package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 1, 0, 0, 1}, 3))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		if n == 0 {
			return true
		}

		if i+2 < len(flowerbed) {
			return false
		}

		if flowerbed[i] == 0 {
			if i+1 < len(flowerbed) {
				if flowerbed[i+1] != 1 {
					flowerbed[i] = 1
					n -= 1
				} else {
					continue
				}
			} else {
				flowerbed[i] = 1
				n -= 1
			}
		}

		i += 1

		if (n*2)-1 > len(flowerbed[i:]) {
			return false
		}

	}
	if n != 0 {
		return false
	}
	return true
}

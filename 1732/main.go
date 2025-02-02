package main

import "fmt"

func largestAltitude(gain []int) int {
	currentAltitude := 0
	highestAltitude := 0

	for _, g := range gain {
		currentAltitude += g
		if currentAltitude > highestAltitude {
			highestAltitude = currentAltitude
		}
	}

	return highestAltitude
}

func main() {
	gain1 := []int{-5, 1, 5, 0, -7}
	fmt.Println("Highest Altitude for gain1:", largestAltitude(gain1))

	gain2 := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println("Highest Altitude for gain2:", largestAltitude(gain2))
}

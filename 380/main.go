package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
}

type RandomizedSet struct {
	valToIndex map[int]int
	values     []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valToIndex: make(map[int]int),
		values:     []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.valToIndex[val]; exists {
		return false
	}
	this.valToIndex[val] = len(this.values)
	this.values = append(this.values, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.valToIndex[val]
	if !exists {
		return false
	}

	lastVal := this.values[len(this.values)-1]
	this.values[idx] = lastVal
	this.valToIndex[lastVal] = idx

	this.values = this.values[:len(this.values)-1]
	delete(this.valToIndex, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	randomIdx := rand.Intn(len(this.values))
	return this.values[randomIdx]
}

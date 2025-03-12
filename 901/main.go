package main

import "fmt"

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(70))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(75))
	fmt.Println(obj.Next(85))
}

type StockSpanner struct {
	Stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{Stack: [][2]int{}}
}

func (this *StockSpanner) Next(price int) int {
	if len(this.Stack) == 0 {
		this.Stack = append(this.Stack, [2]int{price, 1})
		return 1
	}

	count := 1
	for len(this.Stack) > 0 && this.Stack[len(this.Stack)-1][0] <= price {
		count += this.Stack[len(this.Stack)-1][1]
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
	this.Stack = append(this.Stack, [2]int{price, count})
	return count
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(simplifyPath("/../"))
	// fmt.Println(simplifyPath("/.../a/..///////b/c/../d/./"))
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures"))
}

func simplifyPath(path string) string {
	stack := []string{}

	start := 0
	for i := 1; i < len(path); i++ {
		fmt.Println(stack)
		if path[i] == '/' {
			fmt.Println("inside if", stack, start, i)
			ele := path[start+1 : i]
			fmt.Println("ele", ele)
			switch ele {
			case "..":
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			case ".":
			case "/":
			default:
				if ele != "" {
					fmt.Println("appending to stack", path[start+1:i], start, i)
					stack = append(stack, path[start+1:i])
				}
			}

			start = i
		}
	}
	ele := path[start+1:]

	switch ele {
	case "..":
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	case ".":
	case "/":
	default:
		if ele != "" {
			fmt.Println("appending to stack", path[start+1:], start)
			stack = append(stack, path[start+1:])
		}
	}
	fmt.Println(stack, len(stack))

	if len(stack) == 0 {
		return "/"
	} else {
		return "/" + strings.Join(stack, "/")
	}

}

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
		if path[i] == '/' {
			ele := path[start+1 : i]
			switch ele {
			case "..":
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			case ".":
			case "/":
			default:
				if ele != "" {
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
			stack = append(stack, path[start+1:])
		}
	}

	if len(stack) == 0 {
		return "/"
	} else {
		return "/" + strings.Join(stack, "/")
	}

}

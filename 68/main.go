package main

import (
	"fmt"
	"strings"
)

// REVISIT

func main() {
	for _, ele := range fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20) {
		fmt.Println(ele)
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	n := len(words)
	index := 0

	for index < n {
		start := index
		currLen := len(words[index])
		index++

		for index < n && currLen+1+len(words[index]) <= maxWidth {
			currLen += 1 + len(words[index])
			index++
		}

		end := index
		wordsCount := end - start
		ans := ""

		if end == n || wordsCount == 1 {
			ans = strings.Join(words[start:end], " ")
			ans += strings.Repeat(" ", maxWidth-len(ans))
		} else {
			remainingSpaces := maxWidth - currLen + wordsCount - 1 // as we have added spaces to currlen in line 20 by adding 1, we need to minus here to get correct count
			spacesBetween := wordsCount - 1
			spacesToFill := remainingSpaces / spacesBetween
			extraSpace := remainingSpaces % spacesBetween

			for i := start; i < end; i++ {
				ans += words[i]
				if i < end-1 {
					currSpaceToFill := spacesToFill
					if extraSpace > 0 {
						currSpaceToFill++
						extraSpace--
					}

					ans += strings.Repeat(" ", currSpaceToFill)
				}
			}

		}
		result = append(result, ans)
	}
	return result
}

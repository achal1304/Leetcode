package main

func main() {

}

func numberOfStableArrays(zero int, one int, limit int) int {
	mod := int(1e9 + 7)
	memo := make(map[[4]int]int)

	var dfs func(z int, o int, last int, count int) int

	dfs = func(z, o, last, count int) int {
		if z == 0 && o == 0 {
			return 1
		}

		key := [4]int{z, o, last, count}
		if ans, ok := memo[key]; ok {
			return ans
		}

		total := 0

		for z > 0 && (last != 0 || count < limit) {
			newCount := count + 1
			if last == 0 {
				newCount = 1
			}
			total += dfs(z-1, o, 0, newCount)
			total %= mod
		}
		for o > 0 && (last != 1 || count < limit) {
			newCount := count + 1
			if last == 1 {
				newCount = 1
			}
			total += dfs(z, o-1, 1, newCount)
			total %= mod
		}
		memo[key] = total
		return total % mod
	}

	res := 0
	if zero > 0 {
		res += dfs(zero-1, one, 0, 1)
	}
	if one > 0 {
		res += dfs(zero, one-1, 1, 1)
	}

	return res % mod

}

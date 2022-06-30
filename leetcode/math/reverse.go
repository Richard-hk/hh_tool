package main

func Reverse(x int) int {
	var result int
	for x != 0 {
		mod := x % 10
		result = result*10 + mod
		x = x / 10
		if int(int32(result)) != result {
			return 0
		}
	}
	return result
}

package main

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	y := 0
	for x > y {
		y = x%10 + (y * 10)

		if x == y {
			return true
		}
		x = x / 10
	}

	// 121, 12321,

	return x == y
}

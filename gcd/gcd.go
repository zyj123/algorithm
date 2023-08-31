package gcd

func countBeautifulPairs(nums []int) int {
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if isCoprime(nums[i], nums[j]) {
				ans++
			}
		}
	}
	return ans
}

func isCoprime(x, y int) bool {
	return gcd(firstDigit(x), lastDigit(y)) == 1
}

func firstDigit(x int) int {
	for x >= 10 {
		x /= 10
	}
	return x
}

func lastDigit(x int) int {
	return x % 10
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

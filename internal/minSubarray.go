package internal

func minSubarray(nums []int, p int) int {
	remainders := make([]int, len(nums))
	sum := 0
	for k, num := range nums {
		remainders[k] = num % p
		sum += num % p
	}
	remainder := sum % p
	if remainder == 0 {
		return 0
	}
	res := -1
	for l := 0; l < len(remainders); l++ {
		if remainders[l] == remainder {
			return 1
		}
		sum2 := remainders[l]
		for r := l + 1; r < len(remainders); r++ {
			sum2 += remainders[r]
			if sum2 > remainder {
				break
			}
			if sum2 == remainder {
				res = minL(res, r-l+1)
			}
		}
	}
	return res
}

func minL(a, b int) int {
	if a == -1 {
		return b
	}
	if a <= b {
		return a
	}
	return b
}

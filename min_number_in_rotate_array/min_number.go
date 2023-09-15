package min_number_in_rotate_array

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	if numbers[l] < numbers[r] {
		return numbers[l]
	}
	for l < r {
		mid := l + (r-l)/2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else if numbers[mid] == numbers[r] {
			r--
		}
	}
	return numbers[l]
}

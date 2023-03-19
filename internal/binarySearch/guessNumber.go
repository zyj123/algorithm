package binarySearch

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	return 0
}
func guessNumber(n int) int {
	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		res := guess(mid)
		if res == 0 {
			return mid
		}
		if res == -1 {
			right = mid - 1
		}
		if res == 1 {
			left = mid + 1
		}
	}
	return -1
}

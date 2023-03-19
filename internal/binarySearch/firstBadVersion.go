package binarySearch

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		m := l + (r-l)/2
		isBad := isBadVersion(m)
		if !isBad {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

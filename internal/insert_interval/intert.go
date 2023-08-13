package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	i := 0
	interval2 := newInterval
	for _, interval1 := range intervals {
		if !isOverlap(interval1, interval2) {
			interval1, interval2 = sortInterval(interval1, interval2)
			intervals[i] = interval1
			i++
		} else {
			interval2 = merge(interval1, interval2)
		}
	}
	if i < len(intervals) {
		intervals[i] = interval2
	} else {
		intervals = append(intervals, interval2)
	}

	return intervals[:i+1]
}

func isOverlap(a, b []int) bool {
	return a[0] <= b[1] && a[1] >= b[0]
}

func merge(a, b []int) []int {
	var l, r int
	if a[0] < b[0] {
		l = a[0]
	} else {
		l = b[0]
	}
	if a[1] > b[1] {
		r = a[1]
	} else {
		r = b[1]
	}
	return []int{l, r}
}

func sortInterval(a, b []int) ([]int, []int) {
	if a[0] < b[0] {
		return a, b
	} else {
		return b, a
	}
}

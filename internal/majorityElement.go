package internal

func majorityElement(nums []int) int {
	var (
		cur    int
		curCnt int
	)
	for _, v := range nums {
		if curCnt == 0 {
			cur = v
			curCnt++
			continue
		}
		if v == cur {
			curCnt++
			continue
		}
		curCnt--
	}
	return cur
}

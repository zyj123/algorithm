package internal

func minimumRecolors(blocks string, k int) int {
	var (
		l    = 0
		r    = -1
		cnt  = 0
		mCnt = 0
	)
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			cnt++
		}
		r++
	}
	mCnt = cnt
	for ; r < len(blocks); l, r = l+1, r+1 {
		if l == 0 {
			continue
		}
		if blocks[l-1] == 'W' && blocks[r] == 'B' {
			cnt -= 1
		}
		if blocks[l-1] == 'B' && blocks[r] == 'W' {
			cnt += 1
		}
		mCnt = min(cnt, mCnt)
	}
	return mCnt
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

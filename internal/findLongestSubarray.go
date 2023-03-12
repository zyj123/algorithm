package internal

func findLongestSubarray(array []string) []string {
	var l, r int
	mp := make(map[int]int)
	// diff = number - letter
	preDiff := 0
	mp[0] = -1
	for k, v := range array {
		if (v >= "A" && v <= "Z") || (v >= "a" && v <= "z") {
			preDiff -= 1
		} else {
			preDiff += 1
		}
		if _, ok := mp[preDiff]; !ok {
			mp[preDiff] = k
		} else {
			if k-mp[preDiff] > r-l {
				l, r = mp[preDiff]+1, k+1
			}
		}
	}
	return array[l:r]
}

func isNum() {

}

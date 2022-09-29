package internal

func longestCommonPrefix(strs []string) (ret string) {
	var i int
	defer func() {
		ret = strs[0][:i]
	}()
	for ; i < len(strs[0]); i++ {
		var chr byte
		for _, str := range strs {
			if i >= len(str) {
				return
			}
			if chr == 0 {
				chr = str[i]
				continue
			}
			if chr != str[i] {
				return
			}
		}
	}
	return
}

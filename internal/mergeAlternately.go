package internal

func mergeAlternately(word1 string, word2 string) string {
	var (
		l1 = len(word1)
		l2 = len(word2)
	)
	ret := make([]byte, 0, l1+l2)
	i1, i2 := 0, 0
	for ; i1 < l1 && i2 < l2; i1, i2 = i1+1, i2+1 {
		ret = append(ret, word1[i1])
		ret = append(ret, word2[i2])
	}
	if i1 < l1 {
		ret = append(ret, word1[i1:]...)
	}
	if i2 < l2 {
		ret = append(ret, word2[i2:]...)
	}
	return string(ret)
}

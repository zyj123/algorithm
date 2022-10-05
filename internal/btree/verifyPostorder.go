package btree

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	n := len(postorder)
	rootVal := postorder[n-1]
	i := 0
	for ; i < n-1; i++ {
		if postorder[i] > rootVal {
			break
		}
	}
	for j := i; j < n-1; j++ {
		if postorder[j] < rootVal {
			return false
		}
	}
	return verifyPostorder(postorder[:i]) && verifyPostorder(postorder[i:n-1])
}

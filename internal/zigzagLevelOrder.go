package internal

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		curLevel := make([]int, 0, n)
		for i := 0; i < n; i++ {
			if len(ret)%2 == 0 {
				curLevel = append(curLevel, queue[i].Val)
			} else {
				curLevel = append(curLevel, queue[n-i-1].Val)
			}

			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, curLevel)
		queue = queue[n:]
	}
	return ret
}

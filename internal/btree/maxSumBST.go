package btree

func maxSumBST(root *TreeNode) int {
	var maxSum int

	var checkAndSum func(node *TreeNode) (bool, int)
	checkAndSum = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		lCheck, lSum := checkAndSum(node.Left)
		rCheck, rSum := checkAndSum(node.Right)

		if node.Left != nil && node.Left.Val >= node.Val {
			return false, 0
		}
		if node.Right != nil && node.Right.Val <= node.Val {
			return false, 0
		}

		if !(lCheck && rCheck) {
			return false, 0
		}

		sum := node.Val + lSum + rSum
		if sum > maxSum {
			maxSum = sum
		}
		return true, sum
	}

	checkAndSum(root)
	return maxSum
}

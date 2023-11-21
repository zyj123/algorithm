package permute

import "fmt"

func main2() {

	fmt.Println(12^15, 5^15)
}

//算法： 全排列
//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//输入： nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//nums中的所有整数互不相同

func example(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)

	var bt func([]int, []int)

	bt = func(arr []int, item []int) {
		if len(item) == n {
			res = append(res, item)
			return
		}
		for i := range arr {
			arr[0], arr[i] = arr[i], arr[0]
			bt(arr[1:], append(item, arr[0]))
			arr[i], arr[0] = arr[0], arr[i]
		}
	}

	bt(nums, []int{})

	return res
}

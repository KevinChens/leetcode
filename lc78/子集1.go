package lc78

/*
分析：这是一个典型的应用回溯法的题目，简单来说就是穷尽所有可能性,利用回溯算法模板
回溯还不太懂，暂略
时间复杂度：
空间复杂度：O(n)，需要临时空间list
 */

func subsets1(nums []int) [][]int {
	// 保存最终结果
	result := make([][]int, 0)
	// 保存中间结果
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

// nums给定的集合
// pos下次添加到集合中的元素位置索引
// list临时结果集合（每次需要复制保存）
// result最终结果

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	// 把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	// 选择、处理结果、再撤销选择
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0:len(list)-1]
	}
}
package main

import "code/test_utils"

/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

提示：
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109

https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150
*/

func majorityElement(nums []int) int {
	/*
		计数，比较计数大小
	*/
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num] += 1
	}
	maxNum := 0
	maxCount := 0
	for num, count := range counter {
		if count > maxCount {
			maxNum = num
			maxCount = count
		}
	}
	return maxNum
}

func main() {
	test_utils.ArrayCheckResult(
		majorityElement,
		test_utils.ArrayItem{Nums: []int{3, 2, 3}, Result: 3},
		test_utils.ArrayItem{Nums: []int{2, 2, 1, 1, 1, 2, 2}, Result: 2},
	)
}

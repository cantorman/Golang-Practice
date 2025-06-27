package solution

import "sort"

func TwoSum(nums []int, target int) [2]int {
	answer := []int{0, 0}
outer:
	for i := range nums {
		for j := range i {
			if nums[i]+nums[j] == target {
				answer[0] = i
				answer[1] = j
				break outer
			}
		}
	}
	sort.Ints(answer)
	return [2]int{answer[0], answer[1]}
}

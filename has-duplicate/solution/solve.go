package solution

import "fmt"

func HasDuplicate(nums []int) (answer bool) {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	fmt.Println("set: ", set)
	return len(set) != len(nums)
}

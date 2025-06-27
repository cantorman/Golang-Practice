package solution

// import "fmt"

func TopKFrequent(nums []int, k int) (modes []int) {
	for range k {
		mode_i := mode(nums)
		modes = append(modes, mode_i)
		nums = delItem(nums, mode_i)
	}

	return
}

func mode(numList []int) (answer int) {
	counts := map[int]int{}
	maxCount := 0
	for i := range len(numList) {
		num := numList[i]
		counts[num]++
		// fmt.Printf("i %d num %d counts[num] %d\n", i, num, counts[num])
		if counts[num] > maxCount {
			maxCount = counts[num]
			answer = num
		}
	}
	return
}

func delItem(numList []int, item int) (newList []int) {
	for i := range numList {
		if numList[i] != item {
			newList = append(newList, numList[i])
		}
	}
	return
}

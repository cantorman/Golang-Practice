package main

import (
	"fmt"
	"slices"
	"top-k-frequent/solution"
)

func main() {
	/*
		Given an integer array nums and an integer k, return the k most frequent elements within the array.

		The test cases are generated such that the answer is always unique.

		You may return the output in any order.
	*/
	testCases := []solution.TestCase{
		{Input: solution.Input{Nums: []int{1, 2, 2, 3, 3, 3}, K: 2}, Output: solution.Output{2, 3}},
		{Input: solution.Input{Nums: []int{7, 7}, K: 1}, Output: solution.Output{7}},
	}

	suite_status := "PASSED"
	for case_no, testcase := range testCases {
		expected := testcase.Output
		slices.Sort(expected)
		actual := solution.Output(solution.TopKFrequent(testcase.Input.Nums, testcase.Input.K))
		slices.Sort(actual)
		// fmt.Println("expected:", expected, "type:", reflect.TypeOf(expected))
		// fmt.Println("actual:", actual, "type:", reflect.TypeOf(actual))

		if !slices.Equal(actual, expected) {
			suite_status = "FAILED"
			fmt.Println("Test Case:", case_no, ": FAILED")
			fmt.Println("actual", actual, "!= expected", expected)
		} else {
			fmt.Println("Test Case:", case_no, ": PASSED")
		}
	}
	fmt.Println("Suite:", suite_status)
}

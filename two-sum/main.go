package main

import (
	"fmt"
	"two-sum/solution"
)

type input struct {
	nums   []int
	target int
}
type output [2]int

type testcase struct {
	Input  input
	Output output
}

func main() {
	/*
		Given an array of integers nums and an integer target,
		return the indices i and j such that nums[i] + nums[j] == target and i != j.

		You may assume that every input has exactly one pair of indices i and j
		that satisfy the condition.

		Return the answer with the smaller index first.
	*/
	testCases := []testcase{
		{Input: input{[]int{3, 4, 5, 6}, 7}, Output: output{0, 1}},
		{Input: input{[]int{4, 5, 6}, 10}, Output: output{0, 2}},
		{Input: input{[]int{5, 5}, 10}, Output: output{0, 1}},
	}

	suite_status := "PASSED"
	for case_no, testcase := range testCases {
		expected := testcase.Output
		actual := solution.TwoSum(testcase.Input.nums, testcase.Input.target)
		succeeded := actual == expected
		case_status := "PASSED"
		if !succeeded {
			suite_status = "FAILED"
			case_status = "FAILED"
		}
		fmt.Println("Test Case:", case_no, ":", case_status)
	}
	fmt.Println("Suite:", suite_status)
}

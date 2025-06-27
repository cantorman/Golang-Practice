package main

import (
	"fmt"
	"has-duplicate/solution"
)

type Case struct {
	Input  []int
	Output bool
}

func main() {
	/*
		Given an integer array nums,
		return true if any value appears more than once in the array,
		otherwise return false.
	*/
	testCases := []Case{
		{Input: []int{1, 2, 3, 3}, Output: true},
		{Input: []int{1, 2, 3, 4}, Output: false},
	}

	suite_status := "PASSED"
	for case_no, testcase := range testCases {
		expected := testcase.Output
		actual := solution.HasDuplicate(testcase.Input)
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

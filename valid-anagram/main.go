package main

import (
	"fmt"
	"valid-anagram/solution"
)

type stringpair [2]string
type testcase struct {
	Input  stringpair
	Output bool
}

func main() {
	/*
		Given an integer array nums,
		return true if any value appears more than once in the array,
		otherwise return false.
	*/
	testCases := []testcase{
		{Input: stringpair{"racecar", "carrace"}, Output: true},
		{Input: stringpair{"jar", "jam"}, Output: false},
	}

	suite_status := "PASSED"
	for case_no, testcase := range testCases {
		expected := testcase.Output
		actual := solution.IsAnagram(testcase.Input[0], testcase.Input[1])
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

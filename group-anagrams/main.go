package main

import (
	"fmt"
	"group-anagrams/solution"
	"reflect"
)

func main() {
	/*
		Given an array of strings strs, group all anagrams together into sublists.
		You may return the output in any order.

		An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.
	*/
	testCases := []solution.TestCase{
		{Input: solution.Input{"act", "pots", "tops", "cat", "stop", "hat"}, Output: solution.Output{[]string{"hat"}, []string{"act", "cat"}, []string{"stop", "pots", "tops"}}},
		{Input: solution.Input{"x"}, Output: solution.Output{[]string{"x"}}},
	}

	suite_status := "PASSED"
	for case_no, testcase := range testCases {
		expected := testcase.Output
		solution.OrderMyMap(expected)
		actual := solution.Output(solution.GroupAnagrams(testcase.Input))
		solution.OrderMyMap(actual)
		// fmt.Println("expected:", expected, "type:", reflect.TypeOf(expected))
		// fmt.Println("actual:", actual, "type:", reflect.TypeOf(actual))

		succeeded := reflect.DeepEqual(actual, expected)
		case_status := "PASSED"
		if !succeeded {
			suite_status = "FAILED"
			case_status = "FAILED"
		}
		fmt.Println("Test Case:", case_no, ":", case_status)
	}
	fmt.Println("Suite:", suite_status)
}

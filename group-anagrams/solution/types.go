package solution

type Input []string
type Output [][]string

type TestCase struct {
	Input  Input
	Output Output
}

// we need a map whose keys are a comparable thing, and whose values are slices of strings
type mapkeytype string
type mapkeyval []string
type anagramgroups map[mapkeytype]mapkeyval

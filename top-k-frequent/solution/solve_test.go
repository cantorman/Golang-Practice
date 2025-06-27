package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Mode(t *testing.T) {
	testData := []int{1, 1, 2, 3}
	expected := 1

	actual := mode(testData)

	assert.Equal(t, expected, actual)
}

func Test_DelItem(t *testing.T) {
	testData := []int{1, 1, 2, 3}
	target := 1
	expected := []int{2, 3}

	actual := delItem(testData, target)

	assert.Equal(t, expected, actual)
}

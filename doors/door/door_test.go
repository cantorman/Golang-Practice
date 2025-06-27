package door

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetNewDoor(t *testing.T) {
	d := GetNewDoor()

	assert.Equal(t, "CLOSED", d.State)
}

func Test_Toggle(t *testing.T) {
	d := GetNewDoor()
	Toggle(&d)
	assert.Equal(t, "OPEN", d.State)
	Toggle(&d)
	assert.Equal(t, "CLOSED", d.State)
}

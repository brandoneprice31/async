package async

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallel(t *testing.T) {
	// Test no errors.
	x := 0
	ee := Parallel(
		func() error {
			x++
			return nil
		},

		func() error {
			x += 2
			return nil
		},
		func() error {
			x += 3
			return nil
		},
	)
	assert.True(t, ee.IsEmpty())
	assert.Equal(t, 6, x)

	// Test witb errors.
	y := 0
	ee = Parallel(
		func() error {
			y++
			return errors.New("error 1")
		},

		func() error {
			y++
			return errors.New("error 2")
		},

		func() error {
			y--
			return errors.New("error 3")
		},
	)
	assert.False(t, ee.IsEmpty())
	assert.Equal(t, 3, len(ee.All()))
	assert.NotNil(t, ee.ToError())
	assert.Equal(t, 1, y)
}

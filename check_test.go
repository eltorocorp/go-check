package check_test

import (
	"errors"
	"testing"

	"github.com/eltorocorp/go-check"
	"github.com/stretchr/testify/assert"
)

func Test_CheckTrap(t *testing.T) {
	testCases := []struct {
		name        string
		fn          func()
		expectPanic bool
		expectedErr error
	}{
		{
			name:        "no panic",
			fn:          func() {},
			expectPanic: false,
			expectedErr: nil,
		},
		{
			name:        "panic with error",
			fn:          func() { panic(errors.New("test")) },
			expectPanic: false,
			expectedErr: errors.New("test"),
		},
		{
			// Trap will itself panic if the recovered value is not convertable
			// to an error type.
			name:        "panic without error",
			fn:          func() { panic("not error") },
			expectPanic: true,
			expectedErr: nil,
		},
		{
			// Go's recover function doesn't distingtuish between a goroutine
			// that did not panic vs a goroutine that supplied nil to the panic
			// function. Semantically speaking, a nil panic is the same as no
			// panic. Thus, a nil panic is disregarded.
			name:        "nil panic",
			fn:          func() { panic(nil) },
			expectPanic: false,
			expectedErr: nil,
		},
	}
	for _, testCase := range testCases {
		testFn := func(t *testing.T) {
			if testCase.expectPanic {
				assert.Panics(t, func() {
					check.Trap(func() { testCase.fn() })
				})
			} else {
				err := check.Trap(func() { testCase.fn() })
				if testCase.expectedErr == nil {
					assert.NoError(t, err)
				} else {
					assert.EqualError(t, err, testCase.expectedErr.Error())
				}
			}
		}
		t.Run(testCase.name, testFn)
	}
}

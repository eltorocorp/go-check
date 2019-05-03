package check_test

import (
	"errors"
	"testing"

	"github.com/eltorocorp/go-check"
	"github.com/eltorocorp/go-check/mocks/mock_check"
	"github.com/golang/mock/gomock"
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

func Test_CheckTrapTx(t *testing.T) {
	testCases := []struct {
		name        string
		closure     func(check.Tx)
		expRollback bool
		expCommit   bool
		expErr      error
	}{
		{
			name:        "transaction is committed",
			closure:     func(tx check.Tx) {},
			expRollback: false,
			expCommit:   true,
			expErr:      nil,
		},
		{
			name:        "transaction is rolled back",
			closure:     func(tx check.Tx) { panic(errors.New("test error"))},
			expRollback: true,
			expCommit:   false,
			expErr:      errors.New("test error"),
		},
	}

	for _, testCase := range testCases {
		testFn := func(t *testing.T) {
			mc := gomock.NewController(t)
			defer mc.Finish()

			tx := mock_check.NewMockTx(mc)
			if testCase.expCommit {
				tx.EXPECT().Commit().Return(nil).Times(1)
				tx.EXPECT().Rollback().Times(0)
			}
			if testCase.expRollback {
				tx.EXPECT().Commit().Times(0)
				tx.EXPECT().Rollback().Return(nil).Times(1)
			}

			txProvider := mock_check.NewMockTxProvider(mc)
			txProvider.EXPECT().Begin().Return(tx, nil)
			err := check.TrapTx(txProvider, testCase.closure)
			if testCase.expErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, testCase.expErr.Error())
			}
		}
		t.Run(testCase.name, testFn)
	}
}

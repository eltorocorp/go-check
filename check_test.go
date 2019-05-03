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
		txProviderErr error
		closure     func(check.Tx)
		commitErr error
		rollbackErr error
		expRollback bool
		expCommit   bool
		expErr      error
	}{
		{
			name:        "transaction is committed",
			txProviderErr: nil,
			closure:     func(tx check.Tx) {},
			commitErr: nil,
			rollbackErr: nil,
			expRollback: false,
			expCommit:   true,
			expErr:      nil,
		},
		{
			name:        "transaction is rolled back",
			txProviderErr: nil,
			closure:     func(tx check.Tx) { panic(errors.New("test error"))},
			commitErr: nil,
			rollbackErr: nil,
			expRollback: true,
			expCommit:   false,
			expErr:      errors.New("test error"),
		},
		{
			name:        "txProvider error is propogated without calling closure",
			txProviderErr: errors.New("test error"),
			closure:     func(tx check.Tx) { panic(errors.New("this shouldnt get called"))},
			commitErr: nil,
			rollbackErr: nil,
			expRollback: false,
			expCommit:   false,
			expErr:      errors.New("test error"),
		},
		{
			name:        "commit error is propogated",
			txProviderErr: nil,
			closure:     func(tx check.Tx) { },
			commitErr: errors.New("test error"),
			rollbackErr: nil,
			expRollback: false,
			expCommit:   true,
			expErr:      errors.New("test error"),
		},
		{
			// If an error occurs during a rollback, we want to make sure that
			// the original error (the error that required the rollback) is
			// not lost.
			name:        "rollback error includes original error",
			txProviderErr: nil,
			closure:     func(tx check.Tx) { panic(errors.New("original error"))},
			commitErr: nil,
			rollbackErr: errors.New("rollback error"),
			expRollback: true,
			expCommit:   false,
			expErr:      errors.New("rollback error\noriginal error"),
		},
	}

	for _, testCase := range testCases {
		testFn := func(t *testing.T) {
			mc := gomock.NewController(t)
			defer mc.Finish()

			txProvider := mock_check.NewMockTxProvider(mc)
			if testCase.txProviderErr != nil {
				txProvider.EXPECT().Begin().Return(nil, testCase.txProviderErr)
			} else {
				tx := mock_check.NewMockTx(mc)
				if testCase.expCommit {
					tx.EXPECT().Commit().Return(testCase.commitErr).Times(1)
					tx.EXPECT().Rollback().Times(0)
				}
				if testCase.expRollback {
					tx.EXPECT().Commit().Times(0)
					tx.EXPECT().Rollback().Return(testCase.rollbackErr).Times(1)
				}
				txProvider.EXPECT().Begin().Return(tx, nil)
			}

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

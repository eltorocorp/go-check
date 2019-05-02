package check

// Trap executes fn and attempts to recover if fn panics.
// If fn supplied an error type when panicking, trap will cancel the panic and
// return the original error.
// If fn supplied anything other than an error to the original call to panic,
// Trap will also panic. This is done intentionally to focus the
// use of Trap as an error handling mechanism, not to replace go's native
// recover function.
// If fn does not panic, or calls panic(nil), Trap will return nil.
func Trap(fn func()) (err error) {
	return TrapTx(nullDB{}, func(_ Tx) {
		fn()
	})
}

// TrapTx follows the same rules as Trap, but also executes fn within the
// scope of a transaction.
// The transaction is rolled back if fn panics, and is committed if fn does not
// panic.
func TrapTx(txProvider TxProvider, fn func(Tx)) (err error) {
	var tx Tx
	defer func() {
		p := recover()
		if p == nil {
			err = tx.Commit()
			return
		}
		if tx != nil {
			tx.Rollback()
		}
		if e, ok := p.(error); ok {
			err = e
		} else {
			panic(p)
		}
	}()

	tx, err = txProvider.Begin()
	if err != nil {
		return err
	}

	fn(tx)
	return
}

// Bool expects a value and an error.
// If err is not nil, the function will panic.
// Otherwise the value is returned,
func Bool(b bool, err error) bool {
	Err(err)
	return b
}

// Int expects a value and an error.
// If err is not nil, the function will panic.
// Otherwise the value is returned,
func Int(i int, err error) int {
	Err(err)
	return i
}

// Float64 expects a value and an error.
// If err is not nil, the function will panic.
// Otherwise the value is returned,
func Float64(f float64, err error) float64 {
	Err(err)
	return f
}

// String expects a value and an error.
// If err is not nil, the function will panic.
// Otherwise the value is returned,
func String(s string, err error) string {
	Err(err)
	return s
}

// Iface expects a value and an error.
// If err is not nil, the function will panic.
// Otherwise the value is returned,
func Iface(i interface{}, err error) interface{} {
	Err(err)
	return i
}

// Err evaluates an error,
// If the error is not nil, the function will panic.
func Err(err error) {
	if err != nil {
		panic(err)
	}
}

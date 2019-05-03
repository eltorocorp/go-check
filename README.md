go-check
========

`go get github.com/eltorocorp/go-check`


Example
=======
Here's a block of code with conventional go error handling
```go
func (a *API) CalculateContributorScore(user usercontextiface.UserContextAPI, contributorID int) (float64, error) {
	if user == nil {
		return 0, inertiaerrors.ErrPermissionDenied
	}

	admin := &administration.API{DB: a.DB}
	settings, err := admin.GetSettings(user)
	if err != nil {
		return 0, err
	}

	avgGoalScore, err := a.averageRawScoreForContributor(contributorID)
	if err != nil {
		return 0, err
	}

	leadID, err := a.getLeadIDForContributor(contributorID)
	if err != nil {
		return 0, err
	}

	avgLeadScore, err := a.averageScoreForLead(leadID)
	if err != nil {
		return 0, err
	}

	return calculateBiasedScore(avgGoalScore, avgLeadScore, settings.BaseScore), nil
}
```

With `check.Trap`, we can restate as follows:

```go
func (a *API) CalculateContributorScore(user usercontextiface.UserContextAPI, contributorID int) (out float64, err error) {
    err = check.Trap(func() {
        if user == nil {
            panic(ErrPermissionDenied)
        }

        admin := &administration.API{DB: a.DB}
        settings:= check.IFace(admin.GetSettings(user)).(*models.Setting)
        avgGoalScore:= check.Float64(a.averageRawScoreForContributor(contributorID))
        leadID:= check.Int(a.getLeadIDForContributor(contributorID))
        avgLeadScore:= check.Float64(a.averageScoreForLead(leadID))
        out = check.Float64(calculateBiasedScore(avgGoalScore, avgLeadScore, settings.BaseScore))
    })
    return
}
```

Notice how the overall intent of the code is now much more clear since all of the transaction and error handling noise has been abstracted away.

How it Works
============
Helper functions (such as `check.Float64`) are wrapped around functions that return a value and an error. These helper functions panic if the error is not nil, and otherwise return value.
`check.Trap` in turn, traps the panic caused by the helper function, retrieves the error that caused the panic, and returns the error.

Transaction Handling
====================
`check` can also be used to simplify database transaction handling in line with err handling.

This is done with the use of `check.TrapTx`, which is similar to `check.Trap`, but also manages a transaction within the same context as any errors.

Here is an example of a database transaction and errors being handled conventionally:

```go
func (c *Context) ExpireSessions() error {
	if c.SessionToken() == "" {
		return nil
	}

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	sessions, err := models.GetAllSessions(tx)
	if err != nil {
		return err
	}
	for _, session := range sessions {
		if session.PersonID == c.UserID() {
			err = session.Delete(tx)
			if err != nil {
				log.Println(err)
				err = tx.Rollback()
				if err != nil {
					return err
				}
			}
		}
	}
	return tx.Commit()
}
```

Here is the same code, but rewritten with `check.TrapTx`:


```go
func (c *Context) ExpireSessions() error {
    return check.TrapTx(check.UseDB(c.DB), func(tx check.Tx) {
        if c.SessionToken() == "" {
            return
        }

        sessions:= check.IFace(models.GetAllSessions(tx)).([]*models.Session)
        for _, session := range sessions {
            if session.PersonID == c.UserID() {
                check.Err(session.Delete(tx))
            }
        }
    })
}
```

Just as in the `check.Trap` example, notice how the overall intent of the code is now much more clear since all of the transaction and error handling noise has been abstracted away.

How it Works
============
In this case, a database reference is passed into `check.TrapTx`. Internally, `check` will create a transaction for the underlaying database. That transaction is then passed into the closure supplied to `TrapFx`. If any errors occur within the closure, the helper functions (such as `check.Err`) will panic. `check` will then recover from the panic, and automatically rollback the transaction. If the closure returns without any panicks, `check` will automatically commit the transaction.

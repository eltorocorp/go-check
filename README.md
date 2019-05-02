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

How it Works
============
Helper functions (such as `check.Float64`) are wrapped around functions that return a value and an error. These helper functions panic if the error is not nil, and otherwise return value.
`check.Trap` in turn, traps the panic caused by the helper function, retrieves the error that caused the panic, and returns the error.

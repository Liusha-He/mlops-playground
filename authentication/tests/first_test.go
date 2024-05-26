package tests

import (
	"github.com/cucumber/godog"
	"testing"
)

func iEat(arg1 int) error {
	return godog.ErrPending
}

func thereAreGodogs(arg1 int) error {
	return godog.ErrPending
}

func thereShouldBeRemaining(arg1 int) error {
	return godog.ErrPending
}

func TestFeature(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			ctx.Step(`^there are (\d+) godogs$`, thereAreGodogs)
			ctx.Step(`^I eat (\d+)$`, iEat)
			ctx.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

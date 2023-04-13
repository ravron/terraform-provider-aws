package events_test

import (
	"testing"

	tfevents "github.com/hashicorp/terraform-provider-aws/internal/service/events"
)

func TestRuleEnabledFromState(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		TestName        string
		State           string
		ExpectedError   bool
		ExpectedEnabled bool
	}{
		{
			TestName:      "empty state",
			ExpectedError: true,
		},
		{
			TestName:      "invalid state",
			State:         "UNKNOWN",
			ExpectedError: true,
		},
		{
			TestName:        "enabled",
			State:           "ENABLED",
			ExpectedEnabled: true,
		},
		{
			TestName:        "disabled",
			State:           "DISABLED",
			ExpectedEnabled: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.TestName, func(t *testing.T) {
			t.Parallel()

			gotEnabled, err := tfevents.RuleEnabledFromState(testCase.State)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error, got no error")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("got unexpected error: %s", err)
			}

			if gotEnabled != testCase.ExpectedEnabled {
				t.Errorf("got enabled %t, expected %t", gotEnabled, testCase.ExpectedEnabled)
			}
		})
	}
}

func TestRuleStateFromEnabled(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		TestName      string
		Enabled       bool
		ExpectedState string
	}{
		{
			TestName:      "enabled",
			Enabled:       true,
			ExpectedState: "ENABLED",
		},
		{
			TestName:      "disabled",
			Enabled:       false,
			ExpectedState: "DISABLED",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.TestName, func(t *testing.T) {
			t.Parallel()

			gotState := tfevents.RuleStateFromEnabled(testCase.Enabled)

			if gotState != testCase.ExpectedState {
				t.Errorf("got enabled %s, expected %s", gotState, testCase.ExpectedState)
			}
		})
	}
}

package models

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSubTypes(t *testing.T) {
	var typesSubTypesUnderTest = []struct {
		activityType  Activity_Types
		validSubTypes []Activity_SubTypes
	}{
		{activityType: Breast, validSubTypes: []Activity_SubTypes{Undefined}},
		{activityType: Bottle, validSubTypes: []Activity_SubTypes{Formula, Breast_Milk, Goats_Milk, Cows_Milk}},
	}

	for _, test := range typesSubTypesUnderTest {
		t.Run(string(test.activityType), func(t *testing.T) {
			t.Parallel()
			var activity Activity = Activity{
				Type: test.activityType,
			}
			var got []Activity_SubTypes = activity.getSubTypes()
			if diff := cmp.Diff(test.validSubTypes, got); diff != "" {
				t.Logf("Retrieved sub types: %q, are not equal", got)
				t.Fatalf("Activity Type (%q) valid sub types (%q) and expected valid sub types (%q) do not match.", test.activityType, got, test.validSubTypes)
			} else {
				t.Logf("Retrieved sub types: %q, are equal", got)
			}
		})
	}
}

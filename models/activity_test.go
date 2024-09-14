package models

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
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
			var diff = cmp.Diff(test.validSubTypes, got)
			assert.Equal(t, "", diff)
		})
	}
}

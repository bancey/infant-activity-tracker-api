package models

import "testing"

func TestActivityTypes(t *testing.T) {
	var typesUnderTest = []struct {
		subType        Activity_Types
		expectedResult string
	}{
		{Breast, "breast"},
		{Bottle, "bottle"},
		{Nappy, "nappy"},
	}

	for _, test := range typesUnderTest {
		t.Run(string(test.subType), func(t *testing.T) {
			t.Parallel()
			if string(test.subType) != test.expectedResult {
				t.Fatalf("Activity Type (%q) does not match expected string value (%q).", test.subType, test.expectedResult)
			}
		})
	}
}

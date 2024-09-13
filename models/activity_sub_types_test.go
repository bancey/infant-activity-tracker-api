package models

import "testing"

func TestActivitySubTypes(t *testing.T) {
	var typesUnderTest = []struct {
		subType        Activity_SubTypes
		expectedResult string
	}{
		{Left, "left_breast"},
		{Right, "right_breast"},
		{Breast_Milk, "breast_milk"},
	}

	for _, test := range typesUnderTest {
		t.Run(string(test.subType), func(t *testing.T) {
			t.Parallel()
			if string(test.subType) != test.expectedResult {
				t.Fatalf("Activity Sub Type (%q) does not match expected string value (%q).", test.subType, test.expectedResult)
			}
		})
	}
}

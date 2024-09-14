package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			assert.Equal(t, test.expectedResult, string(test.subType))
		})
	}
}

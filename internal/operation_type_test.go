package internal

import "testing"

func TestOperationTypeFromString(t *testing.T) {
	testCases := []struct {
		typeAsString          string
		expectedOperationType OperationType
	}{
		{"avg", Average},
		{"min", Minimum},
		{"max", Maximum},
		{"bla", UnknownOperation},
	}

	for _, tc := range testCases {
		currentOperationType := OperationTypeFromString(tc.typeAsString)
		if currentOperationType != tc.expectedOperationType {
			t.Errorf("Wrongly detected operation type. Expected %d, got %d.", tc.expectedOperationType, currentOperationType)
		}
	}
}

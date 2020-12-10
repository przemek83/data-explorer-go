package internal

import "testing"

func TestOperationFromString(t *testing.T) {
	testCases := []struct {
		asString          string
		expectedOperation Operation
	}{
		{"avg", Average},
		{"min", Minimum},
		{"max", Maximum},
		{"bla", UnknownOperation},
	}

	for _, tc := range testCases {
		currentOperation := OperationFromString(tc.asString)
		if currentOperation != tc.expectedOperation {
			t.Errorf("Wrongly detected operation. Expected %d, got %d.", tc.expectedOperation, currentOperation)
		}
	}
}

package internal

import "testing"

func TestColumnTypeFromString(t *testing.T) {
	testCases := []struct {
		typeAsString       string
		expectedColumnType ColumnType
	}{
		{"integer", NumericColumn},
		{"string", StringColumn},
		{"", Unknown},
		{"bla", Unknown},
	}

	for _, tc := range testCases {
		currentColumnType := ColumnTypeFromString(tc.typeAsString)
		if currentColumnType != tc.expectedColumnType {
			t.Errorf("Wrongly detected column type. Expected %d, got %d.", tc.expectedColumnType, currentColumnType)
		}
	}
}

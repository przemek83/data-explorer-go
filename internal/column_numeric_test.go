package internal

import "testing"

func TestGetType(t *testing.T) {
	column := ColumnNumeric{}
	expectedType := NumericColumn
	if currentType := column.GetType(); currentType != expectedType {
		t.Errorf("Wrong type returned, got %v, want %v", currentType, expectedType)
	}
}

package internal

import "testing"

func TestGetType(t *testing.T) {
	column := ColumnNumeric{}
	expectedType := NumericColumn
	if currentType := column.GetType(); currentType != expectedType {
		t.Errorf("Wrong type returned, got %v, want %v", currentType, expectedType)
	}
}

func TestGet(t *testing.T) {
	expectedValues := []int{1, 2, 3, 4}
	column := ColumnNumeric{expectedValues}
	for i, expectedValue := range expectedValues {
		if currentValue := column.Get(i); currentValue != expectedValue {
			t.Errorf("Wrong value returned, got %v, want %v", currentValue, expectedValue)
		}
	}
}

package internal

import "testing"

func TestColumnNumericGetType(t *testing.T) {
	column := ColumnNumeric{}
	expectedType := NumericColumn
	if currentType := column.GetType(); currentType != expectedType {
		t.Errorf("Wrong type returned, got %v, want %v", currentType, expectedType)
	}
}

func TestColumnNumericGet(t *testing.T) {
	expectedValues := []int{1, 2, 3, 4}
	column := ColumnNumeric{expectedValues}
	for i, expectedValue := range expectedValues {
		if currentValue := column.Get(i); currentValue != expectedValue {
			t.Errorf("Wrong value returned, got %v, want %v", currentValue, expectedValue)
		}
	}
}

func TestColumnNumericAppend(t *testing.T) {
	values := []int{1, 2, 3, 4}
	column := ColumnNumeric{values}
	expectedValue := 7
	column.Append(expectedValue)
	if currentValue := column.Get(len(values)); currentValue != expectedValue {
		t.Errorf("Wrong value returned, got %v, want %v", currentValue, expectedValue)
	}
}

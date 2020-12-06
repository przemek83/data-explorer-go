package internal

import "testing"

func TestColumnStringGetType(t *testing.T) {
	column := ColumnString{}
	expectedType := StringColumn
	if currentType := column.GetType(); currentType != expectedType {
		t.Errorf("Wrong type returned, got %v, want %v", currentType, expectedType)
	}
}

func TestColumnStringGet(t *testing.T) {
	expectedValues := []string{"a", "b", "c", "d"}
	column := ColumnString{expectedValues}
	for i, expectedValue := range expectedValues {
		if currentValue := column.Get(i); currentValue != expectedValue {
			t.Errorf("Wrong value returned, got %v, want %v", currentValue, expectedValue)
		}
	}
}

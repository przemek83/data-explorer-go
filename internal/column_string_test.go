package internal

import (
	"testing"
)

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

func TestNumericColumnAppend(t *testing.T) {
	values := []string{"a", "b", "c", "d"}
	column := ColumnString{values}
	expectedValue := "j"
	column.Append(expectedValue)
	if currentValue := column.Get(len(values)); currentValue != expectedValue {
		t.Errorf("Wrong value returned, got %v, want %v", currentValue, expectedValue)
	}
}

func TestColumnStringGetSize(t *testing.T) {
	tests := []struct {
		name   string
		column *ColumnString
		want   int
	}{
		{"Empty.", &ColumnString{}, 0},
		{"Not empty.", &ColumnString{[]string{"a", "b", "c", "d"}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.column.GetSize(); got != tt.want {
				t.Errorf("ColumnString.GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

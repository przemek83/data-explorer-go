package internal

import (
	"testing"
)

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

func TestColumnNumericGetSize(t *testing.T) {
	tests := []struct {
		name   string
		column *ColumnNumeric
		want   int
	}{
		{"Empty.", &ColumnNumeric{}, 0},
		{"Not empty.", &ColumnNumeric{[]int{1, 2, 3, 4}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.column.GetSize(); got != tt.want {
				t.Errorf("ColumnNumeric.GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

package internal

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const validInputString = `first_name;age;movie_name;score
string;integer;string;integer
tim;26;inception;8
tim;26;pulp_fiction;8
tamas;44;inception;7
tamas;44;pulp_fiction;4
dave;0;inception;8
dave;0;ender's_game;8`

const emptyDataInput = ``

const inputWithWrongColumnTypeName = `bla;bla;bla;bla
string;integer;bla;integer
tim;26;inception;8`

const inputWithWrongColumnCount = `bla;bla;bla;bla
string;integer;integer
tim;26;inception;8`

const inputWithoutData = `bla;bla;bla
string;integer;integer`

func TestLoad(t *testing.T) {
	tests := []struct {
		name string
		data string
		want bool
	}{
		{"Valid input string", validInputString, true},
		{"Input without data", inputWithoutData, true},
		{"Empty data input", emptyDataInput, false},
		{"Input with wrong column type name", inputWithWrongColumnTypeName, false},
		{"Input with wrong column count", inputWithWrongColumnCount, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.data))
			loader := NewFileDataLoader(reader)
			if got := loader.Load(); got != tt.want {
				t.Errorf("FileDataLoader.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHeaders(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(validInputString))
	loader := NewFileDataLoader(reader)
	loader.Load()
	currentHeaders := loader.GetHeaders()
	expectedHeaders := []string{"first_name", "age", "movie_name", "score"}
	if !reflect.DeepEqual(currentHeaders, expectedHeaders) {
		t.Errorf("Wrong headers. Expected %v, got %v", expectedHeaders, currentHeaders)
	}
}

func TestGetColumnTypes(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(validInputString))
	loader := NewFileDataLoader(reader)
	loader.Load()
	currentColumnTypes := loader.GetColumnTypes()
	expectedColumnTypes := []ColumnType{StringColumn, NumericColumn, StringColumn, NumericColumn}
	if !reflect.DeepEqual(currentColumnTypes, expectedColumnTypes) {
		t.Errorf("Wrong column types. Expected %v, got %v", expectedColumnTypes, currentColumnTypes)
	}
}

func (column *ColumnNumeric) setData(data []int) {
	column.data = data
}

func (column *ColumnString) setData(data []string) {
	column.data = data
}

func prepareData() []Column {
	column1 := ColumnString{}
	column1.setData([]string{"tim", "tim", "tamas", "tamas", "dave", "dave"})
	var columns []Column
	columns = append(columns, &column1)

	column2 := ColumnNumeric{}
	column2.setData([]int{26, 26, 44, 44, 0, 0})
	columns = append(columns, &column2)

	column3 := ColumnString{}
	column3.setData([]string{"inception", "pulp_fiction", "inception", "pulp_fiction", "inception", "ender's_game"})
	columns = append(columns, &column3)

	column4 := ColumnNumeric{}
	column4.setData([]int{8, 8, 7, 4, 8, 8})
	columns = append(columns, &column4)

	return columns
}

func TestGetData(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(validInputString))
	loader := NewFileDataLoader(reader)
	loader.Load()
	currentData := loader.GetData()
	expectedData := prepareData()
	if len(currentData) != len(expectedData) {
		t.Errorf("Wrong data. Expected number of columns %d, got %d", len(expectedData), len(currentData))
	}

	for i, expectedColumn := range expectedData {
		currentColumn := currentData[i]
		if !reflect.DeepEqual(currentColumn, expectedColumn) {
			t.Errorf("Wrong data in column \"%s\". Expected %v, got %v", loader.GetHeaders()[i], expectedColumn, currentColumn)
		}
	}
}

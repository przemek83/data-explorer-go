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
	testCases := []struct {
		data           string
		expectedResult bool
	}{
		{validInputString, true},
		{inputWithoutData, true},
		{emptyDataInput, false},
		{inputWithWrongColumnTypeName, false},
		{inputWithWrongColumnCount, false},
	}

	for _, tc := range testCases {
		loader := FileDataLoader{}
		reader := bufio.NewReader(strings.NewReader(tc.data))
		ok := loader.Load(reader)
		if ok != tc.expectedResult {
			t.Errorf("Wrong loading result. Expected %t, got %t, data \"%s\"", tc.expectedResult, ok, tc.data)
		}
	}
}

func TestGetHeaders(t *testing.T) {
	loader := FileDataLoader{}
	reader := bufio.NewReader(strings.NewReader(validInputString))
	loader.Load(reader)
	currentHeaders := loader.GetHeaders()
	expectedHeaders := []string{"first_name", "age", "movie_name", "score"}
	if !reflect.DeepEqual(currentHeaders, expectedHeaders) {
		t.Errorf("Wrong headers. Expected %v, got %v", expectedHeaders, currentHeaders)
	}
}

func TestGetColumnTypes(t *testing.T) {
	loader := FileDataLoader{}
	reader := bufio.NewReader(strings.NewReader(validInputString))
	loader.Load(reader)
	currentColumnTypes := loader.GetColumnTypes()
	expectedColumnTypes := []ColumnType{String, Integer, String, Integer}
	if !reflect.DeepEqual(currentColumnTypes, expectedColumnTypes) {
		t.Errorf("Wrong column types. Expected %v, got %v", expectedColumnTypes, currentColumnTypes)
	}
}

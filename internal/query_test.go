package internal

import (
	"reflect"
	"testing"
)

func TestMakeQuery(t *testing.T) {
	testCases := []struct {
		name     string
		headers  []string
		args     []string
		errOccur bool
		query    Query
	}{{"Proper number of args, avg operation, proper columns.", []string{"a", "b", "c"}, []string{"avg", "a", "b"}, false, Query{Average, 0, 1}}}

	for _, tc := range testCases {
		loader := newFileDataLoaderStub(tc.headers, []ColumnType{}, []Column{}, true)
		_, dataset := MakeDataset(loader)
		query, err := MakeQuery(tc.args, &dataset)
		if tc.errOccur && err == nil {
			t.Errorf("Error unexpectedly not raised.")
		}
		if !tc.errOccur && err != nil {
			t.Errorf("Error unexpectedly raised.")
		}
		if !reflect.DeepEqual(query, tc.query) {
			t.Errorf("Queries not equal. Expected %v, got %v.", tc.query, query)
		}
	}
}

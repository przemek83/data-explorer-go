package internal

import (
	"reflect"
	"testing"
)

func TestMakeQuery(t *testing.T) {
	testCases := []struct {
		name      string
		headers   []string
		args      []string
		wantError bool
		wantQuery Query
	}{{"Proper number of args, avg operation, proper columns.", []string{"a", "b", "c"}, []string{"avg", "a", "b"}, false, Query{Average, 0, 1}}}

	for _, tc := range testCases {
		loader := newFileDataLoaderStub(tc.headers, []ColumnType{}, []Column{}, true)
		_, dataset := MakeDataset(loader)
		gotQuery, gotErr := MakeQuery(tc.args, &dataset)
		if tc.wantError && gotErr == nil {
			t.Errorf("Error unexpectedly not raised.")
		}
		if !tc.wantError && gotErr != nil {
			t.Errorf("Error unexpectedly raised.")
		}
		if !reflect.DeepEqual(gotQuery, tc.wantQuery) {
			t.Errorf("Queries not equal. Expected %v, got %v.", tc.wantQuery, gotQuery)
		}
	}
}

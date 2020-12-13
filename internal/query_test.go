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
	}{
		{
			"Proper number of args, avg operation, proper columns.",
			[]string{"a", "b", "c"},
			[]string{"avg", "a", "b"},
			false,
			Query{Average, 0, 1},
		},
		{
			"Proper number of args, min operation, proper columns.",
			[]string{"a", "b", "c"},
			[]string{"min", "a", "b"},
			false,
			Query{Minimum, 0, 1},
		},
		{
			"Proper number of args, max operation, proper columns.",
			[]string{"a", "b", "c"},
			[]string{"max", "a", "b"},
			false,
			Query{Maximum, 0, 1},
		},
		{
			"Proper number of args, wrong operation, proper columns.",
			[]string{"a", "b", "c"},
			[]string{"ble", "a", "b"},
			true,
			Query{},
		},
		{
			"Args list too long, avg operation, proper columns.",
			[]string{"a", "b", "c"},
			[]string{"avg", "a", "b", "d"},
			true,
			Query{},
		},
		{
			"Args list too short, avg operation, proper columns.",
			[]string{"a", "b", "c"},
			[]string{"avg", "a"},
			true,
			Query{},
		},
		{
			"Proper number of args, avg operation, wrong aggregate column.",
			[]string{"a", "b", "c"},
			[]string{"avg", "d", "a"},
			true,
			Query{},
		},
		{
			"Proper number of args, avg operation, wrong grouping column.",
			[]string{"a", "b", "c"},
			[]string{"avg", "a", "d"},
			true,
			Query{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			loader := newFileDataLoaderStub(tc.headers, []ColumnType{}, []Column{}, true)
			dataset, _ := MakeDataset(loader)
			gotQuery, gotErr := MakeQuery(tc.args, &dataset)
			if tc.wantError && gotErr == nil {
				t.Errorf("Error unexpectedly not raised.")
			}
			if !tc.wantError && gotErr != nil {
				t.Errorf("Error unexpectedly raised.")
			}
			if gotErr == nil && !reflect.DeepEqual(gotQuery, tc.wantQuery) {
				t.Errorf("Queries not equal. Expected %v, got %v.", tc.wantQuery, gotQuery)
			}
		})
	}
}

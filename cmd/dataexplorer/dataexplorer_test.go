package main

import (
	"reflect"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		wantOK bool
		want   []string
	}{
		{"5 args passed, four returned.", []string{"binary", "One", "Two", "Three", "Four"}, true, []string{"One", "Two", "Three", "Four"}},
		{"No params passed", []string{}, false, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ok, got := parseArgs(tt.args); !reflect.DeepEqual(got, tt.want) || ok != tt.wantOK {
				t.Errorf("parseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

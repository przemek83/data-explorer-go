package internal

import (
	"testing"
)

type dataLoaderStub struct {
	Headers     []string
	ColumnTypes []ColumnType
	Data        []Column
	LoadOK      bool
}

func newFileDataLoaderStub(headers []string, columnTypes []ColumnType, columns []Column, loadOK bool) *dataLoaderStub {
	return &dataLoaderStub{headers, columnTypes, columns, loadOK}
}

func (loader *dataLoaderStub) Load() bool {
	return loader.LoadOK
}

func (loader *dataLoaderStub) GetHeaders() []string {
	return loader.Headers
}

func (loader *dataLoaderStub) GetColumnTypes() []ColumnType {
	return loader.ColumnTypes
}

func (loader *dataLoaderStub) GetData() []Column {
	return loader.Data
}

func TestMakeDataset(t *testing.T) {
	tests := []struct {
		name       string
		dataLoader DataLoader
		want       bool
	}{
		{"Loading OK", newFileDataLoaderStub([]string{}, []ColumnType{}, []Column{}, true), true},
		{"Loading not OK", newFileDataLoaderStub([]string{}, []ColumnType{}, []Column{}, false), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := MakeDataset(tt.dataLoader); got != tt.want {
				t.Errorf("MakeDataset returned %v, want %v", got, tt.want)
			}
		})
	}
}

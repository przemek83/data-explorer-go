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

func TestColumnNameToIDPositive(t *testing.T) {
	headers := []string{"a", "b", "c"}
	loader := newFileDataLoaderStub(headers, []ColumnType{}, []Column{}, true)
	_, dataset := MakeDataset(loader)
	for i, header := range headers {
		ok, gotID := dataset.ColumnNameToID(header)
		if !ok {
			t.Errorf("Error finding header named %s", header)
		}
		if gotID != i {
			t.Errorf("Wrong id found for header named %s. Got %d, expected %d", header, gotID, i)
		}

	}
}

func TestColumnNameToIDNegative(t *testing.T) {
	headers := []string{"a", "b", "c"}
	loader := newFileDataLoaderStub(headers, []ColumnType{}, []Column{}, true)
	_, dataset := MakeDataset(loader)
	wrongHeaderName := "d"
	ok, _ := dataset.ColumnNameToID(wrongHeaderName)
	if ok {
		t.Errorf("Header named %s found unexpectedly", wrongHeaderName)
	}
}

func TestColumnIDToName(t *testing.T) {
	headers := []string{"a", "b", "c"}
	loader := newFileDataLoaderStub(headers, []ColumnType{}, []Column{}, true)
	_, dataset := MakeDataset(loader)
	for i, header := range headers {
		ok, gotName := dataset.ColumnIDToName(i)
		if !ok {
			t.Errorf("Error finding header with id %d", i)
		}
		if gotName != header {
			t.Errorf("Wrong header name found for given id = %d. Got %s, expected %s", i, gotName, header)
		}

	}
}

func TestColumnIDToNameWrongID(t *testing.T) {
	headers := []string{"a", "b", "c"}

	tests := []struct {
		name    string
		wrongID int
	}{
		{"Id too big", len(headers)},
		{"Id negative", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loader := newFileDataLoaderStub(headers, []ColumnType{}, []Column{}, true)
			_, dataset := MakeDataset(loader)
			ok, _ := dataset.ColumnIDToName(tt.wrongID)
			if ok {
				t.Errorf("Header with index %d found unexpectedly", tt.wrongID)
			}
		})
	}
}

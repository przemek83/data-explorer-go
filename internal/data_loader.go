package internal

// DataLoader : interface for data loading.
type DataLoader interface {
	Load() bool
	GetHeaders() []string
	GetColumnTypes() []ColumnType
	GetData() []Column
}

package internal

// ColumnType - Column type.
type ColumnType int

const (
	// NumericColumn - NumericColumn column type.
	NumericColumn ColumnType = iota
	// StringColumn - StringColumn column type.
	StringColumn
	// Unknown - Unknown column type.
	Unknown
)

func (d ColumnType) String() string {
	return [...]string{"integer", "string", ""}[d]
}

// ColumnTypeFromString - return ColumnType for given string.
func ColumnTypeFromString(columnTypeString string) ColumnType {
	switch columnTypeString {
	case NumericColumn.String():
		return NumericColumn
	case StringColumn.String():
		return StringColumn
	}
	return Unknown
}

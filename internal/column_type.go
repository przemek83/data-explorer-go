package internal

// ColumnType - Column type.
type ColumnType int

const (
	// Integer - Integer column type.
	Integer ColumnType = iota
	// String - String column type.
	String
	// Unknown - Unknown column type.
	Unknown
)

func (d ColumnType) String() string {
	return [...]string{"integer", "string", ""}[d]
}

// ColumnTypeFromString - return ColumnType for given string.
func ColumnTypeFromString(columnTypeString string) ColumnType {
	switch columnTypeString {
	case Integer.String():
		return Integer
	case String.String():
		return String
	}
	return Unknown
}

package internal

// Operation - Operation type.
type Operation int

const (
	// Average - Calculate average.
	Average Operation = iota
	// Minimum - Calculate minimum.
	Minimum
	// Maximum - Calculate maximum.
	Maximum
	// UnknownOperation - Unknown operation.
	UnknownOperation
)

func (d Operation) String() string {
	return [...]string{"avg", "min", "max"}[d]
}

// OperationFromString - return Operation for given string.
func OperationFromString(operationString string) Operation {
	switch operationString {
	case Average.String():
		return Average
	case Minimum.String():
		return Minimum
	case Maximum.String():
		return Maximum
	}
	return UnknownOperation
}

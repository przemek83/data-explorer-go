package internal

// OperationType - Operation type.
type OperationType int

const (
	// Average - Calculate average.
	Average OperationType = iota
	// Minimum - Calculate minimum.
	Minimum
	// Maximum - Calculate maximum.
	Maximum
	// UnknownOperation - Unknown operation.
	UnknownOperation
)

func (d OperationType) String() string {
	return [...]string{"avg", "min", "max"}[d]
}

// OperationTypeFromString - return OperationType for given string.
func OperationTypeFromString(operationTypeString string) OperationType {
	switch operationTypeString {
	case Average.String():
		return Average
	case Minimum.String():
		return Minimum
	case Maximum.String():
		return Maximum
	}
	return UnknownOperation
}

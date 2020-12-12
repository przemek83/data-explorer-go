package internal

import "errors"

// Query - query to be execututed.
type Query struct {
	operation         Operation
	aggregateColumnID int
	groupingColumnID  int
}

// MakeQuery  - create query.
func MakeQuery(args []string, dataset *Dataset) (Query, error) {
	if len(args) != 3 {
		return Query{}, errors.New("Wrong number of args")
	}
	operation := OperationFromString(args[0])
	if operation == UnknownOperation {
		return Query{}, errors.New("Unknown operation")
	}
	ok, aggregateColumnID := dataset.ColumnNameToID(args[1])
	if !ok {
		return Query{}, errors.New("Unknown column " + args[1])
	}
	ok, groupingColumnID := dataset.ColumnNameToID(args[2])
	if !ok {
		return Query{}, errors.New("Unknown column " + args[2])
	}
	return Query{operation, aggregateColumnID, groupingColumnID}, nil
}

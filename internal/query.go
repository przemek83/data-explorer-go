package internal

import "errors"

// Query - query to be execututed.
type Query struct {
	operation         Operation
	aggregateColumnID int
	groupingColumnID  int
}

// MakeQuery  - create query.
func MakeQuery(args []string) (Query, error) {
	return Query{}, errors.New("Not implemented")
}

// ColumnsAreValid - check if columns on input are columns in dataset.
func (query *Query) ColumnsAreValid(dataset *Dataset) error {
	return errors.New("Not implemented")
}

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
	return Query{}, errors.New("Not implemented")
}

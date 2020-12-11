package internal

// Query - query to be execututed.
type Query struct {
	operation         Operation
	aggregateColumnID int
	groupingColumnID  int
}

// MakeQuery  - create query.
func MakeQuery(args []string) (bool, Query) {
	return false, Query{}
}

// ColumnsAreValid - check if columns on input are columns in dataset.
func (query *Query) ColumnsAreValid(dataset *Dataset) bool {
	return false
}

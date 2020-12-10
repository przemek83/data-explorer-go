package internal

// Query - query to be execututed.
type Query struct {
	operation         Operation
	aggregateColumnID int
	groupingColumnID  int
}

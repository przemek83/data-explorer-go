package internal

import (
	"errors"
	"fmt"
)

// Calculator - Calculates results according to query.
type Calculator struct {
	dataset Dataset
}

// MakeCalculator - Create  Calculator.
func MakeCalculator(dataset Dataset) Calculator {
	return Calculator{dataset}
}

// Execute - Execute given query on dataset.
func (calculator *Calculator) Execute(query Query) (map[string]float32, error) {
	if err := calculator.checkColumnTypes(query); err != nil {
		return map[string]float32{}, err
	}

	aggregationColumn := calculator.getAggregationColumn(query)
	groupingColumn := calculator.getGroupingColumn(query)

	switch query.operation {
	case Average:
		return calculator.computeAvg(aggregationColumn, groupingColumn), nil
	case Maximum:
		return calculator.computeMax(aggregationColumn, groupingColumn), nil
	case Minimum:
		return calculator.computeMin(aggregationColumn, groupingColumn), nil
	}
	return map[string]float32{}, errors.New("Operation unknown")
}

func (calculator *Calculator) getAggregationColumn(query Query) *ColumnNumeric {
	_, aggregationColumn := calculator.dataset.GetData(query.aggregateColumnID)
	columnNumeric, _ := aggregationColumn.(*ColumnNumeric)
	return columnNumeric
}

func (calculator *Calculator) getGroupingColumn(query Query) *ColumnString {
	_, groupingColumn := calculator.dataset.GetData(query.groupingColumnID)
	columnString, _ := groupingColumn.(*ColumnString)
	return columnString
}

func (calculator *Calculator) checkColumnTypes(query Query) error {
	ok, aggregationColumnType := calculator.dataset.GetColumnType(query.aggregateColumnID)
	if !ok || aggregationColumnType != NumericColumn {
		errorString := fmt.Sprintf("Aggregate column with id %d not found or not numeric type", query.aggregateColumnID)
		return errors.New(errorString)
	}
	ok, groupingColumnType := calculator.dataset.GetColumnType(query.groupingColumnID)
	if !ok || groupingColumnType != StringColumn {
		errorString := fmt.Sprintf("Grouping column with id %d not found or not string type", query.groupingColumnID)
		return errors.New(errorString)
	}
	return nil
}

type pair struct {
	count, sum float32
}

func (calculator *Calculator) computeAvg(aggregationColumn *ColumnNumeric, groupingColumn *ColumnString) map[string]float32 {
	sums := calculator.computeSum(aggregationColumn, groupingColumn)
	results := map[string]float32{}
	for grouping, values := range sums {
		results[grouping] = values.sum / values.count
	}
	return results
}

func (calculator *Calculator) computeSum(aggregationColumn *ColumnNumeric, groupingColumn *ColumnString) map[string]pair {
	results := map[string]pair{}
	for i := 0; i < aggregationColumn.GetSize(); i++ {
		value := float32(aggregationColumn.Get(i))
		key := groupingColumn.Get(i)
		currentValue, exists := results[key]
		if exists {
			results[key] = pair{currentValue.count + 1, currentValue.sum + value}
		} else {
			results[key] = pair{1, value}
		}
	}
	return results
}

func (calculator *Calculator) computeMax(aggregationColumn *ColumnNumeric, groupingColumn *ColumnString) map[string]float32 {
	greater := func(i, j float32) bool {
		return i > j
	}
	return calculator.computeExtreme(aggregationColumn, groupingColumn, greater)
}

func (calculator *Calculator) computeMin(aggregationColumn *ColumnNumeric, groupingColumn *ColumnString) map[string]float32 {
	lower := func(i, j float32) bool {
		return i < j
	}
	return calculator.computeExtreme(aggregationColumn, groupingColumn, lower)
}

func (calculator *Calculator) computeExtreme(aggregationColumn *ColumnNumeric,
	groupingColumn *ColumnString,
	condition func(left, right float32) bool) map[string]float32 {
	results := map[string]float32{}
	for i := 0; i < aggregationColumn.GetSize(); i++ {
		value := float32(aggregationColumn.Get(i))
		key := groupingColumn.Get(i)
		currentValue, exists := results[key]
		if !exists {
			results[key] = value
			continue
		}
		if condition(value, currentValue) {
			results[key] = value
		}
	}
	return results
}

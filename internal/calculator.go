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

	aggregationColumn, err := calculator.getAggregationColumn(query)
	if err != nil {
		return map[string]float32{}, err
	}

	groupingColumn, err := calculator.getGroupingColumn(query)
	if err != nil {
		return map[string]float32{}, err
	}

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

func (calculator *Calculator) getAggregationColumn(query Query) (*ColumnNumeric, error) {
	ok, aggregationColumn := calculator.dataset.GetData(query.aggregateColumnID)
	if !ok {
		errorString := fmt.Sprintf("Aggregate column with id %d not found", query.aggregateColumnID)
		return nil, errors.New(errorString)
	}
	columnNumeric, valid := aggregationColumn.(*ColumnNumeric)
	if !valid {
		return nil, errors.New("Cast to ColumnNumeric failed")
	}
	return columnNumeric, nil
}

func (calculator *Calculator) getGroupingColumn(query Query) (*ColumnString, error) {
	ok, groupingColumn := calculator.dataset.GetData(query.groupingColumnID)
	if !ok {
		errorString := fmt.Sprintf("Grouping column with id %d not found", query.groupingColumnID)
		return nil, errors.New(errorString)
	}
	columnString, valid := groupingColumn.(*ColumnString)
	if !valid {
		return nil, errors.New("Cast to ColumnString failed")
	}
	return columnString, nil
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

func (calculator *Calculator) computeAvg(aggregationColumn *ColumnNumeric, groupingColumn *ColumnString) map[string]float32 {
	results := map[string]float32{}
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

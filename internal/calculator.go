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
	ok, aggregationColumn := calculator.dataset.GetData(query.aggregateColumnID)
	if !ok {
		errorString := fmt.Sprintf("Aggregate column with id %d not found", query.aggregateColumnID)
		return map[string]float32{}, errors.New(errorString)
	}
	ok, groupingColumn := calculator.dataset.GetData(query.groupingColumnID)
	if !ok {
		errorString := fmt.Sprintf("Grouping column with id %d not found", query.groupingColumnID)
		return map[string]float32{}, errors.New(errorString)
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

func (calculator *Calculator) computeAvg(aggregationColumn Column, groupingColumn Column) map[string]float32 {
	return map[string]float32{}
}

func (calculator *Calculator) computeMax(aggregationColumn Column, groupingColumn Column) map[string]float32 {
	greater := func(i, j int) bool {
		return i > j
	}
	return calculator.computeExtreme(aggregationColumn, groupingColumn, greater)
}

func (calculator *Calculator) computeMin(aggregationColumn Column, groupingColumn Column) map[string]float32 {
	lower := func(i, j int) bool {
		return i < j
	}
	return calculator.computeExtreme(aggregationColumn, groupingColumn, lower)
}

func (calculator *Calculator) computeExtreme(aggregationColumn Column,
	groupingColumn Column,
	condition func(left, right int) bool) map[string]float32 {
	return map[string]float32{}
}

package internal

// Calculator - Calculates results according to query.
type Calculator struct {
	dataset Dataset
}

// MakeCalculator - Create  Calculator.
func MakeCalculator(dataset Dataset) Calculator {
	return Calculator{dataset}
}

// Execute - Execute given query on dataset.
func (calculator *Calculator) Execute(query Query) map[string]float32 {
	switch query.operation {
	case Average:
		return map[string]float32{}
	case Maximum:
		return map[string]float32{}
	case Minimum:
		return map[string]float32{}
	}
	return map[string]float32{}
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

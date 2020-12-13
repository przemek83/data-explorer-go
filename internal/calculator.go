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
func (Calculator *Dataset) Execute(query Query) map[string]float32 {
	return map[string]float32{}
}

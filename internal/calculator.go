package internal

// Calculator - Calculates results according to query.
type Calculator struct {
	dataset Dataset
}

// MakeCalculator - Create  Calculator.
func MakeCalculator(dataset Dataset) Calculator {
	return Calculator{dataset}
}

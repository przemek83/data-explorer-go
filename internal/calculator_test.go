package internal

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestCalculatorExecute(t *testing.T) {
	const firstName int = 0
	const age int = 1
	const movieName int = 2
	const score int = 3
	tests := []struct {
		name  string
		query Query
		want  map[string]float32
	}{
		{
			"Max age grouped by movie name",
			Query{Maximum, age, movieName},
			map[string]float32{"inception": 44, "pulp_fiction": 44, "ender's_game": 0},
		},
		{
			"Max score grouped by movie name",
			Query{Maximum, score, movieName},
			map[string]float32{"inception": 8, "pulp_fiction": 8, "ender's_game": 8},
		},
		{
			"Max score grouped by first name",
			Query{Maximum, score, firstName},
			map[string]float32{"tim": 8, "tamas": 7, "dave": 8},
		},
		{
			"Min age grouped by movie name",
			Query{Minimum, age, movieName},
			map[string]float32{"inception": 0, "pulp_fiction": 26, "ender's_game": 0},
		},
		{
			"Min score grouped by movie name",
			Query{Minimum, score, movieName},
			map[string]float32{"inception": 7, "pulp_fiction": 4, "ender's_game": 8},
		},
		{
			"Min score grouped by first name",
			Query{Minimum, score, firstName},
			map[string]float32{"tim": 8, "tamas": 4, "dave": 8},
		},
		{
			"Average age grouped by movie name",
			Query{Average, age, movieName},
			map[string]float32{"inception": 70 / 3, "pulp_fiction": 35, "ender's_game": 0},
		},
		{
			"Average score grouped by movie name",
			Query{Average, score, movieName},
			map[string]float32{"inception": 7.67, "pulp_fiction": 6, "ender's_game": 8},
		},
		{
			"Average score grouped by first name",
			Query{Average, score, firstName},
			map[string]float32{"tim": 8, "tamas": 5.5, "dave": 8},
		},
	}
	reader := bufio.NewReader(strings.NewReader(validInputString))
	loader := NewFileDataLoader(reader)
	loader.Load()
	dataset, _ := MakeDataset(loader)
	calculator := MakeCalculator(dataset)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := calculator.Execute(tt.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculator.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

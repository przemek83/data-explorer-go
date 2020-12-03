package internal

type ColumnNumeric struct {
	Column
	data []int
}

func (column *ColumnNumeric) Append(value int) {
	column.data = append(column.data, value)
}

func (column *ColumnNumeric) Get(index int) int {
	return column.data[index]
}

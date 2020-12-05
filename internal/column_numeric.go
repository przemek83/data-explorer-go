package internal

type ColumnNumeric struct {
	data []int
}

func (column *ColumnNumeric) GetType() ColumnType {
	return NumericColumn
}

func (column *ColumnNumeric) Append(value int) {
	column.data = append(column.data, value)
}

func (column *ColumnNumeric) Get(index int) int {
	return column.data[index]
}

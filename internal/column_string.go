package internal

type ColumnString struct {
	Column
	data []string
}

func (column *ColumnString) Append(value string) {
	column.data = append(column.data, value)
}

func (column *ColumnString) Get(index int) string {
	return column.data[index]
}

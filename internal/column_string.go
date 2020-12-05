package internal

type ColumnString struct {
	data []string
}

func (column *ColumnString) GetType() ColumnType {
	return StringColumn
}

func (column *ColumnString) Append(value string) {
	column.data = append(column.data, value)
}

func (column *ColumnString) Get(index int) string {
	return column.data[index]
}

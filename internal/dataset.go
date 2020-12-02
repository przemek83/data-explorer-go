package internal

type Dataset struct {
	loader      FileDataLoader
	headers     []string
	columnTypes []ColumnType
}

func MakeDataset(loader FileDataLoader) Dataset {
	return Dataset{loader, []string{}, []ColumnType{}}
}

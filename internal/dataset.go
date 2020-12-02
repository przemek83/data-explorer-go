package internal

type Dataset struct {
	loader      FileDataLoader
	headers     []string
	columnTypes []ColumnType
}

func MakeDataset(loader FileDataLoader) Dataset {
	return Dataset{loader, []string{}, []ColumnType{}}
}

func (dataset *Dataset) initialize() bool {
	return false
}

func (dataset *Dataset) ColumnNameToID() (bool, int) {
	return false, -1
}

func (dataset *Dataset) ColumnIDToName() (bool, string) {
	return false, ""
}

func (dataset *Dataset) GetColumnType() (bool, ColumnType) {
	return false, Unknown
}

func (dataset *Dataset) GetData(columnID int) (bool, []string) {
	return false, []string{}
}

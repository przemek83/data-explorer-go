package internal

type Dataset struct {
	headers     []string
	columnTypes []ColumnType
	data        []Column
}

func MakeDataset(loader DataLoader) (bool, Dataset) {
	if !loader.Load() {
		return false, Dataset{}
	}
	return true, Dataset{loader.GetHeaders(), loader.GetColumnTypes(), loader.GetData()}
}

func (dataset *Dataset) ColumnNameToID(name string) (bool, int) {
	return false, -1
}

func (dataset *Dataset) ColumnIDToName(id int) (bool, string) {
	return false, ""
}

func (dataset *Dataset) GetColumnType(id int) (bool, ColumnType) {
	return false, Unknown
}

func (dataset *Dataset) GetData(id int) (bool, []Column) {
	return false, dataset.data
}

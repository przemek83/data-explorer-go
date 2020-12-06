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

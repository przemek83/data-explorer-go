package internal

type Column interface {
	Append(value interface{})
	Get(index int) interface{}
}

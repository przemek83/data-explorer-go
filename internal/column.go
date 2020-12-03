package internal

type Column interface {
	append()
	get(index int) interface{}
}

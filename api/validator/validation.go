package validator

type IError struct {
	Field string
	Tag   string
	Value string
}

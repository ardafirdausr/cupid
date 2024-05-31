package validator

type Validator interface {
	ValidateStruct(data interface{}) (map[string]string, error)
}

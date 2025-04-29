package main

type Validator interface {
	Validate() error
}

func ValidateModel(s Validator) error {
	return s.Validate()
}

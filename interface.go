package prompt

// Validator is an interface that is used to accept values
// from user input and return an error if the value is
// not valid.
type Validator interface {
	Validate(value string) error
}

package prompt

// Validator is a func that is used to accept values
// from user input and return an error if the value is
// not valid.
type Validator func(string) error

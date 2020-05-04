package prompt

// InputOptions are passed to the individual questions to
// set defaults and validators of the input.
type InputOptions struct {
	Default   string
	Validator Validator
}

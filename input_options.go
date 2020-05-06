package prompt

// InputOptions are passed to the individual questions to
// set defaults and validators of the input.
type InputOptions struct {
	Default   string
	Validator Validator
}

// SelectOptions are used on the select prompts and allow
// setting a default index option which is non-zero
// indexed.
type SelectOptions struct {
	Default   int
	Validator Validator
}

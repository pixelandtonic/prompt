package prompt

// Options are set globally for the prompt package to reduce
// common prompt tasks like appending space after the prompt,
// appending questions marks to ask prompts, and showing
// the default value in a prompt.
type Options struct {
	// AppendQuestionMarksOnAsk will append the
	// question marks on .Ask input prompts
	// so you don't need to add it to all
	// prompts.
	AppendQuestionMarksOnAsk bool
	// AppendSpace will automatically add
	// a space after prompts to avoid
	// adding strings like "question "
	AppendSpace bool
	// ShowDefaultInPrompt is used to embed
	// the "default" for an input in quotes
	// (e.g. (default is 42)
	ShowDefaultInPrompt bool
}

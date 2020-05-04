package prompt

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Prompt is a struct that contains the reader, writer,
// and options that are applied to all prompts.
type Prompt struct {
	Reader io.Reader
	Writer io.Writer
	*Options
}

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

// InputOptions are passed to the individual questions to
// set defaults and validators of the input.
type InputOptions struct {
	Default   string
	Validator Validator
}

// Ask is used to gather input from a user in the form of a question.
// by default, it will add a ? to the provided question.
func (p *Prompt) Ask(text string, opts *InputOptions) (string, error) {
	format := "%s"
	if p.AppendQuestionMarksOnAsk == true {
		format = format + "?"
	}
	if p.ShowDefaultInPrompt && opts.Default != "" {
		format = format + " [" + opts.Default + "]"
	}
	if p.AppendSpace == true {
		format = format + " "
	}

	fmt.Fprintf(p.Writer, format, text)

	rdr := bufio.NewReader(p.Reader)
	resp, err := rdr.ReadString('\n')
	if err != nil {
		return "", err
	}

	input := strings.TrimSpace(resp)

	if input == "" && opts.Default == "" {
		return "", errors.New("no value provided")
	}

	if input == "" && opts.Default != "" {
		return opts.Default, nil
	}

	return resp, nil
}

func (p *Prompt) Confirm(text string, opts *InputOptions) (bool, error) {
	format := "%s"
	if p.AppendQuestionMarksOnAsk == true {
		format = format + "?"
	}
	if p.ShowDefaultInPrompt && opts.Default != "" {
		format = format + " [" + opts.Default + "]"
	}
	if p.AppendSpace == true {
		format = format + " "
	}

	fmt.Fprintf(p.Writer, format, text)

	rdr := bufio.NewReader(p.Reader)
	resp, err := rdr.ReadString('\n')
	if err != nil {
		return false, err
	}

	input := strings.TrimSpace(resp)

	if input == "" && opts.Default == "" {
		return false, errors.New("no value provided")
	}

	if input == "" && opts.Default != "" {
		input = opts.Default
	}

	if strings.ContainsAny(input, "y") {
		return true, nil
	}

	return false, nil
}

// NewPrompt is used to quickly create a new prompt
// with the default options used on projects.
func NewPrompt() *Prompt {
	return &Prompt{
		Reader: os.Stdin,
		Writer: os.Stdout,
		Options: &Options{
			AppendQuestionMarksOnAsk: true,
			AppendSpace:              true,
			ShowDefaultInPrompt:      true,
		},
	}
}

// NewPromptWithOptions will allow you to override the
// global options for the prompt library.
func NewPromptWithOptions(opts *Options) *Prompt {
	return &Prompt{
		Reader:  os.Stdin,
		Writer:  os.Stdout,
		Options: opts,
	}
}

package prompt

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Prompt is a struct that contains the reader, writer,
// and options that are applied to all prompts.
type Prompt struct {
	Reader io.Reader
	Writer io.Writer
	*Options
}

// Ask is used to gather input from a user in the form of a question.
// by default, it will add a ? to the provided question.
func (p *Prompt) Ask(text string, opts *InputOptions) (string, error) {
	format := p.fmtInputOptions(opts)

	resp, err := p.read(text, format)
	if err != nil {
		return "", err
	}

	input := strings.TrimSpace(resp)

	// no input and no opts
	if input == "" && opts == nil {
		return "", errors.New("no input provided")
	}

	// if opts is not nil and there is a validator
	if opts != nil && opts.Validator != nil {
		if err := opts.Validator(input); err != nil {
			return "", err
		}
	}

	// no input, no opts, and a default is set
	if input == "" && opts != nil && opts.Default != "" {
		return opts.Default, nil
	}

	return input, nil
}

func (p *Prompt) Confirm(text string, opts *InputOptions) (bool, error) {
	format := p.fmtInputOptions(opts)

	resp, err := p.read(text, format)
	if err != nil {
		return false, err
	}

	if resp == "" && opts.Default == "" {
		return false, errors.New("no value provided")
	}

	if opts != nil && opts.Validator != nil {
		if err := opts.Validator(resp); err != nil {
			return false, err
		}
	}

	if resp == "" && opts != nil && opts.Default != "" {
		resp = opts.Default
	}

	if strings.ContainsAny(resp, "yes") {
		return true, nil
	}

	return false, nil
}

// Select will return a prompt with a list of options for a user to select. It will return the
// selected string, the index (as it appears in 0-based index - not the displayed output),
// and an error if something goes wrong.
func (p *Prompt) Select(text string, list []string, opts *SelectOptions) (string, int, error) {
	if len(list) == 0 {
		return "", 0, errors.New("list must be greater than 0")
	}

	format := p.fmtSelectOptions(opts)

	var selectedIndex int
	var selectedText string

	// print the list before the prompt to
	// provide clear options
	for i, l := range list {
		f := "  %d - %s\n"
		if i == len(list)-1 {
			f = "  %d - %s\n"
		}

		_, err := p.Writer.Write([]byte(fmt.Sprintf(f, i+1, l)))
		if err != nil {
			return "", 0, err
		}
	}

	resp, err := p.read(text, format)
	if err != nil {
		return "", 0, err
	}

	if resp == "" && opts.Default != 0 {
		// minus 1 to account for zero index
		resp = strconv.Itoa(opts.Default - 1)
	} else {
		// convert resp to string
		e, err := strconv.Atoi(resp)
		if err != nil {
			return "", 0, err
		}

		// make sure its a valid option before we minus one
		if len(list) < e {
			return "", 0, errors.New("invalid option provided")
		}

		// minus one
		resp = strconv.Itoa(e - 1)
	}

	selectedIndex, err = strconv.Atoi(resp)
	if err != nil {
		return "", 0, err
	}

	selectedText = list[selectedIndex]

	return selectedText, selectedIndex, nil
}

func (p *Prompt) read(text string, format string) (string, error) {
	fmt.Fprintf(p.Writer, format, text)

	rdr := bufio.NewReader(p.Reader)

	resp, err := rdr.ReadString('\n')

	return strings.TrimSpace(resp), err
}

func (p *Prompt) fmtInputOptions(opts *InputOptions) string {
	format := "%s"
	if p.Options != nil && p.AppendQuestionMarksOnAsk == true {
		format = format + "?"
	}
	if p.Options != nil && p.ShowDefaultInPrompt && opts.Default != "" {
		format = format + " [" + opts.Default + "]"
	}
	if p.Options != nil && p.AppendSpace == true {
		format = format + " "
	}

	return format
}

func (p *Prompt) fmtSelectOptions(opts *SelectOptions) string {
	format := "%s"
	if p.Options != nil && p.AppendQuestionMarksOnAsk == true {
		format = format + "?"
	}
	if p.Options != nil && p.ShowDefaultInPrompt && opts.Default != 0 {
		format = format + " [" + strconv.Itoa(opts.Default) + "]"
	}
	if p.Options != nil && p.AppendSpace == true {
		format = format + " "
	}

	return format
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

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
		return "", errors.New("no input provided")
	}

	if input == "" && opts.Default != "" {
		return opts.Default, nil
	}

	if opts.Validator != nil {
		if err := opts.Validator(input); err != nil {
			return "", err
		}
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

	if opts.Validator != nil {
		if err := opts.Validator(input); err != nil {
			return false, err
		}
	}

	if strings.ContainsAny(input, "y") {
		return true, nil
	}

	return false, nil
}

func (p *Prompt) Select(text string, list []string, opts *InputOptions) (string, int, error) {
	if len(list) == 0 {
		return "", 0, errors.New("list must be greater than 0")
	}

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

	var selectedIndex int
	var selectedText string
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

	fmt.Fprintf(p.Writer, format, text)

	rdr := bufio.NewReader(p.Reader)

	resp, err := rdr.ReadString('\n')
	if err != nil {
		return "", 0, err
	}

	// TODO check if the list has that index item

	input := strings.TrimSpace(resp)

	if input == "" && opts.Default != "" {
		input = opts.Default
	}

	selectedIndex, err = strconv.Atoi(input)
	selectedText = list[selectedIndex]

	return selectedText, selectedIndex, nil
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

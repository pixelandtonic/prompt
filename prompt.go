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

var (
	DefaultOptions = &Options{
		AppendQuestionMarksOnAsk: false,
		AppendSpace:              true,
		ShowDefaultInPrompt:      true,
	}
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

	// show me what you're working with
	switch input {
	case "":
		// check the opts
		switch opts {
		case nil:
			// no options and no input means we return an error
			return "", errors.New("no input or default value provided")
		default:
			// check if there is a default to return
			if opts.Default != "" {
				return opts.Default, nil
			}

			if opts.Validator != nil {
				// validate in provided input - even if empty
				if err := opts.Validator(input); err != nil {
					return "", err
				}
			}
		}
	default:
		switch opts {
		case nil:
			// there are no options, so just return the input
			return input, nil
		default:
			if opts.Validator != nil {
				// validate in provided input
				if err := opts.Validator(input); err != nil {
					return "", err
				}
			}
		}
	}

	return input, nil
}

// Confirm is used to prompt a user for a yes/no question. It will always
// return an bool or error if something goes horribly wrong. Like the
// other prompts, it also takes options to provide a default value
// as well as an optional validator to verify input entered.
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

	switch resp {
	case "y":
		return true, nil
	case "yes":
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
	if (p.Options != nil && p.AppendQuestionMarksOnAsk == true) || (opts != nil && opts.AppendQuestionMark == true) {
		format = format + "?"
	}
	if p.Options != nil && p.ShowDefaultInPrompt && opts != nil && opts.Default != "" {
		format = format + " [" + opts.Default + "]"
	}
	if p.Options != nil && p.AppendSpace == true {
		format = format + " "
	}

	return format
}

func (p *Prompt) fmtSelectOptions(opts *SelectOptions) string {
	format := "%s"
	if (p.Options != nil && p.AppendQuestionMarksOnAsk == true) || (opts != nil && opts.AppendQuestionMark == true) {
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
		Reader:  os.Stdin,
		Writer:  os.Stdout,
		Options: DefaultOptions,
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

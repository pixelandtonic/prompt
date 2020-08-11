package prompt

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestPrompt_Ask(t *testing.T) {
	type fields struct {
		Reader  io.Reader
		Writer  io.Writer
		Options *Options
	}
	type args struct {
		text string
		opts *InputOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "complete example will return validator error when input is invalid when provided",
			fields: fields{
				Reader:  bytes.NewReader([]byte("this is wrong\n")),
				Writer:  ioutil.Discard,
				Options: DefaultOptions,
			},
			args: args{
				text: "input will be returned",
				opts: &InputOptions{
					Default: "the-default",
					Validator: func(s string) error {
						if s == "this is wrong" {
							return errors.New("wrong input")
						}
						return nil
					},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "complete example will return provided input when provided",
			fields: fields{
				Reader:  bytes.NewReader([]byte("not-the-default\n")),
				Writer:  ioutil.Discard,
				Options: DefaultOptions,
			},
			args: args{
				text: "input will be returned",
				opts: &InputOptions{
					Default: "the-default",
					Validator: func(s string) error {
						return nil
					},
				},
			},
			want:    "not-the-default",
			wantErr: false,
		},
		{
			name: "complete example will return default when nothing is provided",
			fields: fields{
				Reader:  bytes.NewReader([]byte("\n")),
				Writer:  ioutil.Discard,
				Options: DefaultOptions,
			},
			args: args{
				text: "default will be returned",
				opts: &InputOptions{
					Default: "the-default",
					Validator: func(s string) error {
						return nil
					},
				},
			},
			want:    "the-default",
			wantErr: false,
		},
		{
			name: "provided input will return when no options are present",
			fields: fields{
				Reader:  bytes.NewReader([]byte("some input\n")),
				Writer:  ioutil.Discard,
				Options: DefaultOptions,
			},
			args: args{
				text: "some input will be returned",
				opts: nil,
			},
			want:    "some input",
			wantErr: false,
		},
		{
			name: "empty input will return an error when no options are present",
			fields: fields{
				Reader:  bytes.NewReader([]byte("\n")),
				Writer:  ioutil.Discard,
				Options: DefaultOptions,
			},
			args: args{
				text: "no input error will be returned",
				opts: nil,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prompt{
				Reader:  tt.fields.Reader,
				Writer:  tt.fields.Writer,
				Options: tt.fields.Options,
			}
			got, err := p.Ask(tt.args.text, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrompt_Confirm(t *testing.T) {
	type fields struct {
		Reader  io.Reader
		Writer  io.Writer
		Options *Options
	}
	type args struct {
		text string
		opts *InputOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "returns true when user enters yes",
			fields: fields{
				Reader:  bytes.NewBuffer([]byte("yes\n")),
				Writer:  ioutil.Discard,
				Options: nil,
			},
			args: args{
				text: "Do you agree?",
				opts: nil,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "catches y as true",
			fields: fields{
				Reader:  bytes.NewBuffer([]byte("y\n")),
				Writer:  ioutil.Discard,
				Options: nil,
			},
			args: args{
				text: "Do you agree?",
				opts: nil,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "returns the default",
			fields: fields{
				Reader:  bytes.NewBuffer([]byte("\n")),
				Writer:  ioutil.Discard,
				Options: nil,
			},
			args: args{
				text: "Do you agree?",
				opts: &InputOptions{Default: "no"},
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prompt{
				Reader:  tt.fields.Reader,
				Writer:  tt.fields.Writer,
				Options: tt.fields.Options,
			}
			got, err := p.Confirm(tt.args.text, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Confirm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Confirm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPrompt(t *testing.T) {
	tests := []struct {
		name string
		want *Prompt
	}{
		{
			name: "defaults are defined",
			want: &Prompt{
				Reader: os.Stdin,
				Writer: os.Stdout,
				Options: &Options{
					AppendQuestionMarksOnAsk: false,
					AppendSpace:              true,
					ShowDefaultInPrompt:      true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrompt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPromptWithOptions(t *testing.T) {
	type args struct {
		opts *Options
	}
	tests := []struct {
		name string
		args args
		want *Prompt
	}{
		{
			name: "can override the defaults",
			args: args{opts: &Options{
				AppendSpace: true,
			}},
			want: &Prompt{
				Reader: os.Stdin,
				Writer: os.Stdout,
				Options: &Options{
					AppendSpace: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPromptWithOptions(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPromptWithOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrompt_format(t *testing.T) {
	type fields struct {
		Reader  io.Reader
		Writer  io.Writer
		Options *Options
	}
	type args struct {
		opts *InputOptions
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "all options return the correct fmtInputOptions",
			fields: fields{
				Reader: bytes.NewBuffer([]byte("\n")),
				Writer: ioutil.Discard,
				Options: &Options{
					AppendQuestionMarksOnAsk: true,
					AppendSpace:              true,
					ShowDefaultInPrompt:      true,
				},
			},
			args: args{opts: &InputOptions{
				Default:   "this is my default",
				Validator: nil,
			}},
			want: "%s? [this is my default] ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prompt{
				Reader:  tt.fields.Reader,
				Writer:  tt.fields.Writer,
				Options: tt.fields.Options,
			}
			if got := p.fmtInputOptions(tt.args.opts); got != tt.want {
				t.Errorf("fmtInputOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrompt_Select(t *testing.T) {
	type fields struct {
		Reader  io.Reader
		Writer  io.Writer
		Options *Options
	}
	type args struct {
		text string
		list []string
		opts *SelectOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		{
			name: "out of range selections return an error",
			fields: fields{
				Reader: bytes.NewBuffer([]byte("3\n")),
				Writer: ioutil.Discard,
				Options: &Options{
					AppendQuestionMarksOnAsk: false,
					AppendSpace:              true,
					ShowDefaultInPrompt:      true,
				},
			},
			args: args{
				opts: &SelectOptions{
					Default: 1,
				},
				list: []string{"testing", "select"},
			},
			want:    "",
			want1:   0,
			wantErr: true,
		},
		{
			name: "when providing an option the provided value accounts for zero indexing",
			fields: fields{
				Reader: bytes.NewBuffer([]byte("2\n")),
				Writer: ioutil.Discard,
				Options: &Options{
					AppendQuestionMarksOnAsk: false,
					AppendSpace:              true,
					ShowDefaultInPrompt:      true,
				},
			},
			args: args{
				opts: &SelectOptions{
					Default: 1,
				},
				list: []string{"testing", "select"},
			},
			want:    "select",
			want1:   1,
			wantErr: false,
		},
		{
			name: "returns a default when no options are provided",
			fields: fields{
				Reader: bytes.NewBuffer([]byte("\n")),
				Writer: ioutil.Discard,
				Options: &Options{
					AppendQuestionMarksOnAsk: false,
					AppendSpace:              true,
					ShowDefaultInPrompt:      true,
				},
			},
			args: args{
				opts: &SelectOptions{
					Default: 1,
				},
				list: []string{"testing", "select"},
			},
			want:    "testing",
			want1:   0,
			wantErr: false,
		},
		{
			name: "out of bound options returns an error",
			fields: fields{
				Reader:  bytes.NewBuffer([]byte("4\n")),
				Writer:  ioutil.Discard,
				Options: nil,
			},
			args: args{
				opts: &SelectOptions{
					Default: 1,
				},
				list: []string{"testing", "select"},
			},
			want:    "",
			want1:   0,
			wantErr: true,
		},
		{
			name: "empty lists returns an error",
			args: args{
				text: "some text",
				list: nil,
				opts: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prompt{
				Reader:  tt.fields.Reader,
				Writer:  tt.fields.Writer,
				Options: tt.fields.Options,
			}
			got, got1, err := p.Select(tt.args.text, tt.args.list, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Select() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Select() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
			name: "provided input is returned",
			fields: fields{
				Reader:  bytes.NewBuffer([]byte("output\n")),
				Writer:  ioutil.Discard,
				Options: nil,
			},
			args: args{
				text: "What do you want to see?",
				opts: nil,
			},
			want:    "output",
			wantErr: false,
		},
		{
			name: "defaults are returned when no data is provided",
			fields: fields{
				Reader:  bytes.NewBuffer([]byte("\n")),
				Writer:  ioutil.Discard,
				Options: nil,
			},
			args: args{
				text: "What do you want to see?",
				opts: &InputOptions{
					Default: "my default",
				},
			},
			want:    "my default",
			wantErr: false,
		},
		{
			name: "validator is run when provided",
			fields: fields{
				Reader:  bytes.NewBuffer([]byte("not42\n")),
				Writer:  ioutil.Discard,
				Options: nil,
			},
			args: args{
				text: "What is the meaning of life?",
				opts: &InputOptions{
					Validator: func(s string) error {
						if s != "42" {
							return errors.New("wrong answer")
						}
						return nil
					},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "no input and no default returns an error",
			fields: fields{
				Reader:  bytes.NewBuffer([]byte("\n")),
				Writer:  ioutil.Discard,
				Options: nil,
			},
			args: args{
				text: "This is going to fail",
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
			name: "catches any part of the word yes as true",
			fields: fields{
				Reader:  bytes.NewBuffer([]byte("es\n")),
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
					AppendQuestionMarksOnAsk: true,
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

func TestPrompt_Confirm1(t *testing.T) {
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
		// TODO: Add test cases.
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
			name: "all options return the correct format",
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
			if got := p.format(tt.args.opts); got != tt.want {
				t.Errorf("format() = %v, want %v", got, tt.want)
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
		opts *InputOptions
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
			name: "returns a default when no options are provided",
			fields: fields{
				Reader: bytes.NewBuffer([]byte("\n")),
				Writer: ioutil.Discard,
				Options: &Options{
					AppendQuestionMarksOnAsk: true,
					AppendSpace:              true,
					ShowDefaultInPrompt:      true,
				},
			},
			args: args{
				opts: &InputOptions{
					Default: "1",
				},
				list: []string{"testing", "select"},
			},
			want:    "select",
			want1:   1,
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
				opts: &InputOptions{
					Default: "1",
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

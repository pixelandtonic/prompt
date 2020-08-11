package main

import (
	"fmt"
	"os"

	"github.com/pixelandtonic/prompt"
)

func main() {
	p := prompt.NewPrompt()

	remove, err := p.Confirm(fmt.Sprintf("Are you sure you want to permanently remove the database %q", "somedatabase"), &prompt.InputOptions{
		Default:            "no",
		Validator:          nil,
		AppendQuestionMark: true,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(remove)
}

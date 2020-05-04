package main

import (
	"fmt"
	"log"

	"github.com/pixelandtonic/prompt"
)

func main() {
	// provide global options
	p := prompt.NewPrompt()

	// each prompt type allows for question specific overrides
	answer, err := p.Ask("what is the meaning of life", &prompt.InputOptions{Default: "42"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(answer)
}

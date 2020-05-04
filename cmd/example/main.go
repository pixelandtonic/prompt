package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/pixelandtonic/prompt"
)

func validateAsk(text string) error {
	if text != "42" {
		return errors.New("the answer must be 42")
	}

	return nil
}

func main() {
	// provide global options
	p := prompt.NewPrompt()

	// each prompt type allows for question specific overrides
	answer, err := p.Ask("what is the meaning of life", &prompt.InputOptions{Default: "42", Validator: validateAsk})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("answered:", answer)

	// you can also use Confirm to always return a boolean which defaults to false
	confirm, err := p.Confirm("Do you confirm these changes", &prompt.InputOptions{Default: "yes"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("confirmed:", confirm)

	// provide a select option
	selected, index, err := p.Select("Select an option", []string{"Ludicrous mode", "Normal mode"}, &prompt.InputOptions{Default: "1"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("selected option:", selected)
	fmt.Println("selected index:", index)
}

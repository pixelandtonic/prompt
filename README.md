# Prompt

A simple Go package to make terminal input styled a little more like console input in Craft CMS commands.

## Installation

## Example Usage

To create a new default prompt:

```go
p := prompt.NewPrompt()
``` 

### Passing Options
You can also override the options for a prompt using:

```go
p := prompt.NewWithOptions(opts)
```

### Asking for User Input
To prompt a user to enter input, you simply `.Ask`:

```go
answer, err := p.Ask("What is the answer to the Ultimate Question of Life, the Universe, and Everything", &prompt.InputOptions{Default: "42"})
if err != nil {
	log.Println(err)
}
fmt.Println("the user answered:", answer)
```

> Note: Notice how we did not add a space or ? at the end of the question. The default options when calling `.NewPrompt()` will always add a space and question mark at the end of the output question. If you want to override that pass options using [`prompt.Options`](#passing-options)

If the user does not provide an input, the default value passed in the options will be returned.

### Validating Input
You can also validate input with a simple func:

```go
answer, err := p.Ask("What is the answer to the Ultimate Question of Life, the Universe, and Everything", &prompt.InputOptions{Default: "42", Validator: validateTheMeaning})
if err != nil {
	log.Println(err)
}
fmt.Println("answered:", answer)
```

The validate function is simple and takes a string and returns an error:

```go
func validateTheMeaning(input string) error {
    if input != "42" {
        return errors.New("wrong, the answer is 42")
    }
    return nil
}
```

### Confirming Actions

You can also use Confirm to always return a boolean which defaults to false

```go
confirm, err := p.Confirm("Do you confirm these changes", &prompt.InputOptions{Default: "yes"})
if err != nil {
	log.Println(err)
}
fmt.Println("confirmed:", confirm)
```

You can also pass a validator to confirm, but this will default to checking if the input contains a `y`.

### Providing a Selection

You can also create a selection of items to choose from:

```go
selected, index, err := p.Select("Select an speed", []string{"Ludicrous mode", "Normal mode"}, &prompt.InputOptions{Default: "1"})
if err != nil {
	log.Println(err)
}
```

package prompt

import (
	"errors"
	"github.com/manifoldco/promptui"
	"strings"
)

// String show a password and parse to string.
func StringPwd(name string) (string, error) {
	var prompt = promptui.Prompt{
		Label:     name,
		Validate:  validateEmptyInput,
		Templates: defaultTemplate(),
		Mask:      '*',
	}
	return prompt.Run()
}

// String show a prompt and parse to string.
func String(name string, required bool) (string, error) {
	var prompt promptui.Prompt

	if required {
		prompt = promptui.Prompt{
			Label:     name,
			Validate:  validateEmptyInput,
			Templates: defaultTemplate(),
		}
	} else {
		prompt = promptui.Prompt{
			Label:     name,
			Templates: defaultTemplate(),
		}
	}

	return prompt.Run()
}


// List show a prompt with options and parse to string.
func List(name string, items []string) (string, error) {
	prompt := promptui.Select{
		Items:     items,
		Templates: defaultSelectTemplate(name),
	}
	_, result, err := prompt.Run()
	return result, err
}

func defaultTemplate() *promptui.PromptTemplates {
	return &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | bold }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
}

func defaultSelectTemplate(label string) *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label: label,
	}
}

func validateEmptyInput(input string) error {
	if len(strings.TrimSpace(input)) < 1 {
		return errors.New("this input must not be empty")
	}
	return nil
}
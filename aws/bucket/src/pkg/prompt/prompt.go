package prompt

import (
	"github.com/manifoldco/promptui"
)

// list show a prompt with options and parse to string.
func List(name string, items []string) (string, error) {
	prompt := promptui.Select{
		Items:     items,
		Templates: defaultSelectTemplate(name),
	}
	_, result, err := prompt.Run()
	return result, err
}

func defaultSelectTemplate(label string) *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label: label,
	}
}


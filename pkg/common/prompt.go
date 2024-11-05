package common

import (
	"strconv"

	"github.com/manifoldco/promptui"
)

func PromptString(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: ValidateEmptyInput,
	}

	return prompt.Run()
}

func PromptInteger(name string) (int64, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: ValidateIntegerNumberInput,
	}

	propmtResult, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	parseInt, _ := strconv.ParseInt(propmtResult, 0, 64)
	return parseInt, err
}

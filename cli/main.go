package main

import (
	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Receive payment?",
		Items: []string{"Yes"},
	}

	_, _, error := prompt.Run()
	if error != nil {
		panic(error)
	}
}

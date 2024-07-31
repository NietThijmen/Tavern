package input

import (
	"fmt"
	"strconv"
)

func Select(options []string) string {
	// print the options with a number in front of it
	// read the input
	// return the selected option

	for id, option := range options {
		println("[" + strconv.Itoa(id) + "] " + option)
	}

	var selected string
	_, err := fmt.Scanln(&selected)
	if err != nil {
		return ""
	}

	selectedAsInt, err := strconv.Atoi(selected)
	if err != nil {
		return ""
	}

	if selectedAsInt < 0 || selectedAsInt >= len(options) {
		return ""
	}

	return options[selectedAsInt]
}

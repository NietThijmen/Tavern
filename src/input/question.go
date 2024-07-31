package input

import "fmt"

func Question(question string) string {
	println(question)
	var answer string
	_, err := fmt.Scanln(&answer)
	if err != nil {
		return ""
	}

	return answer
}

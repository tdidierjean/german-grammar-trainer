package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tdidierjean/german_grammar/german_grammar_cli"
)

const exerciseTypeObjects = 1
const exerciseTypePrepositions = 2

func main() {
	fmt.Println("Choose an exercise type:")
	fmt.Println("1. Cases for direct or indirect objects")
	fmt.Println("2. Cases following prepositions")
	choice, err := strconv.Atoi(strings.TrimSpace(readInput()))
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	for {
		var exercise *german_grammar_cli.Exercise
		switch choice {
		case exerciseTypeObjects:
			exercise = german_grammar_cli.GetExercise()
			break
		case exerciseTypePrepositions:
			exercise = german_grammar_cli.GetPrepositionExercise()
			break
		default:
			fmt.Println("Invalid choice")
			return
		}
		fmt.Printf("%s (%s)", exercise.Sentence, exercise.Hint)
		fmt.Println()
		response := strings.TrimSpace(readInput())

		switch response {
		case "":
			return
		case exercise.Answer:
			fmt.Println("Richtig!")
			break
		default:
			fmt.Printf("Falsch! (%s)\n", exercise.Answer)
		}
	}
}

func readInput() string {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()
}

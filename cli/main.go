package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tdidierjean/german_grammar/exercises"
)

const exerciseTypeObjectsParam = 1
const exerciseTypePrepositionsParam = 2
const exerciseTypeAdjectivesParam = 3

var exerciseGenerator exercises.ExerciseGenerator

func main() {
	var randomizer = new(exercises.Randomizer)
	exerciseGenerator := exercises.ExerciseGenerator{Randomizer: randomizer}

	fmt.Println("Choose an exercise type:")
	fmt.Println("1. Cases for direct or indirect objects")
	fmt.Println("2. Cases following prepositions")
	fmt.Println("3. Cases for adjectives")
	choice, err := strconv.Atoi(strings.TrimSpace(readInput()))
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	for {
		var exercise *exercises.Exercise
		switch choice {
		case exerciseTypeObjectsParam:
			exercise = exerciseGenerator.GetObjectExercise()
			break
		case exerciseTypePrepositionsParam:
			exercise = exerciseGenerator.GetPrepositionExercise()
			break
		case exerciseTypeAdjectivesParam:
			exercise = exerciseGenerator.GetAdjectiveExercise()
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

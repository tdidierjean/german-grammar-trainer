package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	grammarexercise "github.com/tdidierjean/german_grammar/german_grammar_server/app"
)

const exerciseTypeObjectsParam = 1
const exerciseTypePrepositionsParam = 2
const exerciseTypeAdjectivesParam = 3

var exerciseGenerator grammarexercise.ExerciseGenerator

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/tdidierjean/german_grammar/german_grammar_server/.env")); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	exerciseGenerator := grammarexercise.CreateExerciseGenerator()

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
		exerciseType := ""
		switch choice {
		case exerciseTypeObjectsParam:
			exerciseType = grammarexercise.ExerciseTypeObject
			break
		case exerciseTypePrepositionsParam:
			exerciseType = grammarexercise.ExerciseTypePreposition
			break
		case exerciseTypeAdjectivesParam:
			exerciseType = grammarexercise.ExerciseTypeAdjective
			break
		default:
			fmt.Println("Invalid choice")
			return
		}

		exercise, err := exerciseGenerator.GetExercise([]string{exerciseType})

		if err != nil {
			fmt.Println("Unexpected error!")
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

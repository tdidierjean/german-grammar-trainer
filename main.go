package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		exercise := getExercise()
		fmt.Printf("%s (%s)", exercise.sentence, exercise.hint)
		fmt.Println()
		response := readInput()

		switch response {
		case "":
			return
		case exercise.answer:
			fmt.Println("Richtig!")
			break
		default:
			fmt.Printf("Falsch! (%s)\n", exercise.answer)
		}
	}
}

func readInput() string {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()
}

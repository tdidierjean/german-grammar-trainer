package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	exercise := getExercise()
	fmt.Printf("%s (%s)", exercise.sentence, exercise.hint)
	fmt.Println()
	response := readInput()

	if response == exercise.answer {
		fmt.Println("Richtig!")
	} else {
		fmt.Printf("Falsch! (%s)\n", exercise.answer)
	}
}

func readInput() string {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()
}

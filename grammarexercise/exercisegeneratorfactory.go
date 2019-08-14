package grammarexercise

// CreateExerciseGenerator factory method to get an exercise generator
func CreateExerciseGenerator() *ExerciseGenerator {
	var randomizer = new(Randomizer)
	return &ExerciseGenerator{Randomizer: randomizer}
}

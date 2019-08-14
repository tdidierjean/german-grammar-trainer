package exercise

import "testing"

func TestGetObjectExercise(t *testing.T) {
	var randomizer = new(Randomizer)
	exerciseGenerator := ExerciseGenerator{Randomizer: randomizer}
	exercise := exerciseGenerator.GetObjectExercise()

	if exercise == nil {
		t.Errorf("expected an exercise")
	}

	// got := Hello()
	// want := "Hello, world"

	// if got != want {
	// t.Errorf("got %q want %q", got, want)
	// }
}

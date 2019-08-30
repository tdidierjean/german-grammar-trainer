package grammarexercise

import "testing"

func assertEquals(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetObjectExercise(t *testing.T) {
	var randomizer = new(Randomizer)
	exerciseGenerator := ExerciseGenerator{Randomizer: randomizer}

	noun := Noun{"Mann", Masculine}
	template := ExerciseTemplate{"Ich habe ... gegessen", []Noun{noun}, Accusative}

	exercise := exerciseGenerator.GetObjectExercise([]ExerciseTemplate{template})

	if exercise == nil {
		t.Errorf("expected an exercise")
	}

	assertEquals(t, exercise.Sentence, template.sentence)
}

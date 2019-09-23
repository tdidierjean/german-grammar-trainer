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

	exercise := exerciseGenerator.getObjectExercise([]ExerciseTemplate{template}, DefiniteArticlesCases)

	if exercise == nil {
		t.Errorf("expected an exercise")
	}

	assertEquals(t, exercise.Sentence, template.sentence)
	assertEquals(t, exercise.Answer, "den")
}

func TestGetPrepositionExercise(t *testing.T) {
	var randomizer = new(Randomizer)
	exerciseGenerator := ExerciseGenerator{Randomizer: randomizer}

	noun := Noun{"Park", Masculine}
	template := ExerciseTemplate{"Ich bin %s ... gefahren", []Noun{noun}, ""}

	preposition := Preposition{"mit", Dative}

	exercise := exerciseGenerator.getPrepositionExercise([]ExerciseTemplate{template}, DefiniteArticlesCases, []Preposition{preposition})

	if exercise == nil {
		t.Errorf("expected an exercise")
	}

	assertEquals(t, exercise.Sentence, "Ich bin mit ... gefahren")
	assertEquals(t, exercise.Answer, "dem")
}

func TestGetAdjectiveExercise(t *testing.T) {
	var randomizer = new(Randomizer)
	exerciseGenerator := ExerciseGenerator{Randomizer: randomizer}

	noun := Noun{"Birne", Feminine}
	template := ExerciseTemplate{"Ich esse %s ... %s.", []Noun{noun}, Accusative}
	adjective := Adjective{"klein"}

	exercise := exerciseGenerator.getAdjectiveExercise([]ExerciseTemplate{template}, DefiniteArticlesCases, DefiniteArticlesAdjectiveCaseEndings, []Adjective{adjective})

	if exercise == nil {
		t.Errorf("expected an exercise")
	}

	assertEquals(t, exercise.Sentence, "Ich esse die ... Birne.")
	assertEquals(t, exercise.Answer, "kleine")
}

func BenchmarkGetGetExercises(b *testing.B) {
	var randomizer = new(Randomizer)
	exerciseGenerator := ExerciseGenerator{Randomizer: randomizer}
	exerciseGenerator.GetExercises([]string{"object"}, 10000, false)
}

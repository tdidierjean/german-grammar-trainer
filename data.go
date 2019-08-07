package german_grammar_cli

import (
	"fmt"
	"math/rand"
	"time"
)

type randomPickable interface{}

type Exercise struct {
	Sentence string
	Hint     string
	Answer   string
}

type ExerciseTemplate struct {
	sentence     string
	nouns        []Noun
	grammar_case string
}

var nouns = []Noun{
	{"Birne", Feminine},
	{"Pfirsich", Masculine},
	{"Kraut", Neuter},
	{"Mann", Masculine},
	{"Mädchen", Neuter},
	{"Frau", Feminine},
	{"Park", Masculine},
	{"Gebaüde", Neuter},
	{"See", Feminine},
}

var exerciseTemplates = []ExerciseTemplate{
	{"Ich habe ... gegessen", nouns[0:2], Accusative},
	{"Ich gebe ... ein Buch", nouns[3:6], Dative},
}

func GetExercise() *Exercise {
	exercise := new(Exercise)

	exerciseTemplate := exerciseTemplates[getRandIndex(len(exerciseTemplates))]
	exercise.Sentence = exerciseTemplate.sentence

	articles := articles[rand.Intn(len(articles))]
	noun := exerciseTemplate.nouns[rand.Intn(len(exerciseTemplate.nouns))]

	exercise.Hint = articles.nominative[noun.gender] + " " + noun.word
	switch exerciseTemplate.grammar_case {
	case Accusative:
		exercise.Answer = articles.accusative[noun.gender]
		break
	case Dative:
		exercise.Answer = articles.dative[noun.gender]
		break
	}

	return exercise
}

func GetPrepositionExercise() *Exercise {

	preposition := prepostitions[getRandIndex(len(prepostitions))]

	var prepositionTemplates = []ExerciseTemplate{
		{"Ich habe %s ... gefahren", nouns[7:9], preposition.grammar_case},
	}

	exercise := new(Exercise)

	exerciseTemplate := prepositionTemplates[getRandIndex(len(prepositionTemplates))]
	exercise.Sentence = fmt.Sprintf(exerciseTemplate.sentence, preposition.preposition)

	articles := articles[rand.Intn(len(articles))]
	noun := exerciseTemplate.nouns[rand.Intn(len(exerciseTemplate.nouns))]

	exercise.Hint = articles.nominative[noun.gender] + " " + noun.word
	switch exerciseTemplate.grammar_case {
	case Accusative:
		exercise.Answer = articles.accusative[noun.gender]
		break
	case Dative:
		exercise.Answer = articles.dative[noun.gender]
		break
	}

	return exercise
}

// Get a random number between 0 and length-1
func getRandIndex(length int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(length)
}

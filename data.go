package main

import (
	"math/rand"
	"time"
)

type Exercise struct {
	sentence string
	hint     string
	answer   string
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
}

var exerciseTemplates = []ExerciseTemplate{
	{"Ich habe ... gegessen", nouns[0:2], Accusative},
	{"Ich gebe ... ein Buch", nouns[3:6], Dative},
}

func getExercise() *Exercise {
	exercise := new(Exercise)

	rand.Seed(time.Now().Unix())
	exerciseTemplate := exerciseTemplates[rand.Intn(len(exerciseTemplates))]
	exercise.sentence = exerciseTemplate.sentence

	articles := articles[rand.Intn(len(articles))]
	noun := exerciseTemplate.nouns[rand.Intn(len(exerciseTemplate.nouns))]

	exercise.hint = articles.nominative[noun.gender] + " " + noun.word
	switch exerciseTemplate.grammar_case {
	case Accusative:
		exercise.answer = articles.accusative[noun.gender]
		break
	case Dative:
		exercise.answer = articles.dative[noun.gender]
		break
	}

	return exercise
}

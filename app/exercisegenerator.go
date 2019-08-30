package grammarexercise

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

const exerciseTypeObject = "object"
const exerciseTypePreposition = "preposition"
const exerciseTypeAdjective = "adjective"

var ExerciseTypes = []string{
	exerciseTypeObject,
	exerciseTypePreposition,
	exerciseTypeAdjective,
}

type randomPickable interface{}

type Exercise struct {
	Sentence string
	Hint     string
	Answer   string
}

type ExerciseTemplate struct {
	sentence    string
	nouns       []Noun
	grammarCase string
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

var adjectives = []Adjective{
	{"klein"},
	{"jung"},
	{"blau"},
	{"neu"},
}

var ObjectExerciseTemplates = []ExerciseTemplate{
	{"Ich habe ... gegessen", nouns[0:2], Accusative},
	{"Ich gebe ... ein Buch", nouns[3:6], Dative},
}

var PrepositionTemplates = []ExerciseTemplate{
	{"Ich bin %s ... gefahren", nouns[7:9], ""},
}

var AdjectiveTemplates = []ExerciseTemplate{
	{"%s ... %s ist hier", nouns[3:6], Nominative},
	{"Ich esse %s ... %s.", nouns[0:3], Accusative},
}

type ExerciseGenerator struct {
	Randomizer *Randomizer
}

// GetExercises Get a list of generated exercises according to the types and size requested
// possible types are object, preposition and adjective
// if multiple types are specified, the type of each exercise is randomly picked within those types
func (e *ExerciseGenerator) GetExercises(exerciseTypes []string, count int) ([]*Exercise, error) {
	var exercises []*Exercise
	for i := 0; i < count; i++ {
		switch exerciseTypes[e.Randomizer.getRandIndex(len(exerciseTypes))] {
		case exerciseTypeObject:
			exercises = append(exercises, e.GetObjectExercise(ObjectExerciseTemplates))
			break
		case exerciseTypePreposition:
			exercises = append(exercises, e.GetPrepositionExercise(PrepositionTemplates))
			break
		case exerciseTypeAdjective:
			exercises = append(exercises, e.GetAdjectiveExercise(AdjectiveTemplates))
			break
		default:
			return nil, errors.New("Invalid exercise type requested")
		}
	}

	return exercises, nil
}

// GetObjectExercise Get a single exercise of type "object"
func (e *ExerciseGenerator) GetObjectExercise(templates []ExerciseTemplate) *Exercise {
	exercise := new(Exercise)

	exerciseTemplate := templates[e.Randomizer.getRandIndex(len(templates))]
	exercise.Sentence = exerciseTemplate.sentence

	articles := articles[e.Randomizer.getRandIndex(len(articles))]
	noun := exerciseTemplate.nouns[e.Randomizer.getRandIndex(len(exerciseTemplate.nouns))]

	exercise.Hint = articles.nominative[noun.gender] + " " + noun.word

	exercise.Answer = reflect.ValueOf(articles).FieldByName(exerciseTemplate.grammarCase).Index(int(noun.gender)).String()

	return exercise
}

// GetPrepositionExercise Get a single exercise of type "Preposition"
func (e *ExerciseGenerator) GetPrepositionExercise(templates []ExerciseTemplate) *Exercise {
	exercise := new(Exercise)
	preposition := prepostitions[e.Randomizer.getRandIndex(len(prepostitions))]

	exerciseTemplate := templates[e.Randomizer.getRandIndex(len(templates))]
	exercise.Sentence = fmt.Sprintf(exerciseTemplate.sentence, preposition.preposition)

	articles := articles[e.Randomizer.getRandIndex(len(articles))]
	noun := exerciseTemplate.nouns[e.Randomizer.getRandIndex(len(exerciseTemplate.nouns))]

	exercise.Hint = articles.nominative[noun.gender] + " " + noun.word

	exercise.Answer = reflect.ValueOf(articles).FieldByName(preposition.grammarCase).Index(int(noun.gender)).String()

	return exercise
}

// GetAdjectiveExercise Get a single exercise of type "Adjective"
func (e *ExerciseGenerator) GetAdjectiveExercise(templates []ExerciseTemplate) *Exercise {
	adjective := adjectives[e.Randomizer.getRandIndex(len(adjectives))]

	var articles Cases
	var adjectiveEndings Cases
	// randomly pick either definite articles ("der") or indefinite articles ("ein") for the exercise
	switch e.Randomizer.getRandIndex(2) {
	case 0:
		articles = DefiniteArticlesCases
		adjectiveEndings = DefiniteArticlesAdjectiveCaseEndings
		break
	case 1:
		articles = IndefiniteArticlesCases
		adjectiveEndings = IndefiniteArticlesAdjectiveCaseEndings
		break
	}

	exercise := new(Exercise)

	exerciseTemplate := templates[e.Randomizer.getRandIndex(len(templates))]
	noun := exerciseTemplate.nouns[e.Randomizer.getRandIndex(len(exerciseTemplate.nouns))]
	caseArticles := reflect.ValueOf(articles).FieldByName(exerciseTemplate.grammarCase)
	exercise.Sentence = fmt.Sprintf(exerciseTemplate.sentence, caseArticles.Index(int(noun.gender)), noun.word)

	exercise.Hint = adjective.word
	ending := reflect.ValueOf(adjectiveEndings).FieldByName(exerciseTemplate.grammarCase).Index(int(noun.gender))
	exercise.Answer = adjective.word + ending.String()

	return exercise
}

type Randomizer struct {
	initialized bool
}

// Get a random number between 0 and length-1
func (r *Randomizer) getRandIndex(length int) int {
	// Set the seed once and only once
	if r.initialized != true {
		rand.Seed(time.Now().Unix())
		r.initialized = true
	}

	return rand.Intn(length)
}

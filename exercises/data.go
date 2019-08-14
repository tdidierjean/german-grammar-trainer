package exercises

import (
	"errors"
	"fmt"
	"math/rand"
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
	{"groß"},
	{"blau"},
	{"neu"},
}

var exerciseTemplates = []ExerciseTemplate{
	{"Ich habe ... gegessen", nouns[0:2], Accusative},
	{"Ich gebe ... ein Buch", nouns[3:6], Dative},
}

type ExerciseGenerator struct {
	Randomizer *Randomizer
}

// GetExercises Get a list of generated exercises according to the types and size requested
func (e *ExerciseGenerator) GetExercises(exerciseTypes []string, count int) ([]*Exercise, error) {
	var exercises []*Exercise
	for i := 0; i < count; i++ {
		switch exerciseTypes[e.Randomizer.getRandIndex(len(exerciseTypes))] {
		case exerciseTypeObject:
			exercises = append(exercises, e.GetObjectExercise())
			break
		case exerciseTypePreposition:
			exercises = append(exercises, e.GetPrepositionExercise())
			break
		case exerciseTypeAdjective:
			exercises = append(exercises, e.GetAdjectiveExercise())
			break
		default:
			return nil, errors.New("Invalid exercise type requested")
		}
	}

	return exercises, nil
}

// GetObjectExercise Get a single exercise of type "object"
func (e *ExerciseGenerator) GetObjectExercise() *Exercise {
	exercise := new(Exercise)

	exerciseTemplate := exerciseTemplates[e.Randomizer.getRandIndex(len(exerciseTemplates))]
	exercise.Sentence = exerciseTemplate.sentence

	articles := articles[e.Randomizer.getRandIndex(len(articles))]
	noun := exerciseTemplate.nouns[e.Randomizer.getRandIndex(len(exerciseTemplate.nouns))]

	exercise.Hint = articles.nominative[noun.gender] + " " + noun.word
	switch exerciseTemplate.grammarCase {
	case Accusative:
		exercise.Answer = articles.accusative[noun.gender]
		break
	case Dative:
		exercise.Answer = articles.dative[noun.gender]
		break
	}

	return exercise
}

// GetPrepositionExercise Get a single exercise of type "Preposition"
func (e *ExerciseGenerator) GetPrepositionExercise() *Exercise {

	preposition := prepostitions[e.Randomizer.getRandIndex(len(prepostitions))]

	var prepositionTemplates = []ExerciseTemplate{
		{"Ich bin %s ... gefahren", nouns[7:9], preposition.grammarCase},
	}

	exercise := new(Exercise)

	exerciseTemplate := prepositionTemplates[e.Randomizer.getRandIndex(len(prepositionTemplates))]
	exercise.Sentence = fmt.Sprintf(exerciseTemplate.sentence, preposition.preposition)

	articles := articles[e.Randomizer.getRandIndex(len(articles))]
	noun := exerciseTemplate.nouns[e.Randomizer.getRandIndex(len(exerciseTemplate.nouns))]

	exercise.Hint = articles.nominative[noun.gender] + " " + noun.word
	switch exerciseTemplate.grammarCase {
	case Accusative:
		exercise.Answer = articles.accusative[noun.gender]
		break
	case Dative:
		exercise.Answer = articles.dative[noun.gender]
		break
	}

	return exercise
}

// GetPrepositionExercise Get a single exercise of type "Adjective"
func (e *ExerciseGenerator) GetAdjectiveExercise() *Exercise {

	adjective := adjectives[e.Randomizer.getRandIndex(len(adjectives))]

	var adjectiveTemplates = []ExerciseTemplate{
		{"%s ... %s heißt Tobias", nouns[3:4], Nominative},
	}

	var articles Cases
	var adjectiveEndings Cases
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

	exerciseTemplate := adjectiveTemplates[e.Randomizer.getRandIndex(len(adjectiveTemplates))]
	noun := exerciseTemplate.nouns[e.Randomizer.getRandIndex(len(exerciseTemplate.nouns))]
	exercise.Sentence = fmt.Sprintf(exerciseTemplate.sentence, articles.nominative[0], noun.word)

	exercise.Hint = adjective.word
	switch exerciseTemplate.grammarCase {
	case Nominative:
		exercise.Answer = adjective.word + adjectiveEndings.nominative[0]
		break
	}

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

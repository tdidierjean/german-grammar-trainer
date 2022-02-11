package grammarexercise

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

const ExerciseTypeObject = "object"
const ExerciseTypePreposition = "preposition"
const ExerciseTypeAdjective = "adjective"

var ExerciseTypes = []string{
	ExerciseTypeObject,
	ExerciseTypePreposition,
	ExerciseTypeAdjective,
}

type Exercise struct {
	Sentence     string
	Hint         string
	Answer       string
	Keywords     []string
	Translations map[string]string
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

var defaultAdjectives = []Adjective{
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

// CreateExerciseGenerator factory method to get an exercise generator
func CreateExerciseGenerator() *ExerciseGenerator {
	var randomizer = new(Randomizer)
	return &ExerciseGenerator{Randomizer: randomizer}
}

// GetExercises Get a list of generated exercises according to the types and size requested
// possible types are object, preposition and adjective
// if multiple types are specified, the type of each exercise is randomly picked within those types
func (e *ExerciseGenerator) GetExercises(exerciseTypes []string, count int, withTranslations bool) ([]*Exercise, error) {
	var exercise *Exercise
	var exercises []*Exercise
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		switch exerciseTypes[e.Randomizer.getRandIndex(len(exerciseTypes))] {
		case ExerciseTypeObject:
			exercise = e.GetObjectExerciseDefault(ObjectExerciseTemplates)
		case ExerciseTypePreposition:
			exercise = e.GetPrepositionExerciseDefault(PrepositionTemplates)
		case ExerciseTypeAdjective:
			exercise = e.GetAdjectiveExerciseDefault(AdjectiveTemplates)
		default:
			return nil, errors.New("invalid exercise type requested")
		}

		exercises = append(exercises, exercise)

		// we're using a wait group to account for all the pending asynchronous translation requests
		if withTranslations {
			wg.Add(1)
			go addKeywordsTranslations(exercise, &wg)
		}
	}

	// wait until all the translation requests have returned
	wg.Wait()

	return exercises, nil
}

// GetExercise Get a single generated exercise according to the types requested
// possible types are object, preposition and adjective
// if multiple types are specified, the type of the exercise is randomly picked within those types
func (e *ExerciseGenerator) GetExercise(exerciseTypes []string) (*Exercise, error) {
	exercises, err := e.GetExercises(exerciseTypes, 1, false)

	if err != nil {
		return nil, err
	}

	return exercises[0], nil
}

// getObjectExercise Get a single exercise of type "object", specifying a subset of available articles
func (e *ExerciseGenerator) getObjectExercise(templates []ExerciseTemplate, articles Cases) *Exercise {
	exercise := new(Exercise)

	exerciseTemplate := templates[e.Randomizer.getRandIndex(len(templates))]
	exercise.Sentence = exerciseTemplate.sentence

	noun := exerciseTemplate.nouns[e.Randomizer.getRandIndex(len(exerciseTemplate.nouns))]
	exercise.Keywords = append(exercise.Keywords, noun.String())

	exercise.Hint = articles.nominative[noun.gender] + " " + noun.word
	exercise.Answer = reflect.ValueOf(articles).FieldByName(exerciseTemplate.grammarCase).Index(int(noun.gender)).String()

	return exercise
}

// GetObjectExerciseDefault Get a single exercise of type "object" with a random class of articles
func (e *ExerciseGenerator) GetObjectExerciseDefault(templates []ExerciseTemplate) *Exercise {
	return e.getObjectExercise(templates, AllArticles[e.Randomizer.getRandIndex(len(AllArticles))])
}

// getPrepositionExercise Get a single exercise of type "Preposition"
func (e *ExerciseGenerator) getPrepositionExercise(templates []ExerciseTemplate, articles Cases, prepositions []Preposition) *Exercise {
	exercise := new(Exercise)
	preposition := prepositions[e.Randomizer.getRandIndex(len(prepositions))]

	exerciseTemplate := templates[e.Randomizer.getRandIndex(len(templates))]
	exercise.Sentence = fmt.Sprintf(exerciseTemplate.sentence, preposition.preposition)

	noun := exerciseTemplate.nouns[e.Randomizer.getRandIndex(len(exerciseTemplate.nouns))]
	exercise.Keywords = append(exercise.Keywords, noun.String())

	exercise.Hint = articles.nominative[noun.gender] + " " + noun.word

	exercise.Answer = reflect.ValueOf(articles).FieldByName(preposition.grammarCase).Index(int(noun.gender)).String()

	return exercise
}

// GetPrepositionExerciseDefault Get a single exercise of type "preposition" with a random class of articles
func (e *ExerciseGenerator) GetPrepositionExerciseDefault(templates []ExerciseTemplate) *Exercise {
	return e.getPrepositionExercise(templates, AllArticles[e.Randomizer.getRandIndex(len(AllArticles))], Prepositions)
}

// getAdjectiveExercise Get a single exercise of type "Adjective"
func (e *ExerciseGenerator) getAdjectiveExercise(templates []ExerciseTemplate, articles Cases, adjectiveEndings Cases, adjectives []Adjective) *Exercise {
	adjective := adjectives[e.Randomizer.getRandIndex(len(adjectives))]

	exercise := new(Exercise)

	exerciseTemplate := templates[e.Randomizer.getRandIndex(len(templates))]
	noun := exerciseTemplate.nouns[e.Randomizer.getRandIndex(len(exerciseTemplate.nouns))]
	exercise.Keywords = append(exercise.Keywords, noun.String())

	caseArticles := reflect.ValueOf(articles).FieldByName(exerciseTemplate.grammarCase)
	exercise.Sentence = fmt.Sprintf(exerciseTemplate.sentence, caseArticles.Index(int(noun.gender)), noun.word)

	exercise.Hint = adjective.word
	ending := reflect.ValueOf(adjectiveEndings).FieldByName(exerciseTemplate.grammarCase).Index(int(noun.gender))
	exercise.Answer = adjective.word + ending.String()

	return exercise
}

// GetAdjectiveExerciseDefault Get a single exercise of type "adjective" using default adjetives
func (e *ExerciseGenerator) GetAdjectiveExerciseDefault(templates []ExerciseTemplate) *Exercise {
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

	return e.getAdjectiveExercise(templates, articles, adjectiveEndings, defaultAdjectives)
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

// Fetch translations for the exercise's keywords
func addKeywordsTranslations(exercise *Exercise, wg *sync.WaitGroup) {
	exercise.Translations = make(map[string]string)

	for _, keyword := range exercise.Keywords {
		exercise.Translations[keyword] = TranslateText(keyword, "de", "en")
		wg.Done()
	}
}

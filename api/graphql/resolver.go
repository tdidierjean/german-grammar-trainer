package graphql

import (
	"context"
	"errors"

	grammarexercise "github.com/tdidierjean/german_grammar/german_grammar_server/app"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	connection *grammarexercise.Connection
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

// Exercises return a list of newly generated exercises
// query parameters are always passed as pointers by gqlgen
func (r *queryResolver) Exercises(ctx context.Context, count *int, exerciseType *string) ([]*Exercise, error) {
	exerciseGenerator := grammarexercise.CreateExerciseGenerator()

	if exerciseType == nil {
		return nil, errors.New("No exercise type specified")
	}

	rawExercises, err := exerciseGenerator.GetExercises([]string{*exerciseType}, *count, true)

	if err != nil {
		return nil, err
	}

	var exercises []*Exercise
	for _, exercise := range rawExercises {
		exercises = append(exercises, r.transformExerciseToGraphQL(exercise))
	}

	return exercises, nil
}

// ExerciseTypes return a list of valid exercise types
func (r *queryResolver) ExerciseTypes(ctx context.Context) ([]string, error) {
	return grammarexercise.ExerciseTypes, nil
}

// Transform the business exercise entity into the graphql entity
func (r *Resolver) transformExerciseToGraphQL(exercise *grammarexercise.Exercise) *Exercise {
	var translation Translation
	var translations []*Translation
	for k, v := range exercise.Translations {
		translation = Translation{k, v}
		translations = append(translations, &translation)
	}

	return &Exercise{
		Question:     exercise.Sentence,
		Hint:         exercise.Hint,
		Answer:       exercise.Answer,
		Translations: translations,
	}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

// UpdateExerciseType update the default exercise type for a user
// Note: user ID is currently hardcoded
func (r *mutationResolver) UpdateExerciseType(ctx context.Context, input NewExerciseType) (*string, error) {

	err := r.connection.UpdateUserExerciseType(1, input.ExerciseType)
	return &input.ExerciseType, err
}

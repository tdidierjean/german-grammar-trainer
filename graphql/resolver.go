package graphql

import (
	"context"
	"errors"

	"github.com/tdidierjean/german_grammar/german_grammar_cli"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	exercises []*Exercise
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

// Query parameters are always passed as pointers by gqlgen
func (r *queryResolver) Exercises(ctx context.Context, count *int, exerciseType *string) ([]*Exercise, error) {
	var randomizer = new(german_grammar_cli.Randomizer)
	exerciseGenerator := german_grammar_cli.ExerciseGenerator{Randomizer: randomizer}

	if exerciseType == nil {
		return nil, errors.New("No exercise type specified")
	}

	rawExercises, err := exerciseGenerator.GetExercises([]string{*exerciseType}, *count)

	if err != nil {
		return nil, err
	}

	var exercises []*Exercise
	for _, exercise := range rawExercises {
		exercises = append(exercises, r.transformExeciseToGraphQL(exercise))
	}

	return exercises, nil
}

func (r *Resolver) transformExeciseToGraphQL(exercise *german_grammar_cli.Exercise) *Exercise {
	return &Exercise{
		Question: exercise.Sentence,
		Hint:     exercise.Hint,
		Answer:   exercise.Answer,
	}
}
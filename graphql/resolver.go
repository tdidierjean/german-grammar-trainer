package graphql

import (
	"context"

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
func (r *queryResolver) Exercises(ctx context.Context, count *int) ([]*Exercise, error) {
	var randomizer = new(german_grammar_cli.Randomizer)
	exerciseGenerator := german_grammar_cli.ExerciseGenerator{Randomizer: randomizer}

	var exercises []*Exercise
	for i := 0; i < *count; i++ {
		exercises = append(exercises, r.transformExeciseToGraphQL(exerciseGenerator.GetExercise()))
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

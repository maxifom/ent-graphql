package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"ent-graphql/graph/generated"
	"ent-graphql/graph/model"
)

func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (*model.Note, error) {
	createdNote, err := r.EntClient.Note.Create().SetBody(input.Body).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Note{
		ID:         createdNote.ID,
		Body:       createdNote.Body,
		CreateTime: createdNote.CreateTime,
		UpdateTime: createdNote.UpdateTime,
	}, nil
}

func (r *queryResolver) Notes(ctx context.Context) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) NoteByID(ctx context.Context) (*model.Note, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

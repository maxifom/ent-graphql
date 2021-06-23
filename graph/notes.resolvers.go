package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"ent-graphql/ent"
	"ent-graphql/ent/note"
	"ent-graphql/graph/model"
)

func (r *mutationResolver) CreateNote(ctx context.Context, input model.CreateNotePayload) (*model.Note, error) {
	createdNote, err := r.entClient.Note.Create().SetBody(input.Body).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Note{
		ID:         strconv.Itoa(createdNote.ID),
		Body:       createdNote.Body,
		CreateTime: createdNote.CreateTime,
		UpdateTime: createdNote.UpdateTime,
	}, nil
}

func (r *mutationResolver) UpdateNote(ctx context.Context, id string, input model.UpdateNotePayload) (*model.Note, error) {
	parsedInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	savedID, err := r.entClient.Note.Update().Where(note.ID(parsedInt)).SetBody(input.Body).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("update error: %s", err)
	}

	if savedID == 0 {
		return nil, fmt.Errorf("note not found by id %s", id)
	}

	selectedNote, err := r.entClient.Note.Get(ctx, savedID)
	if err != nil {
		return nil, fmt.Errorf("get error: %s", err)
	}

	return &model.Note{
		ID:         strconv.Itoa(selectedNote.ID),
		Body:       selectedNote.Body,
		CreateTime: selectedNote.CreateTime,
		UpdateTime: selectedNote.UpdateTime,
	}, nil
}

func (r *mutationResolver) DeleteNote(ctx context.Context, id string) (*model.DeleteNoteResponse, error) {
	parsedInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	err = r.entClient.Note.DeleteOneID(parsedInt).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &model.DeleteNoteResponse{Success: false}, nil
		}
		return nil, err
	}

	return &model.DeleteNoteResponse{Success: true}, nil
}

func (r *queryResolver) Notes(ctx context.Context) ([]model.Note, error) {
	notes, err := r.entClient.Note.Query().Order(ent.Desc(note.FieldCreateTime)).All(ctx)
	if err != nil {
		return nil, err
	}

	var resultNotes []model.Note
	for _, n := range notes {
		resultNotes = append(resultNotes, model.Note{
			ID:         strconv.Itoa(n.ID),
			Body:       n.Body,
			CreateTime: n.CreateTime,
			UpdateTime: n.UpdateTime,
		})
	}

	return resultNotes, nil
}

func (r *queryResolver) NoteByID(ctx context.Context, id string) (*model.Note, error) {
	parsedInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	selectedNote, err := r.entClient.Note.Get(ctx, parsedInt)
	if err != nil {
		return nil, err
	}

	return &model.Note{
		ID:         strconv.Itoa(selectedNote.ID),
		Body:       selectedNote.Body,
		CreateTime: selectedNote.CreateTime,
		UpdateTime: selectedNote.UpdateTime,
	}, nil
}

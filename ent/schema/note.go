package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Note holds the schema definition for the Note entity.
type Note struct {
	ent.Schema
}

// Fields of the Note.
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.Text("body"),
	}
}

func (Note) Mixin() []ent.Mixin  {
	return []ent.Mixin {
		mixin.Time{},
	}
}

// Edges of the Note.
func (Note) Edges() []ent.Edge {
	return nil
}

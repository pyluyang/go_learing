package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.Int("age"),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}

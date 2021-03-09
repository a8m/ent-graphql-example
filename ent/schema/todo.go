package schema

import (
	"entgo.io/ent/schema/edge"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			Values("in_progress", "completed").
			Default("in_progress"),
		field.Int("priority").
			Default(0),
	}
}
// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Todo.Type).
			Unique().
			From("children"),
	}
}

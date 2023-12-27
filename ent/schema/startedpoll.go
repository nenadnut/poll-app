package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StartedPoll holds the schema definition for the StartedPoll entity.
type StartedPoll struct {
	ent.Schema
}

// Fields of the StartedPoll.
func (StartedPoll) Fields() []ent.Field {
	return []ent.Field{
		field.Int("poll_id"),
		field.Int("user_id"),
		field.Bool("completed").Default(false),
		field.Time("created_at").Default(time.Now).StructTag(`json:"-"`),
		field.Time("updated_at").Default(time.Now).StructTag(`json:"-"`),
	}
}

// Edges of the StartedPoll.
func (StartedPoll) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("poll", Poll.Type).Ref("started_polls").Field("poll_id").Required().Unique(),
		edge.From("user", User.Type).Ref("started_polls").Field("user_id").Required().Unique(),
		edge.To("completed_questions", CompletedQuestion.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

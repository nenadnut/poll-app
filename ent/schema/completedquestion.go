package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CompletedQuestion holds the schema definition for the CompletedQuestion entity.
type CompletedQuestion struct {
	ent.Schema
}

// Fields of the CompletedQuestion.
func (CompletedQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("started_poll_id"),
		field.Int("question_id"),
		field.JSON("answers", []int{}), // dummy but quicker
		field.Time("created_at").Default(time.Now).StructTag(`json:"-"`),
		field.Time("updated_at").Default(time.Now).StructTag(`json:"-"`),
	}
}

// Edges of the CompletedQuestion.
func (CompletedQuestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("started_poll", StartedPoll.Type).Ref("completed_questions").Field("started_poll_id").Required().Unique(),
		edge.From("question", Question.Type).Ref("completed_questions").Field("question_id").Required().Unique(),
	}
}

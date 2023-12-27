package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// QuestionOption holds the schema definition for the QuestionOption entity.
type QuestionOption struct {
	ent.Schema
}

// Fields of the QuestionOption.
func (QuestionOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Time("created_at").Default(time.Now).StructTag(`json:"-"`),
		field.Time("updated_at").Default(time.Now).StructTag(`json:"-"`),
		field.Int("question_id"),
	}
}

// Edges of the QuestionOption.
func (QuestionOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("next_option", QuestionOption.Type).Unique().From("next_option_inv").Unique(),
		edge.From("question", Question.Type).Ref("options").Field("question_id").Required().Unique(),
	}
}

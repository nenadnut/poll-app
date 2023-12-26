package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Bool("head").Default(false).StructTag(`json:"-"`),
		field.Int("num_of_answers").Default(1),
		field.Time("created_at").Default(time.Now).StructTag(`json:"-"`),
		field.Time("updated_at").Default(time.Now).StructTag(`json:"-"`),
		field.Int("poll_id"),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("options", QuestionOption.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("next_question", Question.Type).Unique().
			// Field("next_question_id").
			From("next_question_inv"),
		edge.From("poll", Poll.Type).Ref("questions").Field("poll_id").Required().Unique(),
	}
}

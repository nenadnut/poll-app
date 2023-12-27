package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name"),
		field.String("last_name"),
		field.String("email").Unique(),
		field.String("password").StructTag(`json:"-"`),
		field.Enum("role").Values("regular", "admin").Default("regular"),
		field.Time("created_at").Default(time.Now).StructTag(`json:"-"`),
		field.Time("updated_at").Default(time.Now).StructTag(`json:"-"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("polls", Poll.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("started_polls", StartedPoll.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

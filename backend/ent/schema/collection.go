package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Unique().Immutable(),
		// Enum, 0 for "wish", 1 for "watched", 2 for "doing", 3 for "on_hold", 4 for "dropped"
		field.Uint8("type").Max(4).Min(0),
		field.Bool("has_comment").Default(false),
		field.String("comment").Default("").MaxLen(100),
		field.Uint8("score").Default(0),
		field.String("add_time").Default("2000-01-01"), // add date
		field.Uint8("ep_status").Default(0),            // watching progress
		// maybe private
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", Members.Type).Ref("collections").Unique(),
		edge.From("subject", Subject.Type).Ref("collections").Unique(),
	}
}

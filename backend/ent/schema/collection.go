package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("type"),
		field.String("comment").Default("").MaxLen(100),
		field.Int8("score").Default(10),
		field.String("time").Default(time.Now().Format("2006-01-02")),
		field.Uint8("episode").Default(0),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", Members.Type).Ref("collections").Unique(),
		edge.From("subject", Subject.Type).Ref("collections").Unique(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("uid"),
		field.Int64("sub_id"),
		field.Uint8("type"),
		field.Bool("if_comment").Default(false),
		field.String("comment").MaxLen(100).MinLen(5),
		field.Int8("score").Default(100),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return nil
}

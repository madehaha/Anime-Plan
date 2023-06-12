package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Members holds the schema definition for the Members entity.
type Members struct {
	ent.Schema
}

// Fields of the Members.
func (Members) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").StorageKey("uid").Unique(),
		field.String("username").MaxLen(30).Default(""),
		field.String("nickname").MaxLen(30),
		field.String("avatar").MaxLen(255),
		field.Uint8("gid").Default(0), // group id
		field.Time("register_time"),
	}
}

// Edges of the Members.
func (Members) Edges() []ent.Edge {
	return nil
}

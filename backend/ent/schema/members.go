package schema

import (
	"regexp"

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
		field.Uint32("id").StorageKey("uid").Unique().Immutable(),
		field.String("username").MaxLen(30).Nillable().Unique().Match(
			regexp.
				MustCompile("^[a-zA-Z_]{1,}[a-zA-Z0-9_]*$"),
		),
		field.String("email").MaxLen(50).Unique(),
		field.String("password"),
		field.String("nickname").MaxLen(30),
		field.String("avatar").MaxLen(255).Default("https://lain.bgm.tv/pic/user/l/icon.jpg"),
		field.Uint8("gid").Default(0), // group id
		field.String("register_time").Immutable(),
	}
}

// Edges of the Members.
func (Members) Edges() []ent.Edge {
	return nil
}

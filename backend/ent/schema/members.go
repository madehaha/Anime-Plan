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
		field.String("email").MaxLen(50),
		field.String("password").MaxLen(30).MinLen(5),
		field.String("nickname").MaxLen(30),
		field.String("avatar").MaxLen(255).Default("https://lain.bgm.tv/pic/user/l/icon.jpg"),
		field.Uint8("gid").Default(0), // group id
		field.String("register_time"),
	}
}

// Edges of the Members.
func (Members) Edges() []ent.Edge {
	return nil
}

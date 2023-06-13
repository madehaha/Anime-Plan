package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{

		field.String("image").Default("https://lain.bgm.tv/pic/user/l/icon.jpg"),
		field.String("summary").MaxLen(300),
		field.String("name"),
		field.String("date"), //release date
		field.String("name_cn"),
		field.Uint32("on_hold").Default(0),
		field.Uint32("wish").Default(0),
		field.Uint32("doing").Default(0),
		field.Uint8("subject_type").Default(0),
		field.Uint32("collect").Default(0),
	}
}

// Annotations of the User.
func (Subject) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "subjects"},
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return nil
}

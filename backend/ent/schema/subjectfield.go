package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubjectField holds the schema definition for the SubjectField entity.
type SubjectField struct {
	ent.Schema
}

// Fields of the SubjectField.
func (SubjectField) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("rate_1").Default(0),
		field.Uint32("rate_2").Default(0),
		field.Uint32("rate_3").Default(0),
		field.Uint32("rate_4").Default(0),
		field.Uint32("rate_5").Default(0),
		field.Uint32("rate_6").Default(0),
		field.Uint32("rate_7").Default(0),
		field.Uint32("rate_8").Default(0),
		field.Uint32("rate_9").Default(0),
		field.Uint32("rate_10").Default(0),
		field.Float("average_score").Max(10.0).Default(0),
		field.Uint32("rank").Default(0),
		field.Uint32("year"),                // release year
		field.Uint8("month").Min(1).Max(12), // release month
		field.Uint8("date").Min(1).Max(31),  // release  date
		field.Uint8("weekday").Min(1).Max(7),
	}
}

// Edges of the SubjectField.
func (SubjectField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subject", Subject.Type).Ref("subject_field").Unique(),
	}
}

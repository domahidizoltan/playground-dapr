package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Balance holds the schema definition for the Balance entity.
type Balance struct {
	ent.Schema
}

// Fields of the Balance.
func (Balance) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique(),
		field.Float("balance").Min(0).Default(0),
		field.Float("locked").Min(0).Default(0),
	}
}

// Edges of the Balance.
func (Balance) Edges() []ent.Edge {
	return nil
}

// Annotations of the Balance.
func (Balance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "balance"},
	}
}

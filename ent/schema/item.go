package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("note"),
		field.Int("amount").Positive(),
		field.String("status").NotEmpty().Default("new"),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("store", Store.Type).Ref("items").Unique(),
		edge.From("category", Category.Type).Ref("items").Unique(),
	}
}

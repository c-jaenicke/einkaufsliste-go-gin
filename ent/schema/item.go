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
		field.Int("store_id").Optional(),
		field.Int("category_id").Optional(),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("store", Store.Type).Ref("items").Field("store_id").Unique(),
		edge.From("category", Category.Type).Ref("items").Field("category_id").Unique(),
	}
}

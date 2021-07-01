package schema

import "entgo.io/ent"

// OrderAddress holds the schema definition for the OrderAddress entity.
type OrderAddress struct {
	ent.Schema
}

// Fields of the OrderAddress.
func (OrderAddress) Fields() []ent.Field {
	return nil
}

// Edges of the OrderAddress.
func (OrderAddress) Edges() []ent.Edge {
	return nil
}

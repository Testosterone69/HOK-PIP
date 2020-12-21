package schema

import "github.com/facebookincubator/ent"

// HEE holds the schema definition for the HEE entity.
type HEE struct {
	ent.Schema
}

// Fields of the HEE.
func (HEE) Fields() []ent.Field {
	return nil
}

// Edges of the HEE.
func (HEE) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"github.com/facebook/ent"
)

// Pikachu holds the schema definition for the Pikachu entity.
type Pikachu struct {
	ent.Schema
}

// Fields of the Pikachu.
func (Pikachu) Fields() []ent.Field {
	return nil
}

// Edges of the Pikachu.
func (Pikachu) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
 )

// Heeted holds the schema definition for the Heeted entity.
type Heeted struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Positive(),
		field.String("name").NotEmpty(),
	}
 }

// Edges of the Heeted.
func (Heeted) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Devices holds the schema definition for the Devices entity.
type Devices struct {
	ent.Schema
}

// Fields of the Device.
func (Devices) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip"),
		field.String("device"),
	}
}

// Edges of the Device.
func (Devices) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Users.Type).
			Ref("devices").
			Unique(),
	}
}

func (Devices) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}

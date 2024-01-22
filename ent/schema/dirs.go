package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Dirs holds the schema definition for the Dirs entity.
type Dirs struct {
	ent.Schema
}

// Fields of the Dir.
func (Dirs) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name"),
	}
}

// Edges of the Dir.
func (Dirs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Users.Type).
			Ref("dirs").
			Unique(),
		edge.To("saves", Saves.Type),

		edge.To("subdir", Dirs.Type),

		edge.From("pdir", Dirs.Type).
			Ref("subdir"),
	}
}

func (Dirs) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}

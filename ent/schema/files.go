package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Files holds the schema definition for the Files entity.
type Files struct {
	ent.Schema
}

// Fields of the File.
func (Files) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("hash").
			Unique().
			Comment("ipfs hash"),
		field.
			String("name").
			NotEmpty().
			Comment("file name"),
		field.
			Uint64("size").
			Comment("file size, number of bytes in the stored file"),
	}
}

// Edges of the File.
func (Files) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("saves", Saves.Type),
	}
}

func (Files) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash").
			Unique(),
	}
}

func (Files) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

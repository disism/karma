package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Saves holds the schema definition for the Saves entity.
type Saves struct {
	ent.Schema
}

// Fields of the Saved.
func (Saves) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(255),
		field.
			String("caption").
			MaxLen(255).
			Optional().
			Comment("the descriptive text or title of a document, image, or other media element. It is used to provide a short description of the content, characteristics or context of a document."),
	}
}

// Edges of the Saved.
func (Saves) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("file", Files.Type).
			Ref("saves").
			Unique(),
		edge.
			From("owner", Users.Type).
			Ref("saves").
			Unique(),
		edge.
			From("dir", Dirs.Type).
			Ref("saves"),
	}
}

func (Saves) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}

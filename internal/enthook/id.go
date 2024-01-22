package enthook

import (
	"context"
	"fmt"
	"github.com/disism/karma/ent"
	"github.com/sony/sonyflake"
)

// IDHook initializes a hook for generating unique IDs using a sonyflake generator.
//
// It takes a client as a parameter, which is an instance of ent.Client.
// The function does not return any values.
func IDHook(client *ent.Client) {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})

	type IDSetter interface {
		SetID(uint64)
	}
	client.Use(
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				is, ok := m.(IDSetter)
				if !ok {
					return nil, fmt.Errorf("id hook failed to %T", m)
				}
				id, err := sf.NextID()
				if err != nil {
					return nil, err
				}
				is.SetID(id)
				return next.Mutate(ctx, m)
			})
		},
	)
}

package database

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/disism/karma/ent"
	"github.com/disism/karma/internal/enthook"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/exp/slog"
)

type Database struct {
	Ctx    context.Context
	Client *ent.Client
}

func SQLiteConnectURL(name string) string {
	return fmt.Sprintf("file:%s.db?&cache=shared&_fk=1", name)
}

func New(ctx context.Context) (*Database, error) {
	client, err := ent.Open(
		dialect.SQLite,
		SQLiteConnectURL("karma"),
	)
	if err != nil {
		slog.Error("connect database error: ", err.Error())
		return nil, err
	}
	if err := client.
		Schema.
		Create(ctx); err != nil {
		return nil, err
	}
	enthook.IDHook(client)

	return &Database{
		Ctx:    ctx,
		Client: client,
	}, nil
}

func (c *Database) Close() error {
	return c.Client.Close()
}

func (c *Database) Tx(ctx context.Context) (*ent.Tx, error) {
	return c.Client.Tx(ctx)
}

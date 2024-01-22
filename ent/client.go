// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/disism/karma/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/disism/karma/ent/devices"
	"github.com/disism/karma/ent/dirs"
	"github.com/disism/karma/ent/files"
	"github.com/disism/karma/ent/saves"
	"github.com/disism/karma/ent/users"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Devices is the client for interacting with the Devices builders.
	Devices *DevicesClient
	// Dirs is the client for interacting with the Dirs builders.
	Dirs *DirsClient
	// Files is the client for interacting with the Files builders.
	Files *FilesClient
	// Saves is the client for interacting with the Saves builders.
	Saves *SavesClient
	// Users is the client for interacting with the Users builders.
	Users *UsersClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Devices = NewDevicesClient(c.config)
	c.Dirs = NewDirsClient(c.config)
	c.Files = NewFilesClient(c.config)
	c.Saves = NewSavesClient(c.config)
	c.Users = NewUsersClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Devices: NewDevicesClient(cfg),
		Dirs:    NewDirsClient(cfg),
		Files:   NewFilesClient(cfg),
		Saves:   NewSavesClient(cfg),
		Users:   NewUsersClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Devices: NewDevicesClient(cfg),
		Dirs:    NewDirsClient(cfg),
		Files:   NewFilesClient(cfg),
		Saves:   NewSavesClient(cfg),
		Users:   NewUsersClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Devices.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Devices.Use(hooks...)
	c.Dirs.Use(hooks...)
	c.Files.Use(hooks...)
	c.Saves.Use(hooks...)
	c.Users.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Devices.Intercept(interceptors...)
	c.Dirs.Intercept(interceptors...)
	c.Files.Intercept(interceptors...)
	c.Saves.Intercept(interceptors...)
	c.Users.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *DevicesMutation:
		return c.Devices.mutate(ctx, m)
	case *DirsMutation:
		return c.Dirs.mutate(ctx, m)
	case *FilesMutation:
		return c.Files.mutate(ctx, m)
	case *SavesMutation:
		return c.Saves.mutate(ctx, m)
	case *UsersMutation:
		return c.Users.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// DevicesClient is a client for the Devices schema.
type DevicesClient struct {
	config
}

// NewDevicesClient returns a client for the Devices from the given config.
func NewDevicesClient(c config) *DevicesClient {
	return &DevicesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `devices.Hooks(f(g(h())))`.
func (c *DevicesClient) Use(hooks ...Hook) {
	c.hooks.Devices = append(c.hooks.Devices, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `devices.Intercept(f(g(h())))`.
func (c *DevicesClient) Intercept(interceptors ...Interceptor) {
	c.inters.Devices = append(c.inters.Devices, interceptors...)
}

// Create returns a builder for creating a Devices entity.
func (c *DevicesClient) Create() *DevicesCreate {
	mutation := newDevicesMutation(c.config, OpCreate)
	return &DevicesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Devices entities.
func (c *DevicesClient) CreateBulk(builders ...*DevicesCreate) *DevicesCreateBulk {
	return &DevicesCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DevicesClient) MapCreateBulk(slice any, setFunc func(*DevicesCreate, int)) *DevicesCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DevicesCreateBulk{err: fmt.Errorf("calling to DevicesClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DevicesCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DevicesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Devices.
func (c *DevicesClient) Update() *DevicesUpdate {
	mutation := newDevicesMutation(c.config, OpUpdate)
	return &DevicesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DevicesClient) UpdateOne(d *Devices) *DevicesUpdateOne {
	mutation := newDevicesMutation(c.config, OpUpdateOne, withDevices(d))
	return &DevicesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DevicesClient) UpdateOneID(id uint64) *DevicesUpdateOne {
	mutation := newDevicesMutation(c.config, OpUpdateOne, withDevicesID(id))
	return &DevicesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Devices.
func (c *DevicesClient) Delete() *DevicesDelete {
	mutation := newDevicesMutation(c.config, OpDelete)
	return &DevicesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DevicesClient) DeleteOne(d *Devices) *DevicesDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DevicesClient) DeleteOneID(id uint64) *DevicesDeleteOne {
	builder := c.Delete().Where(devices.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DevicesDeleteOne{builder}
}

// Query returns a query builder for Devices.
func (c *DevicesClient) Query() *DevicesQuery {
	return &DevicesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDevices},
		inters: c.Interceptors(),
	}
}

// Get returns a Devices entity by its id.
func (c *DevicesClient) Get(ctx context.Context, id uint64) (*Devices, error) {
	return c.Query().Where(devices.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DevicesClient) GetX(ctx context.Context, id uint64) *Devices {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Devices.
func (c *DevicesClient) QueryUser(d *Devices) *UsersQuery {
	query := (&UsersClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(devices.Table, devices.FieldID, id),
			sqlgraph.To(users.Table, users.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, devices.UserTable, devices.UserColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DevicesClient) Hooks() []Hook {
	return c.hooks.Devices
}

// Interceptors returns the client interceptors.
func (c *DevicesClient) Interceptors() []Interceptor {
	return c.inters.Devices
}

func (c *DevicesClient) mutate(ctx context.Context, m *DevicesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DevicesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DevicesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DevicesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DevicesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Devices mutation op: %q", m.Op())
	}
}

// DirsClient is a client for the Dirs schema.
type DirsClient struct {
	config
}

// NewDirsClient returns a client for the Dirs from the given config.
func NewDirsClient(c config) *DirsClient {
	return &DirsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dirs.Hooks(f(g(h())))`.
func (c *DirsClient) Use(hooks ...Hook) {
	c.hooks.Dirs = append(c.hooks.Dirs, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `dirs.Intercept(f(g(h())))`.
func (c *DirsClient) Intercept(interceptors ...Interceptor) {
	c.inters.Dirs = append(c.inters.Dirs, interceptors...)
}

// Create returns a builder for creating a Dirs entity.
func (c *DirsClient) Create() *DirsCreate {
	mutation := newDirsMutation(c.config, OpCreate)
	return &DirsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Dirs entities.
func (c *DirsClient) CreateBulk(builders ...*DirsCreate) *DirsCreateBulk {
	return &DirsCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DirsClient) MapCreateBulk(slice any, setFunc func(*DirsCreate, int)) *DirsCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DirsCreateBulk{err: fmt.Errorf("calling to DirsClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DirsCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DirsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Dirs.
func (c *DirsClient) Update() *DirsUpdate {
	mutation := newDirsMutation(c.config, OpUpdate)
	return &DirsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DirsClient) UpdateOne(d *Dirs) *DirsUpdateOne {
	mutation := newDirsMutation(c.config, OpUpdateOne, withDirs(d))
	return &DirsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DirsClient) UpdateOneID(id uint64) *DirsUpdateOne {
	mutation := newDirsMutation(c.config, OpUpdateOne, withDirsID(id))
	return &DirsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dirs.
func (c *DirsClient) Delete() *DirsDelete {
	mutation := newDirsMutation(c.config, OpDelete)
	return &DirsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DirsClient) DeleteOne(d *Dirs) *DirsDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DirsClient) DeleteOneID(id uint64) *DirsDeleteOne {
	builder := c.Delete().Where(dirs.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DirsDeleteOne{builder}
}

// Query returns a query builder for Dirs.
func (c *DirsClient) Query() *DirsQuery {
	return &DirsQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDirs},
		inters: c.Interceptors(),
	}
}

// Get returns a Dirs entity by its id.
func (c *DirsClient) Get(ctx context.Context, id uint64) (*Dirs, error) {
	return c.Query().Where(dirs.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DirsClient) GetX(ctx context.Context, id uint64) *Dirs {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Dirs.
func (c *DirsClient) QueryOwner(d *Dirs) *UsersQuery {
	query := (&UsersClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dirs.Table, dirs.FieldID, id),
			sqlgraph.To(users.Table, users.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dirs.OwnerTable, dirs.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySaves queries the saves edge of a Dirs.
func (c *DirsClient) QuerySaves(d *Dirs) *SavesQuery {
	query := (&SavesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dirs.Table, dirs.FieldID, id),
			sqlgraph.To(saves.Table, saves.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, dirs.SavesTable, dirs.SavesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubdir queries the subdir edge of a Dirs.
func (c *DirsClient) QuerySubdir(d *Dirs) *DirsQuery {
	query := (&DirsClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dirs.Table, dirs.FieldID, id),
			sqlgraph.To(dirs.Table, dirs.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, dirs.SubdirTable, dirs.SubdirPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPdir queries the pdir edge of a Dirs.
func (c *DirsClient) QueryPdir(d *Dirs) *DirsQuery {
	query := (&DirsClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dirs.Table, dirs.FieldID, id),
			sqlgraph.To(dirs.Table, dirs.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, dirs.PdirTable, dirs.PdirPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DirsClient) Hooks() []Hook {
	return c.hooks.Dirs
}

// Interceptors returns the client interceptors.
func (c *DirsClient) Interceptors() []Interceptor {
	return c.inters.Dirs
}

func (c *DirsClient) mutate(ctx context.Context, m *DirsMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DirsCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DirsUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DirsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DirsDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Dirs mutation op: %q", m.Op())
	}
}

// FilesClient is a client for the Files schema.
type FilesClient struct {
	config
}

// NewFilesClient returns a client for the Files from the given config.
func NewFilesClient(c config) *FilesClient {
	return &FilesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `files.Hooks(f(g(h())))`.
func (c *FilesClient) Use(hooks ...Hook) {
	c.hooks.Files = append(c.hooks.Files, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `files.Intercept(f(g(h())))`.
func (c *FilesClient) Intercept(interceptors ...Interceptor) {
	c.inters.Files = append(c.inters.Files, interceptors...)
}

// Create returns a builder for creating a Files entity.
func (c *FilesClient) Create() *FilesCreate {
	mutation := newFilesMutation(c.config, OpCreate)
	return &FilesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Files entities.
func (c *FilesClient) CreateBulk(builders ...*FilesCreate) *FilesCreateBulk {
	return &FilesCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *FilesClient) MapCreateBulk(slice any, setFunc func(*FilesCreate, int)) *FilesCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &FilesCreateBulk{err: fmt.Errorf("calling to FilesClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*FilesCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &FilesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Files.
func (c *FilesClient) Update() *FilesUpdate {
	mutation := newFilesMutation(c.config, OpUpdate)
	return &FilesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FilesClient) UpdateOne(f *Files) *FilesUpdateOne {
	mutation := newFilesMutation(c.config, OpUpdateOne, withFiles(f))
	return &FilesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FilesClient) UpdateOneID(id uint64) *FilesUpdateOne {
	mutation := newFilesMutation(c.config, OpUpdateOne, withFilesID(id))
	return &FilesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Files.
func (c *FilesClient) Delete() *FilesDelete {
	mutation := newFilesMutation(c.config, OpDelete)
	return &FilesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FilesClient) DeleteOne(f *Files) *FilesDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FilesClient) DeleteOneID(id uint64) *FilesDeleteOne {
	builder := c.Delete().Where(files.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FilesDeleteOne{builder}
}

// Query returns a query builder for Files.
func (c *FilesClient) Query() *FilesQuery {
	return &FilesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFiles},
		inters: c.Interceptors(),
	}
}

// Get returns a Files entity by its id.
func (c *FilesClient) Get(ctx context.Context, id uint64) (*Files, error) {
	return c.Query().Where(files.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FilesClient) GetX(ctx context.Context, id uint64) *Files {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySaves queries the saves edge of a Files.
func (c *FilesClient) QuerySaves(f *Files) *SavesQuery {
	query := (&SavesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(files.Table, files.FieldID, id),
			sqlgraph.To(saves.Table, saves.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, files.SavesTable, files.SavesColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FilesClient) Hooks() []Hook {
	return c.hooks.Files
}

// Interceptors returns the client interceptors.
func (c *FilesClient) Interceptors() []Interceptor {
	return c.inters.Files
}

func (c *FilesClient) mutate(ctx context.Context, m *FilesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FilesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FilesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FilesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FilesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Files mutation op: %q", m.Op())
	}
}

// SavesClient is a client for the Saves schema.
type SavesClient struct {
	config
}

// NewSavesClient returns a client for the Saves from the given config.
func NewSavesClient(c config) *SavesClient {
	return &SavesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `saves.Hooks(f(g(h())))`.
func (c *SavesClient) Use(hooks ...Hook) {
	c.hooks.Saves = append(c.hooks.Saves, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `saves.Intercept(f(g(h())))`.
func (c *SavesClient) Intercept(interceptors ...Interceptor) {
	c.inters.Saves = append(c.inters.Saves, interceptors...)
}

// Create returns a builder for creating a Saves entity.
func (c *SavesClient) Create() *SavesCreate {
	mutation := newSavesMutation(c.config, OpCreate)
	return &SavesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Saves entities.
func (c *SavesClient) CreateBulk(builders ...*SavesCreate) *SavesCreateBulk {
	return &SavesCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SavesClient) MapCreateBulk(slice any, setFunc func(*SavesCreate, int)) *SavesCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SavesCreateBulk{err: fmt.Errorf("calling to SavesClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SavesCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SavesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Saves.
func (c *SavesClient) Update() *SavesUpdate {
	mutation := newSavesMutation(c.config, OpUpdate)
	return &SavesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SavesClient) UpdateOne(s *Saves) *SavesUpdateOne {
	mutation := newSavesMutation(c.config, OpUpdateOne, withSaves(s))
	return &SavesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SavesClient) UpdateOneID(id uint64) *SavesUpdateOne {
	mutation := newSavesMutation(c.config, OpUpdateOne, withSavesID(id))
	return &SavesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Saves.
func (c *SavesClient) Delete() *SavesDelete {
	mutation := newSavesMutation(c.config, OpDelete)
	return &SavesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SavesClient) DeleteOne(s *Saves) *SavesDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SavesClient) DeleteOneID(id uint64) *SavesDeleteOne {
	builder := c.Delete().Where(saves.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SavesDeleteOne{builder}
}

// Query returns a query builder for Saves.
func (c *SavesClient) Query() *SavesQuery {
	return &SavesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSaves},
		inters: c.Interceptors(),
	}
}

// Get returns a Saves entity by its id.
func (c *SavesClient) Get(ctx context.Context, id uint64) (*Saves, error) {
	return c.Query().Where(saves.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SavesClient) GetX(ctx context.Context, id uint64) *Saves {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFile queries the file edge of a Saves.
func (c *SavesClient) QueryFile(s *Saves) *FilesQuery {
	query := (&FilesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(saves.Table, saves.FieldID, id),
			sqlgraph.To(files.Table, files.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, saves.FileTable, saves.FileColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwner queries the owner edge of a Saves.
func (c *SavesClient) QueryOwner(s *Saves) *UsersQuery {
	query := (&UsersClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(saves.Table, saves.FieldID, id),
			sqlgraph.To(users.Table, users.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, saves.OwnerTable, saves.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDir queries the dir edge of a Saves.
func (c *SavesClient) QueryDir(s *Saves) *DirsQuery {
	query := (&DirsClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(saves.Table, saves.FieldID, id),
			sqlgraph.To(dirs.Table, dirs.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, saves.DirTable, saves.DirPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SavesClient) Hooks() []Hook {
	return c.hooks.Saves
}

// Interceptors returns the client interceptors.
func (c *SavesClient) Interceptors() []Interceptor {
	return c.inters.Saves
}

func (c *SavesClient) mutate(ctx context.Context, m *SavesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SavesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SavesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SavesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SavesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Saves mutation op: %q", m.Op())
	}
}

// UsersClient is a client for the Users schema.
type UsersClient struct {
	config
}

// NewUsersClient returns a client for the Users from the given config.
func NewUsersClient(c config) *UsersClient {
	return &UsersClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `users.Hooks(f(g(h())))`.
func (c *UsersClient) Use(hooks ...Hook) {
	c.hooks.Users = append(c.hooks.Users, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `users.Intercept(f(g(h())))`.
func (c *UsersClient) Intercept(interceptors ...Interceptor) {
	c.inters.Users = append(c.inters.Users, interceptors...)
}

// Create returns a builder for creating a Users entity.
func (c *UsersClient) Create() *UsersCreate {
	mutation := newUsersMutation(c.config, OpCreate)
	return &UsersCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Users entities.
func (c *UsersClient) CreateBulk(builders ...*UsersCreate) *UsersCreateBulk {
	return &UsersCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UsersClient) MapCreateBulk(slice any, setFunc func(*UsersCreate, int)) *UsersCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UsersCreateBulk{err: fmt.Errorf("calling to UsersClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UsersCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UsersCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Users.
func (c *UsersClient) Update() *UsersUpdate {
	mutation := newUsersMutation(c.config, OpUpdate)
	return &UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UsersClient) UpdateOne(u *Users) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsers(u))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UsersClient) UpdateOneID(id uint64) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsersID(id))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Users.
func (c *UsersClient) Delete() *UsersDelete {
	mutation := newUsersMutation(c.config, OpDelete)
	return &UsersDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UsersClient) DeleteOne(u *Users) *UsersDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UsersClient) DeleteOneID(id uint64) *UsersDeleteOne {
	builder := c.Delete().Where(users.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UsersDeleteOne{builder}
}

// Query returns a query builder for Users.
func (c *UsersClient) Query() *UsersQuery {
	return &UsersQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUsers},
		inters: c.Interceptors(),
	}
}

// Get returns a Users entity by its id.
func (c *UsersClient) Get(ctx context.Context, id uint64) (*Users, error) {
	return c.Query().Where(users.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UsersClient) GetX(ctx context.Context, id uint64) *Users {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDevices queries the devices edge of a Users.
func (c *UsersClient) QueryDevices(u *Users) *DevicesQuery {
	query := (&DevicesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(users.Table, users.FieldID, id),
			sqlgraph.To(devices.Table, devices.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, users.DevicesTable, users.DevicesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDirs queries the dirs edge of a Users.
func (c *UsersClient) QueryDirs(u *Users) *DirsQuery {
	query := (&DirsClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(users.Table, users.FieldID, id),
			sqlgraph.To(dirs.Table, dirs.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, users.DirsTable, users.DirsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySaves queries the saves edge of a Users.
func (c *UsersClient) QuerySaves(u *Users) *SavesQuery {
	query := (&SavesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(users.Table, users.FieldID, id),
			sqlgraph.To(saves.Table, saves.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, users.SavesTable, users.SavesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UsersClient) Hooks() []Hook {
	return c.hooks.Users
}

// Interceptors returns the client interceptors.
func (c *UsersClient) Interceptors() []Interceptor {
	return c.inters.Users
}

func (c *UsersClient) mutate(ctx context.Context, m *UsersMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UsersCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UsersDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Users mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Devices, Dirs, Files, Saves, Users []ent.Hook
	}
	inters struct {
		Devices, Dirs, Files, Saves, Users []ent.Interceptor
	}
)
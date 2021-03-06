// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Faroukhamadi/Banketeer/teller_service/tellerdb/ent/migrate"

	"github.com/Faroukhamadi/Banketeer/teller_service/tellerdb/ent/teller"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Teller is the client for interacting with the Teller builders.
	Teller *TellerClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Teller = NewTellerClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Teller: NewTellerClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:    ctx,
		config: cfg,
		Teller: NewTellerClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Teller.
//		Query().
//		Count(ctx)
//
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
	c.Teller.Use(hooks...)
}

// TellerClient is a client for the Teller schema.
type TellerClient struct {
	config
}

// NewTellerClient returns a client for the Teller from the given config.
func NewTellerClient(c config) *TellerClient {
	return &TellerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `teller.Hooks(f(g(h())))`.
func (c *TellerClient) Use(hooks ...Hook) {
	c.hooks.Teller = append(c.hooks.Teller, hooks...)
}

// Create returns a builder for creating a Teller entity.
func (c *TellerClient) Create() *TellerCreate {
	mutation := newTellerMutation(c.config, OpCreate)
	return &TellerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Teller entities.
func (c *TellerClient) CreateBulk(builders ...*TellerCreate) *TellerCreateBulk {
	return &TellerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Teller.
func (c *TellerClient) Update() *TellerUpdate {
	mutation := newTellerMutation(c.config, OpUpdate)
	return &TellerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TellerClient) UpdateOne(t *Teller) *TellerUpdateOne {
	mutation := newTellerMutation(c.config, OpUpdateOne, withTeller(t))
	return &TellerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TellerClient) UpdateOneID(id int) *TellerUpdateOne {
	mutation := newTellerMutation(c.config, OpUpdateOne, withTellerID(id))
	return &TellerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Teller.
func (c *TellerClient) Delete() *TellerDelete {
	mutation := newTellerMutation(c.config, OpDelete)
	return &TellerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TellerClient) DeleteOne(t *Teller) *TellerDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TellerClient) DeleteOneID(id int) *TellerDeleteOne {
	builder := c.Delete().Where(teller.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TellerDeleteOne{builder}
}

// Query returns a query builder for Teller.
func (c *TellerClient) Query() *TellerQuery {
	return &TellerQuery{
		config: c.config,
	}
}

// Get returns a Teller entity by its id.
func (c *TellerClient) Get(ctx context.Context, id int) (*Teller, error) {
	return c.Query().Where(teller.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TellerClient) GetX(ctx context.Context, id int) *Teller {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TellerClient) Hooks() []Hook {
	return c.hooks.Teller
}

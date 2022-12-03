// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"anylogibtc/ent/migrate"

	"anylogibtc/ent/transaction"
	"anylogibtc/ent/wallet"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Transaction is the client for interacting with the Transaction builders.
	Transaction *TransactionClient
	// Wallet is the client for interacting with the Wallet builders.
	Wallet *WalletClient
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
	c.Transaction = NewTransactionClient(c.config)
	c.Wallet = NewWalletClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Transaction: NewTransactionClient(cfg),
		Wallet:      NewWalletClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		Transaction: NewTransactionClient(cfg),
		Wallet:      NewWalletClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Transaction.
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
	c.Transaction.Use(hooks...)
	c.Wallet.Use(hooks...)
}

// TransactionClient is a client for the Transaction schema.
type TransactionClient struct {
	config
}

// NewTransactionClient returns a client for the Transaction from the given config.
func NewTransactionClient(c config) *TransactionClient {
	return &TransactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transaction.Hooks(f(g(h())))`.
func (c *TransactionClient) Use(hooks ...Hook) {
	c.hooks.Transaction = append(c.hooks.Transaction, hooks...)
}

// Create returns a builder for creating a Transaction entity.
func (c *TransactionClient) Create() *TransactionCreate {
	mutation := newTransactionMutation(c.config, OpCreate)
	return &TransactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transaction entities.
func (c *TransactionClient) CreateBulk(builders ...*TransactionCreate) *TransactionCreateBulk {
	return &TransactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transaction.
func (c *TransactionClient) Update() *TransactionUpdate {
	mutation := newTransactionMutation(c.config, OpUpdate)
	return &TransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransactionClient) UpdateOne(t *Transaction) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransaction(t))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransactionClient) UpdateOneID(id int) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransactionID(id))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transaction.
func (c *TransactionClient) Delete() *TransactionDelete {
	mutation := newTransactionMutation(c.config, OpDelete)
	return &TransactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransactionClient) DeleteOne(t *Transaction) *TransactionDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TransactionClient) DeleteOneID(id int) *TransactionDeleteOne {
	builder := c.Delete().Where(transaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransactionDeleteOne{builder}
}

// Query returns a query builder for Transaction.
func (c *TransactionClient) Query() *TransactionQuery {
	return &TransactionQuery{
		config: c.config,
	}
}

// Get returns a Transaction entity by its id.
func (c *TransactionClient) Get(ctx context.Context, id int) (*Transaction, error) {
	return c.Query().Where(transaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransactionClient) GetX(ctx context.Context, id int) *Transaction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySender queries the sender edge of a Transaction.
func (c *TransactionClient) QuerySender(t *Transaction) *WalletQuery {
	query := &WalletQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(wallet.Table, wallet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transaction.SenderTable, transaction.SenderColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRecipient queries the recipient edge of a Transaction.
func (c *TransactionClient) QueryRecipient(t *Transaction) *WalletQuery {
	query := &WalletQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(wallet.Table, wallet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transaction.RecipientTable, transaction.RecipientColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TransactionClient) Hooks() []Hook {
	return c.hooks.Transaction
}

// WalletClient is a client for the Wallet schema.
type WalletClient struct {
	config
}

// NewWalletClient returns a client for the Wallet from the given config.
func NewWalletClient(c config) *WalletClient {
	return &WalletClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `wallet.Hooks(f(g(h())))`.
func (c *WalletClient) Use(hooks ...Hook) {
	c.hooks.Wallet = append(c.hooks.Wallet, hooks...)
}

// Create returns a builder for creating a Wallet entity.
func (c *WalletClient) Create() *WalletCreate {
	mutation := newWalletMutation(c.config, OpCreate)
	return &WalletCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Wallet entities.
func (c *WalletClient) CreateBulk(builders ...*WalletCreate) *WalletCreateBulk {
	return &WalletCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Wallet.
func (c *WalletClient) Update() *WalletUpdate {
	mutation := newWalletMutation(c.config, OpUpdate)
	return &WalletUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WalletClient) UpdateOne(w *Wallet) *WalletUpdateOne {
	mutation := newWalletMutation(c.config, OpUpdateOne, withWallet(w))
	return &WalletUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WalletClient) UpdateOneID(id int) *WalletUpdateOne {
	mutation := newWalletMutation(c.config, OpUpdateOne, withWalletID(id))
	return &WalletUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Wallet.
func (c *WalletClient) Delete() *WalletDelete {
	mutation := newWalletMutation(c.config, OpDelete)
	return &WalletDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WalletClient) DeleteOne(w *Wallet) *WalletDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *WalletClient) DeleteOneID(id int) *WalletDeleteOne {
	builder := c.Delete().Where(wallet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WalletDeleteOne{builder}
}

// Query returns a query builder for Wallet.
func (c *WalletClient) Query() *WalletQuery {
	return &WalletQuery{
		config: c.config,
	}
}

// Get returns a Wallet entity by its id.
func (c *WalletClient) Get(ctx context.Context, id int) (*Wallet, error) {
	return c.Query().Where(wallet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WalletClient) GetX(ctx context.Context, id int) *Wallet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySenders queries the senders edge of a Wallet.
func (c *WalletClient) QuerySenders(w *Wallet) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(wallet.Table, wallet.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, wallet.SendersTable, wallet.SendersColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRecipients queries the recipients edge of a Wallet.
func (c *WalletClient) QueryRecipients(w *Wallet) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(wallet.Table, wallet.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, wallet.RecipientsTable, wallet.RecipientsColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WalletClient) Hooks() []Hook {
	return c.hooks.Wallet
}

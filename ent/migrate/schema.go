// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "datetime", Type: field.TypeTime},
		{Name: "amount", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(6,2)", "postgres": "numeric(6,2)"}},
		{Name: "sender_id", Type: field.TypeInt},
		{Name: "recipient_id", Type: field.TypeInt},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_wallets_senders",
				Columns:    []*schema.Column{TransactionsColumns[3]},
				RefColumns: []*schema.Column{WalletsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "transactions_wallets_recipients",
				Columns:    []*schema.Column{TransactionsColumns[4]},
				RefColumns: []*schema.Column{WalletsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "transaction_datetime",
				Unique:  true,
				Columns: []*schema.Column{TransactionsColumns[1]},
			},
		},
	}
	// WalletsColumns holds the columns for the "wallets" table.
	WalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// WalletsTable holds the schema information for the "wallets" table.
	WalletsTable = &schema.Table{
		Name:       "wallets",
		Columns:    WalletsColumns,
		PrimaryKey: []*schema.Column{WalletsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TransactionsTable,
		WalletsTable,
	}
)

func init() {
	TransactionsTable.ForeignKeys[0].RefTable = WalletsTable
	TransactionsTable.ForeignKeys[1].RefTable = WalletsTable
}

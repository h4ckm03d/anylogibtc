package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/shopspring/decimal"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Time("datetime").
			Default(time.Now().UTC()).Immutable(),
		field.Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric(6,2)",
			}),
		field.Int("sender_id"),
		field.Int("recipient_id"),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("sender", Wallet.Type).
			Ref("senders").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique().Field("sender_id").Required(),
		edge.From("recipient", Wallet.Type).
			Ref("recipients").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique().Field("recipient_id").Required(),
	}
}

func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("datetime").Unique(),
	}
}

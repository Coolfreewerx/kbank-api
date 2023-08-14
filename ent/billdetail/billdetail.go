// Code generated by ent, DO NOT EDIT.

package billdetail

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the billdetail type in the database.
	Label = "bill_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChannelCode holds the string denoting the channel_code field in the database.
	FieldChannelCode = "channel_code"
	// FieldSenderBankCode holds the string denoting the sender_bank_code field in the database.
	FieldSenderBankCode = "sender_bank_code"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldTranAmount holds the string denoting the tran_amount field in the database.
	FieldTranAmount = "tran_amount"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeBills holds the string denoting the bills edge name in mutations.
	EdgeBills = "bills"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// Table holds the table name of the billdetail in the database.
	Table = "bill_details"
	// BillsTable is the table that holds the bills relation/edge.
	BillsTable = "bill_details"
	// BillsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillsInverseTable = "bills"
	// BillsColumn is the table column denoting the bills relation/edge.
	BillsColumn = "bill_bill_detail"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "bill_details"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_id"
)

// Columns holds all SQL columns for billdetail fields.
var Columns = []string{
	FieldID,
	FieldChannelCode,
	FieldSenderBankCode,
	FieldStatus,
	FieldCustomerID,
	FieldTranAmount,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "bill_details"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"bill_bill_detail",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultChannelCode holds the default value on creation for the "channel_code" field.
	DefaultChannelCode string
	// DefaultSenderBankCode holds the default value on creation for the "sender_bank_code" field.
	DefaultSenderBankCode string
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(string) error
	// DefaultTranAmount holds the default value on creation for the "tran_amount" field.
	DefaultTranAmount float64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the BillDetail queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChannelCode orders the results by the channel_code field.
func ByChannelCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannelCode, opts...).ToFunc()
}

// BySenderBankCode orders the results by the sender_bank_code field.
func BySenderBankCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSenderBankCode, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCustomerID orders the results by the customer_id field.
func ByCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerID, opts...).ToFunc()
}

// ByTranAmount orders the results by the tran_amount field.
func ByTranAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTranAmount, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByBillsField orders the results by bills field.
func ByBillsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBillsStep(), sql.OrderByField(field, opts...))
	}
}

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}
func newBillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, BillsTable, BillsColumn),
	)
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
	)
}
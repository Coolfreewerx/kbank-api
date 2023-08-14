// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kapi/ent/bill"
	"kapi/ent/billdetail"
	"kapi/ent/customer"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BillDetail is the model entity for the BillDetail schema.
type BillDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ChannelCode holds the value of the "channel_code" field.
	ChannelCode string `json:"channel_code,omitempty"`
	// SenderBankCode holds the value of the "sender_bank_code" field.
	SenderBankCode string `json:"sender_bank_code,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID int `json:"customer_id,omitempty"`
	// TranAmount holds the value of the "tran_amount" field.
	TranAmount float64 `json:"tran_amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillDetailQuery when eager-loading is set.
	Edges            BillDetailEdges `json:"edges"`
	bill_bill_detail *int
	selectValues     sql.SelectValues
}

// BillDetailEdges holds the relations/edges for other nodes in the graph.
type BillDetailEdges struct {
	// Bills holds the value of the bills edge.
	Bills *Bill `json:"bills,omitempty"`
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BillsOrErr returns the Bills value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillDetailEdges) BillsOrErr() (*Bill, error) {
	if e.loadedTypes[0] {
		if e.Bills == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: bill.Label}
		}
		return e.Bills, nil
	}
	return nil, &NotLoadedError{edge: "bills"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillDetailEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[1] {
		if e.Customer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BillDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case billdetail.FieldTranAmount:
			values[i] = new(sql.NullFloat64)
		case billdetail.FieldID, billdetail.FieldCustomerID:
			values[i] = new(sql.NullInt64)
		case billdetail.FieldChannelCode, billdetail.FieldSenderBankCode, billdetail.FieldStatus:
			values[i] = new(sql.NullString)
		case billdetail.FieldCreatedAt, billdetail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case billdetail.ForeignKeys[0]: // bill_bill_detail
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BillDetail fields.
func (bd *BillDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case billdetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bd.ID = int(value.Int64)
		case billdetail.FieldChannelCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_code", values[i])
			} else if value.Valid {
				bd.ChannelCode = value.String
			}
		case billdetail.FieldSenderBankCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender_bank_code", values[i])
			} else if value.Valid {
				bd.SenderBankCode = value.String
			}
		case billdetail.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				bd.Status = value.String
			}
		case billdetail.FieldCustomerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				bd.CustomerID = int(value.Int64)
			}
		case billdetail.FieldTranAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field tran_amount", values[i])
			} else if value.Valid {
				bd.TranAmount = value.Float64
			}
		case billdetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bd.CreatedAt = value.Time
			}
		case billdetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bd.UpdatedAt = value.Time
			}
		case billdetail.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field bill_bill_detail", value)
			} else if value.Valid {
				bd.bill_bill_detail = new(int)
				*bd.bill_bill_detail = int(value.Int64)
			}
		default:
			bd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BillDetail.
// This includes values selected through modifiers, order, etc.
func (bd *BillDetail) Value(name string) (ent.Value, error) {
	return bd.selectValues.Get(name)
}

// QueryBills queries the "bills" edge of the BillDetail entity.
func (bd *BillDetail) QueryBills() *BillQuery {
	return NewBillDetailClient(bd.config).QueryBills(bd)
}

// QueryCustomer queries the "customer" edge of the BillDetail entity.
func (bd *BillDetail) QueryCustomer() *CustomerQuery {
	return NewBillDetailClient(bd.config).QueryCustomer(bd)
}

// Update returns a builder for updating this BillDetail.
// Note that you need to call BillDetail.Unwrap() before calling this method if this BillDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (bd *BillDetail) Update() *BillDetailUpdateOne {
	return NewBillDetailClient(bd.config).UpdateOne(bd)
}

// Unwrap unwraps the BillDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bd *BillDetail) Unwrap() *BillDetail {
	_tx, ok := bd.config.driver.(*txDriver)
	if !ok {
		panic("ent: BillDetail is not a transactional entity")
	}
	bd.config.driver = _tx.drv
	return bd
}

// String implements the fmt.Stringer.
func (bd *BillDetail) String() string {
	var builder strings.Builder
	builder.WriteString("BillDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bd.ID))
	builder.WriteString("channel_code=")
	builder.WriteString(bd.ChannelCode)
	builder.WriteString(", ")
	builder.WriteString("sender_bank_code=")
	builder.WriteString(bd.SenderBankCode)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(bd.Status)
	builder.WriteString(", ")
	builder.WriteString("customer_id=")
	builder.WriteString(fmt.Sprintf("%v", bd.CustomerID))
	builder.WriteString(", ")
	builder.WriteString("tran_amount=")
	builder.WriteString(fmt.Sprintf("%v", bd.TranAmount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(bd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bd.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BillDetails is a parsable slice of BillDetail.
type BillDetails []*BillDetail
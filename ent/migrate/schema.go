// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BillsColumns holds the columns for the "bills" table.
	BillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ref_2", Type: field.TypeInt, Nullable: true},
		{Name: "ref_1", Type: field.TypeInt, Nullable: true},
		{Name: "biller_id", Type: field.TypeInt, Nullable: true},
	}
	// BillsTable holds the schema information for the "bills" table.
	BillsTable = &schema.Table{
		Name:       "bills",
		Columns:    BillsColumns,
		PrimaryKey: []*schema.Column{BillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bills_customers_bills",
				Columns:    []*schema.Column{BillsColumns[2]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bills_stores_bills",
				Columns:    []*schema.Column{BillsColumns[3]},
				RefColumns: []*schema.Column{StoresColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BillDetailsColumns holds the columns for the "bill_details" table.
	BillDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "channel_code", Type: field.TypeString, Default: ""},
		{Name: "sender_bank_code", Type: field.TypeString, Default: ""},
		{Name: "status", Type: field.TypeString},
		{Name: "tran_amount", Type: field.TypeFloat64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "bill_bill_detail", Type: field.TypeInt, Unique: true},
		{Name: "customer_id", Type: field.TypeInt, Nullable: true},
	}
	// BillDetailsTable holds the schema information for the "bill_details" table.
	BillDetailsTable = &schema.Table{
		Name:       "bill_details",
		Columns:    BillDetailsColumns,
		PrimaryKey: []*schema.Column{BillDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bill_details_bills_bill_detail",
				Columns:    []*schema.Column{BillDetailsColumns[7]},
				RefColumns: []*schema.Column{BillsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "bill_details_customers_bill_details",
				Columns:    []*schema.Column{BillDetailsColumns[8]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "first_name", Type: field.TypeString, Default: "Unknown"},
		{Name: "last_name", Type: field.TypeString, Default: ""},
		{Name: "title_name", Type: field.TypeString, Default: ""},
		{Name: "mobile_number", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
	}
	// StoresColumns holds the columns for the "stores" table.
	StoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "account_name", Type: field.TypeString, Default: ""},
		{Name: "service_name", Type: field.TypeString, Default: ""},
	}
	// StoresTable holds the schema information for the "stores" table.
	StoresTable = &schema.Table{
		Name:       "stores",
		Columns:    StoresColumns,
		PrimaryKey: []*schema.Column{StoresColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BillsTable,
		BillDetailsTable,
		CustomersTable,
		StoresTable,
	}
)

func init() {
	BillsTable.ForeignKeys[0].RefTable = CustomersTable
	BillsTable.ForeignKeys[1].RefTable = StoresTable
	BillDetailsTable.ForeignKeys[0].RefTable = BillsTable
	BillDetailsTable.ForeignKeys[1].RefTable = CustomersTable
}

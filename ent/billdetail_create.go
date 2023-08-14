// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kapi/ent/bill"
	"kapi/ent/billdetail"
	"kapi/ent/customer"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BillDetailCreate is the builder for creating a BillDetail entity.
type BillDetailCreate struct {
	config
	mutation *BillDetailMutation
	hooks    []Hook
}

// SetChannelCode sets the "channel_code" field.
func (bdc *BillDetailCreate) SetChannelCode(s string) *BillDetailCreate {
	bdc.mutation.SetChannelCode(s)
	return bdc
}

// SetNillableChannelCode sets the "channel_code" field if the given value is not nil.
func (bdc *BillDetailCreate) SetNillableChannelCode(s *string) *BillDetailCreate {
	if s != nil {
		bdc.SetChannelCode(*s)
	}
	return bdc
}

// SetSenderBankCode sets the "sender_bank_code" field.
func (bdc *BillDetailCreate) SetSenderBankCode(s string) *BillDetailCreate {
	bdc.mutation.SetSenderBankCode(s)
	return bdc
}

// SetNillableSenderBankCode sets the "sender_bank_code" field if the given value is not nil.
func (bdc *BillDetailCreate) SetNillableSenderBankCode(s *string) *BillDetailCreate {
	if s != nil {
		bdc.SetSenderBankCode(*s)
	}
	return bdc
}

// SetStatus sets the "status" field.
func (bdc *BillDetailCreate) SetStatus(s string) *BillDetailCreate {
	bdc.mutation.SetStatus(s)
	return bdc
}

// SetCustomerID sets the "customer_id" field.
func (bdc *BillDetailCreate) SetCustomerID(i int) *BillDetailCreate {
	bdc.mutation.SetCustomerID(i)
	return bdc
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (bdc *BillDetailCreate) SetNillableCustomerID(i *int) *BillDetailCreate {
	if i != nil {
		bdc.SetCustomerID(*i)
	}
	return bdc
}

// SetTranAmount sets the "tran_amount" field.
func (bdc *BillDetailCreate) SetTranAmount(f float64) *BillDetailCreate {
	bdc.mutation.SetTranAmount(f)
	return bdc
}

// SetNillableTranAmount sets the "tran_amount" field if the given value is not nil.
func (bdc *BillDetailCreate) SetNillableTranAmount(f *float64) *BillDetailCreate {
	if f != nil {
		bdc.SetTranAmount(*f)
	}
	return bdc
}

// SetCreatedAt sets the "created_at" field.
func (bdc *BillDetailCreate) SetCreatedAt(t time.Time) *BillDetailCreate {
	bdc.mutation.SetCreatedAt(t)
	return bdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bdc *BillDetailCreate) SetNillableCreatedAt(t *time.Time) *BillDetailCreate {
	if t != nil {
		bdc.SetCreatedAt(*t)
	}
	return bdc
}

// SetUpdatedAt sets the "updated_at" field.
func (bdc *BillDetailCreate) SetUpdatedAt(t time.Time) *BillDetailCreate {
	bdc.mutation.SetUpdatedAt(t)
	return bdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bdc *BillDetailCreate) SetNillableUpdatedAt(t *time.Time) *BillDetailCreate {
	if t != nil {
		bdc.SetUpdatedAt(*t)
	}
	return bdc
}

// SetBillsID sets the "bills" edge to the Bill entity by ID.
func (bdc *BillDetailCreate) SetBillsID(id int) *BillDetailCreate {
	bdc.mutation.SetBillsID(id)
	return bdc
}

// SetBills sets the "bills" edge to the Bill entity.
func (bdc *BillDetailCreate) SetBills(b *Bill) *BillDetailCreate {
	return bdc.SetBillsID(b.ID)
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (bdc *BillDetailCreate) SetCustomer(c *Customer) *BillDetailCreate {
	return bdc.SetCustomerID(c.ID)
}

// Mutation returns the BillDetailMutation object of the builder.
func (bdc *BillDetailCreate) Mutation() *BillDetailMutation {
	return bdc.mutation
}

// Save creates the BillDetail in the database.
func (bdc *BillDetailCreate) Save(ctx context.Context) (*BillDetail, error) {
	bdc.defaults()
	return withHooks(ctx, bdc.sqlSave, bdc.mutation, bdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bdc *BillDetailCreate) SaveX(ctx context.Context) *BillDetail {
	v, err := bdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bdc *BillDetailCreate) Exec(ctx context.Context) error {
	_, err := bdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdc *BillDetailCreate) ExecX(ctx context.Context) {
	if err := bdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bdc *BillDetailCreate) defaults() {
	if _, ok := bdc.mutation.ChannelCode(); !ok {
		v := billdetail.DefaultChannelCode
		bdc.mutation.SetChannelCode(v)
	}
	if _, ok := bdc.mutation.SenderBankCode(); !ok {
		v := billdetail.DefaultSenderBankCode
		bdc.mutation.SetSenderBankCode(v)
	}
	if _, ok := bdc.mutation.TranAmount(); !ok {
		v := billdetail.DefaultTranAmount
		bdc.mutation.SetTranAmount(v)
	}
	if _, ok := bdc.mutation.CreatedAt(); !ok {
		v := billdetail.DefaultCreatedAt()
		bdc.mutation.SetCreatedAt(v)
	}
	if _, ok := bdc.mutation.UpdatedAt(); !ok {
		v := billdetail.DefaultUpdatedAt()
		bdc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bdc *BillDetailCreate) check() error {
	if _, ok := bdc.mutation.ChannelCode(); !ok {
		return &ValidationError{Name: "channel_code", err: errors.New(`ent: missing required field "BillDetail.channel_code"`)}
	}
	if _, ok := bdc.mutation.SenderBankCode(); !ok {
		return &ValidationError{Name: "sender_bank_code", err: errors.New(`ent: missing required field "BillDetail.sender_bank_code"`)}
	}
	if _, ok := bdc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "BillDetail.status"`)}
	}
	if v, ok := bdc.mutation.Status(); ok {
		if err := billdetail.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "BillDetail.status": %w`, err)}
		}
	}
	if _, ok := bdc.mutation.TranAmount(); !ok {
		return &ValidationError{Name: "tran_amount", err: errors.New(`ent: missing required field "BillDetail.tran_amount"`)}
	}
	if _, ok := bdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BillDetail.created_at"`)}
	}
	if _, ok := bdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BillDetail.updated_at"`)}
	}
	if _, ok := bdc.mutation.BillsID(); !ok {
		return &ValidationError{Name: "bills", err: errors.New(`ent: missing required edge "BillDetail.bills"`)}
	}
	return nil
}

func (bdc *BillDetailCreate) sqlSave(ctx context.Context) (*BillDetail, error) {
	if err := bdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bdc.mutation.id = &_node.ID
	bdc.mutation.done = true
	return _node, nil
}

func (bdc *BillDetailCreate) createSpec() (*BillDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &BillDetail{config: bdc.config}
		_spec = sqlgraph.NewCreateSpec(billdetail.Table, sqlgraph.NewFieldSpec(billdetail.FieldID, field.TypeInt))
	)
	if value, ok := bdc.mutation.ChannelCode(); ok {
		_spec.SetField(billdetail.FieldChannelCode, field.TypeString, value)
		_node.ChannelCode = value
	}
	if value, ok := bdc.mutation.SenderBankCode(); ok {
		_spec.SetField(billdetail.FieldSenderBankCode, field.TypeString, value)
		_node.SenderBankCode = value
	}
	if value, ok := bdc.mutation.Status(); ok {
		_spec.SetField(billdetail.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := bdc.mutation.TranAmount(); ok {
		_spec.SetField(billdetail.FieldTranAmount, field.TypeFloat64, value)
		_node.TranAmount = value
	}
	if value, ok := bdc.mutation.CreatedAt(); ok {
		_spec.SetField(billdetail.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bdc.mutation.UpdatedAt(); ok {
		_spec.SetField(billdetail.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := bdc.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   billdetail.BillsTable,
			Columns: []string{billdetail.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.bill_bill_detail = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bdc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   billdetail.CustomerTable,
			Columns: []string{billdetail.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CustomerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BillDetailCreateBulk is the builder for creating many BillDetail entities in bulk.
type BillDetailCreateBulk struct {
	config
	builders []*BillDetailCreate
}

// Save creates the BillDetail entities in the database.
func (bdcb *BillDetailCreateBulk) Save(ctx context.Context) ([]*BillDetail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bdcb.builders))
	nodes := make([]*BillDetail, len(bdcb.builders))
	mutators := make([]Mutator, len(bdcb.builders))
	for i := range bdcb.builders {
		func(i int, root context.Context) {
			builder := bdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillDetailMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bdcb *BillDetailCreateBulk) SaveX(ctx context.Context) []*BillDetail {
	v, err := bdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bdcb *BillDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := bdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdcb *BillDetailCreateBulk) ExecX(ctx context.Context) {
	if err := bdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

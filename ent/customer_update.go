// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kapi/ent/bill"
	"kapi/ent/customer"
	"kapi/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetFirstName sets the "first_name" field.
func (cu *CustomerUpdate) SetFirstName(s string) *CustomerUpdate {
	cu.mutation.SetFirstName(s)
	return cu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableFirstName(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetFirstName(*s)
	}
	return cu
}

// SetLastName sets the "last_name" field.
func (cu *CustomerUpdate) SetLastName(s string) *CustomerUpdate {
	cu.mutation.SetLastName(s)
	return cu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableLastName(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetLastName(*s)
	}
	return cu
}

// SetTitleName sets the "title_name" field.
func (cu *CustomerUpdate) SetTitleName(s string) *CustomerUpdate {
	cu.mutation.SetTitleName(s)
	return cu
}

// SetNillableTitleName sets the "title_name" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableTitleName(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetTitleName(*s)
	}
	return cu
}

// SetMobileNumber sets the "mobile_number" field.
func (cu *CustomerUpdate) SetMobileNumber(s string) *CustomerUpdate {
	cu.mutation.SetMobileNumber(s)
	return cu
}

// SetNillableMobileNumber sets the "mobile_number" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableMobileNumber(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetMobileNumber(*s)
	}
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CustomerUpdate) SetCreatedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableCreatedAt(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (cu *CustomerUpdate) AddBillIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddBillIDs(ids...)
	return cu
}

// AddBills adds the "bills" edges to the Bill entity.
func (cu *CustomerUpdate) AddBills(b ...*Bill) *CustomerUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.AddBillIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (cu *CustomerUpdate) ClearBills() *CustomerUpdate {
	cu.mutation.ClearBills()
	return cu
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (cu *CustomerUpdate) RemoveBillIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveBillIDs(ids...)
	return cu
}

// RemoveBills removes "bills" edges to Bill entities.
func (cu *CustomerUpdate) RemoveBills(b ...*Bill) *CustomerUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.RemoveBillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.FirstName(); ok {
		_spec.SetField(customer.FieldFirstName, field.TypeString, value)
	}
	if value, ok := cu.mutation.LastName(); ok {
		_spec.SetField(customer.FieldLastName, field.TypeString, value)
	}
	if value, ok := cu.mutation.TitleName(); ok {
		_spec.SetField(customer.FieldTitleName, field.TypeString, value)
	}
	if value, ok := cu.mutation.MobileNumber(); ok {
		_spec.SetField(customer.FieldMobileNumber, field.TypeString, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(customer.FieldCreatedAt, field.TypeTime, value)
	}
	if cu.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.BillsTable,
			Columns: []string{customer.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedBillsIDs(); len(nodes) > 0 && !cu.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.BillsTable,
			Columns: []string{customer.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.BillsTable,
			Columns: []string{customer.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomerMutation
}

// SetFirstName sets the "first_name" field.
func (cuo *CustomerUpdateOne) SetFirstName(s string) *CustomerUpdateOne {
	cuo.mutation.SetFirstName(s)
	return cuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableFirstName(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetFirstName(*s)
	}
	return cuo
}

// SetLastName sets the "last_name" field.
func (cuo *CustomerUpdateOne) SetLastName(s string) *CustomerUpdateOne {
	cuo.mutation.SetLastName(s)
	return cuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableLastName(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetLastName(*s)
	}
	return cuo
}

// SetTitleName sets the "title_name" field.
func (cuo *CustomerUpdateOne) SetTitleName(s string) *CustomerUpdateOne {
	cuo.mutation.SetTitleName(s)
	return cuo
}

// SetNillableTitleName sets the "title_name" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableTitleName(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetTitleName(*s)
	}
	return cuo
}

// SetMobileNumber sets the "mobile_number" field.
func (cuo *CustomerUpdateOne) SetMobileNumber(s string) *CustomerUpdateOne {
	cuo.mutation.SetMobileNumber(s)
	return cuo
}

// SetNillableMobileNumber sets the "mobile_number" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableMobileNumber(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetMobileNumber(*s)
	}
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CustomerUpdateOne) SetCreatedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableCreatedAt(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (cuo *CustomerUpdateOne) AddBillIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddBillIDs(ids...)
	return cuo
}

// AddBills adds the "bills" edges to the Bill entity.
func (cuo *CustomerUpdateOne) AddBills(b ...*Bill) *CustomerUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.AddBillIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (cuo *CustomerUpdateOne) ClearBills() *CustomerUpdateOne {
	cuo.mutation.ClearBills()
	return cuo
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (cuo *CustomerUpdateOne) RemoveBillIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveBillIDs(ids...)
	return cuo
}

// RemoveBills removes "bills" edges to Bill entities.
func (cuo *CustomerUpdateOne) RemoveBills(b ...*Bill) *CustomerUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.RemoveBillIDs(ids...)
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cuo *CustomerUpdateOne) Where(ps ...predicate.Customer) *CustomerUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CustomerUpdateOne) Select(field string, fields ...string) *CustomerUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Customer entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Customer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for _, f := range fields {
			if !customer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != customer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.FirstName(); ok {
		_spec.SetField(customer.FieldFirstName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.LastName(); ok {
		_spec.SetField(customer.FieldLastName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.TitleName(); ok {
		_spec.SetField(customer.FieldTitleName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.MobileNumber(); ok {
		_spec.SetField(customer.FieldMobileNumber, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(customer.FieldCreatedAt, field.TypeTime, value)
	}
	if cuo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.BillsTable,
			Columns: []string{customer.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedBillsIDs(); len(nodes) > 0 && !cuo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.BillsTable,
			Columns: []string{customer.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.BillsTable,
			Columns: []string{customer.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}

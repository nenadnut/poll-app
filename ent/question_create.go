// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"poll-app/ent/poll"
	"poll-app/ent/question"
	"poll-app/ent/questionoption"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QuestionCreate is the builder for creating a Question entity.
type QuestionCreate struct {
	config
	mutation *QuestionMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (qc *QuestionCreate) SetText(s string) *QuestionCreate {
	qc.mutation.SetText(s)
	return qc
}

// SetHead sets the "head" field.
func (qc *QuestionCreate) SetHead(b bool) *QuestionCreate {
	qc.mutation.SetHead(b)
	return qc
}

// SetNillableHead sets the "head" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableHead(b *bool) *QuestionCreate {
	if b != nil {
		qc.SetHead(*b)
	}
	return qc
}

// SetNumOfAnswers sets the "num_of_answers" field.
func (qc *QuestionCreate) SetNumOfAnswers(i int) *QuestionCreate {
	qc.mutation.SetNumOfAnswers(i)
	return qc
}

// SetNillableNumOfAnswers sets the "num_of_answers" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableNumOfAnswers(i *int) *QuestionCreate {
	if i != nil {
		qc.SetNumOfAnswers(*i)
	}
	return qc
}

// SetCreatedAt sets the "created_at" field.
func (qc *QuestionCreate) SetCreatedAt(t time.Time) *QuestionCreate {
	qc.mutation.SetCreatedAt(t)
	return qc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableCreatedAt(t *time.Time) *QuestionCreate {
	if t != nil {
		qc.SetCreatedAt(*t)
	}
	return qc
}

// SetUpdatedAt sets the "updated_at" field.
func (qc *QuestionCreate) SetUpdatedAt(t time.Time) *QuestionCreate {
	qc.mutation.SetUpdatedAt(t)
	return qc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableUpdatedAt(t *time.Time) *QuestionCreate {
	if t != nil {
		qc.SetUpdatedAt(*t)
	}
	return qc
}

// SetPollID sets the "poll_id" field.
func (qc *QuestionCreate) SetPollID(i int) *QuestionCreate {
	qc.mutation.SetPollID(i)
	return qc
}

// AddOptionIDs adds the "options" edge to the QuestionOption entity by IDs.
func (qc *QuestionCreate) AddOptionIDs(ids ...int) *QuestionCreate {
	qc.mutation.AddOptionIDs(ids...)
	return qc
}

// AddOptions adds the "options" edges to the QuestionOption entity.
func (qc *QuestionCreate) AddOptions(q ...*QuestionOption) *QuestionCreate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qc.AddOptionIDs(ids...)
}

// AddNextQuestionInvIDs adds the "next_question_inv" edge to the Question entity by IDs.
func (qc *QuestionCreate) AddNextQuestionInvIDs(ids ...int) *QuestionCreate {
	qc.mutation.AddNextQuestionInvIDs(ids...)
	return qc
}

// AddNextQuestionInv adds the "next_question_inv" edges to the Question entity.
func (qc *QuestionCreate) AddNextQuestionInv(q ...*Question) *QuestionCreate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qc.AddNextQuestionInvIDs(ids...)
}

// SetNextQuestionID sets the "next_question" edge to the Question entity by ID.
func (qc *QuestionCreate) SetNextQuestionID(id int) *QuestionCreate {
	qc.mutation.SetNextQuestionID(id)
	return qc
}

// SetNillableNextQuestionID sets the "next_question" edge to the Question entity by ID if the given value is not nil.
func (qc *QuestionCreate) SetNillableNextQuestionID(id *int) *QuestionCreate {
	if id != nil {
		qc = qc.SetNextQuestionID(*id)
	}
	return qc
}

// SetNextQuestion sets the "next_question" edge to the Question entity.
func (qc *QuestionCreate) SetNextQuestion(q *Question) *QuestionCreate {
	return qc.SetNextQuestionID(q.ID)
}

// SetPoll sets the "poll" edge to the Poll entity.
func (qc *QuestionCreate) SetPoll(p *Poll) *QuestionCreate {
	return qc.SetPollID(p.ID)
}

// Mutation returns the QuestionMutation object of the builder.
func (qc *QuestionCreate) Mutation() *QuestionMutation {
	return qc.mutation
}

// Save creates the Question in the database.
func (qc *QuestionCreate) Save(ctx context.Context) (*Question, error) {
	qc.defaults()
	return withHooks(ctx, qc.sqlSave, qc.mutation, qc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QuestionCreate) SaveX(ctx context.Context) *Question {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qc *QuestionCreate) Exec(ctx context.Context) error {
	_, err := qc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qc *QuestionCreate) ExecX(ctx context.Context) {
	if err := qc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qc *QuestionCreate) defaults() {
	if _, ok := qc.mutation.Head(); !ok {
		v := question.DefaultHead
		qc.mutation.SetHead(v)
	}
	if _, ok := qc.mutation.NumOfAnswers(); !ok {
		v := question.DefaultNumOfAnswers
		qc.mutation.SetNumOfAnswers(v)
	}
	if _, ok := qc.mutation.CreatedAt(); !ok {
		v := question.DefaultCreatedAt()
		qc.mutation.SetCreatedAt(v)
	}
	if _, ok := qc.mutation.UpdatedAt(); !ok {
		v := question.DefaultUpdatedAt()
		qc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qc *QuestionCreate) check() error {
	if _, ok := qc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Question.text"`)}
	}
	if _, ok := qc.mutation.Head(); !ok {
		return &ValidationError{Name: "head", err: errors.New(`ent: missing required field "Question.head"`)}
	}
	if _, ok := qc.mutation.NumOfAnswers(); !ok {
		return &ValidationError{Name: "num_of_answers", err: errors.New(`ent: missing required field "Question.num_of_answers"`)}
	}
	if _, ok := qc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Question.created_at"`)}
	}
	if _, ok := qc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Question.updated_at"`)}
	}
	if _, ok := qc.mutation.PollID(); !ok {
		return &ValidationError{Name: "poll_id", err: errors.New(`ent: missing required field "Question.poll_id"`)}
	}
	if _, ok := qc.mutation.PollID(); !ok {
		return &ValidationError{Name: "poll", err: errors.New(`ent: missing required edge "Question.poll"`)}
	}
	return nil
}

func (qc *QuestionCreate) sqlSave(ctx context.Context) (*Question, error) {
	if err := qc.check(); err != nil {
		return nil, err
	}
	_node, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	qc.mutation.id = &_node.ID
	qc.mutation.done = true
	return _node, nil
}

func (qc *QuestionCreate) createSpec() (*Question, *sqlgraph.CreateSpec) {
	var (
		_node = &Question{config: qc.config}
		_spec = sqlgraph.NewCreateSpec(question.Table, sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt))
	)
	if value, ok := qc.mutation.Text(); ok {
		_spec.SetField(question.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := qc.mutation.Head(); ok {
		_spec.SetField(question.FieldHead, field.TypeBool, value)
		_node.Head = value
	}
	if value, ok := qc.mutation.NumOfAnswers(); ok {
		_spec.SetField(question.FieldNumOfAnswers, field.TypeInt, value)
		_node.NumOfAnswers = value
	}
	if value, ok := qc.mutation.CreatedAt(); ok {
		_spec.SetField(question.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := qc.mutation.UpdatedAt(); ok {
		_spec.SetField(question.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := qc.mutation.OptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.OptionsTable,
			Columns: []string{question.OptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(questionoption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.NextQuestionInvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   question.NextQuestionInvTable,
			Columns: []string{question.NextQuestionInvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.NextQuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   question.NextQuestionTable,
			Columns: []string{question.NextQuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.question_next_question = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.PollIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   question.PollTable,
			Columns: []string{question.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PollID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// QuestionCreateBulk is the builder for creating many Question entities in bulk.
type QuestionCreateBulk struct {
	config
	err      error
	builders []*QuestionCreate
}

// Save creates the Question entities in the database.
func (qcb *QuestionCreateBulk) Save(ctx context.Context) ([]*Question, error) {
	if qcb.err != nil {
		return nil, qcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(qcb.builders))
	nodes := make([]*Question, len(qcb.builders))
	mutators := make([]Mutator, len(qcb.builders))
	for i := range qcb.builders {
		func(i int, root context.Context) {
			builder := qcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QuestionMutation)
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
					_, err = mutators[i+1].Mutate(root, qcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, qcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qcb *QuestionCreateBulk) SaveX(ctx context.Context) []*Question {
	v, err := qcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qcb *QuestionCreateBulk) Exec(ctx context.Context) error {
	_, err := qcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcb *QuestionCreateBulk) ExecX(ctx context.Context) {
	if err := qcb.Exec(ctx); err != nil {
		panic(err)
	}
}

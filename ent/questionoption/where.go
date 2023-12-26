// Code generated by ent, DO NOT EDIT.

package questionoption

import (
	"poll-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLTE(FieldID, id))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldText, v))
}

// Chosen applies equality check predicate on the "chosen" field. It's identical to ChosenEQ.
func Chosen(v bool) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldChosen, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldUpdatedAt, v))
}

// QuestionID applies equality check predicate on the "question_id" field. It's identical to QuestionIDEQ.
func QuestionID(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldQuestionID, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldContainsFold(FieldText, v))
}

// ChosenEQ applies the EQ predicate on the "chosen" field.
func ChosenEQ(v bool) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldChosen, v))
}

// ChosenNEQ applies the NEQ predicate on the "chosen" field.
func ChosenNEQ(v bool) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldChosen, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLTE(FieldUpdatedAt, v))
}

// QuestionIDEQ applies the EQ predicate on the "question_id" field.
func QuestionIDEQ(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldQuestionID, v))
}

// QuestionIDNEQ applies the NEQ predicate on the "question_id" field.
func QuestionIDNEQ(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldQuestionID, v))
}

// QuestionIDIn applies the In predicate on the "question_id" field.
func QuestionIDIn(vs ...int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldIn(FieldQuestionID, vs...))
}

// QuestionIDNotIn applies the NotIn predicate on the "question_id" field.
func QuestionIDNotIn(vs ...int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNotIn(FieldQuestionID, vs...))
}

// HasNextOptionInv applies the HasEdge predicate on the "next_option_inv" edge.
func HasNextOptionInv() predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, NextOptionInvTable, NextOptionInvColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNextOptionInvWith applies the HasEdge predicate on the "next_option_inv" edge with a given conditions (other predicates).
func HasNextOptionInvWith(preds ...predicate.QuestionOption) predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := newNextOptionInvStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNextOption applies the HasEdge predicate on the "next_option" edge.
func HasNextOption() predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, NextOptionTable, NextOptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNextOptionWith applies the HasEdge predicate on the "next_option" edge with a given conditions (other predicates).
func HasNextOptionWith(preds ...predicate.QuestionOption) predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := newNextOptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasQuestion applies the HasEdge predicate on the "question" edge.
func HasQuestion() predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWith applies the HasEdge predicate on the "question" edge with a given conditions (other predicates).
func HasQuestionWith(preds ...predicate.Question) predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := newQuestionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.QuestionOption) predicate.QuestionOption {
	return predicate.QuestionOption(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.QuestionOption) predicate.QuestionOption {
	return predicate.QuestionOption(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.QuestionOption) predicate.QuestionOption {
	return predicate.QuestionOption(sql.NotPredicates(p))
}

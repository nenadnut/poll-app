// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePolls holds the string denoting the polls edge name in mutations.
	EdgePolls = "polls"
	// EdgeStartedPolls holds the string denoting the started_polls edge name in mutations.
	EdgeStartedPolls = "started_polls"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PollsTable is the table that holds the polls relation/edge.
	PollsTable = "polls"
	// PollsInverseTable is the table name for the Poll entity.
	// It exists in this package in order to avoid circular dependency with the "poll" package.
	PollsInverseTable = "polls"
	// PollsColumn is the table column denoting the polls relation/edge.
	PollsColumn = "creator_id"
	// StartedPollsTable is the table that holds the started_polls relation/edge.
	StartedPollsTable = "started_polls"
	// StartedPollsInverseTable is the table name for the StartedPoll entity.
	// It exists in this package in order to avoid circular dependency with the "startedpoll" package.
	StartedPollsInverseTable = "started_polls"
	// StartedPollsColumn is the table column denoting the started_polls relation/edge.
	StartedPollsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFirstName,
	FieldLastName,
	FieldEmail,
	FieldPassword,
	FieldRole,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// Role defines the type for the "role" enum field.
type Role string

// RoleRegular is the default value of the Role enum.
const DefaultRole = RoleRegular

// Role values.
const (
	RoleRegular Role = "regular"
	RoleAdmin   Role = "admin"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleRegular, RoleAdmin:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPollsCount orders the results by polls count.
func ByPollsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPollsStep(), opts...)
	}
}

// ByPolls orders the results by polls terms.
func ByPolls(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPollsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStartedPollsCount orders the results by started_polls count.
func ByStartedPollsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStartedPollsStep(), opts...)
	}
}

// ByStartedPolls orders the results by started_polls terms.
func ByStartedPolls(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStartedPollsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPollsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PollsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PollsTable, PollsColumn),
	)
}
func newStartedPollsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StartedPollsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StartedPollsTable, StartedPollsColumn),
	)
}

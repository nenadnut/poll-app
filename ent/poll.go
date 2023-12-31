// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"poll-app/ent/poll"
	"poll-app/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Poll is the model entity for the Poll schema.
type Poll struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"-"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"-"`
	// CreatorID holds the value of the "creator_id" field.
	CreatorID int `json:"creator_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PollQuery when eager-loading is set.
	Edges        PollEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PollEdges holds the relations/edges for other nodes in the graph.
type PollEdges struct {
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// Questions holds the value of the questions edge.
	Questions []*Question `json:"questions,omitempty"`
	// StartedPolls holds the value of the started_polls edge.
	StartedPolls []*StartedPoll `json:"started_polls,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PollEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// QuestionsOrErr returns the Questions value or an error if the edge
// was not loaded in eager-loading.
func (e PollEdges) QuestionsOrErr() ([]*Question, error) {
	if e.loadedTypes[1] {
		return e.Questions, nil
	}
	return nil, &NotLoadedError{edge: "questions"}
}

// StartedPollsOrErr returns the StartedPolls value or an error if the edge
// was not loaded in eager-loading.
func (e PollEdges) StartedPollsOrErr() ([]*StartedPoll, error) {
	if e.loadedTypes[2] {
		return e.StartedPolls, nil
	}
	return nil, &NotLoadedError{edge: "started_polls"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Poll) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case poll.FieldID, poll.FieldCreatorID:
			values[i] = new(sql.NullInt64)
		case poll.FieldTitle, poll.FieldDescription:
			values[i] = new(sql.NullString)
		case poll.FieldCreatedAt, poll.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Poll fields.
func (po *Poll) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case poll.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case poll.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				po.Title = value.String
			}
		case poll.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				po.Description = value.String
			}
		case poll.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case poll.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case poll.FieldCreatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				po.CreatorID = int(value.Int64)
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Poll.
// This includes values selected through modifiers, order, etc.
func (po *Poll) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryCreator queries the "creator" edge of the Poll entity.
func (po *Poll) QueryCreator() *UserQuery {
	return NewPollClient(po.config).QueryCreator(po)
}

// QueryQuestions queries the "questions" edge of the Poll entity.
func (po *Poll) QueryQuestions() *QuestionQuery {
	return NewPollClient(po.config).QueryQuestions(po)
}

// QueryStartedPolls queries the "started_polls" edge of the Poll entity.
func (po *Poll) QueryStartedPolls() *StartedPollQuery {
	return NewPollClient(po.config).QueryStartedPolls(po)
}

// Update returns a builder for updating this Poll.
// Note that you need to call Poll.Unwrap() before calling this method if this Poll
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Poll) Update() *PollUpdateOne {
	return NewPollClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Poll entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Poll) Unwrap() *Poll {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Poll is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Poll) String() string {
	var builder strings.Builder
	builder.WriteString("Poll(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("title=")
	builder.WriteString(po.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(po.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("creator_id=")
	builder.WriteString(fmt.Sprintf("%v", po.CreatorID))
	builder.WriteByte(')')
	return builder.String()
}

// Polls is a parsable slice of Poll.
type Polls []*Poll

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/subject"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Subject is the model entity for the Subject schema.
type Subject struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Summary holds the value of the "summary" field.
	Summary string `json:"summary,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Date holds the value of the "date" field.
	Date string `json:"date,omitempty"`
	// NameCn holds the value of the "name_cn" field.
	NameCn string `json:"name_cn,omitempty"`
	// OnHold holds the value of the "on_hold" field.
	OnHold uint32 `json:"on_hold,omitempty"`
	// Wish holds the value of the "wish" field.
	Wish uint32 `json:"wish,omitempty"`
	// Doing holds the value of the "doing" field.
	Doing uint32 `json:"doing,omitempty"`
	// SubjectType holds the value of the "subject_type" field.
	SubjectType uint8 `json:"subject_type,omitempty"`
	// Collect holds the value of the "collect" field.
	Collect      uint32 `json:"collect,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subject.FieldID, subject.FieldOnHold, subject.FieldWish, subject.FieldDoing, subject.FieldSubjectType, subject.FieldCollect:
			values[i] = new(sql.NullInt64)
		case subject.FieldImage, subject.FieldSummary, subject.FieldName, subject.FieldDate, subject.FieldNameCn:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subject fields.
func (s *Subject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case subject.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				s.Image = value.String
			}
		case subject.FieldSummary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field summary", values[i])
			} else if value.Valid {
				s.Summary = value.String
			}
		case subject.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subject.FieldDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				s.Date = value.String
			}
		case subject.FieldNameCn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_cn", values[i])
			} else if value.Valid {
				s.NameCn = value.String
			}
		case subject.FieldOnHold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field on_hold", values[i])
			} else if value.Valid {
				s.OnHold = uint32(value.Int64)
			}
		case subject.FieldWish:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wish", values[i])
			} else if value.Valid {
				s.Wish = uint32(value.Int64)
			}
		case subject.FieldDoing:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field doing", values[i])
			} else if value.Valid {
				s.Doing = uint32(value.Int64)
			}
		case subject.FieldSubjectType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field subject_type", values[i])
			} else if value.Valid {
				s.SubjectType = uint8(value.Int64)
			}
		case subject.FieldCollect:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field collect", values[i])
			} else if value.Valid {
				s.Collect = uint32(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Subject.
// This includes values selected through modifiers, order, etc.
func (s *Subject) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Subject.
// Note that you need to call Subject.Unwrap() before calling this method if this Subject
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subject) Update() *SubjectUpdateOne {
	return NewSubjectClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subject) Unwrap() *Subject {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subject is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subject) String() string {
	var builder strings.Builder
	builder.WriteString("Subject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("image=")
	builder.WriteString(s.Image)
	builder.WriteString(", ")
	builder.WriteString("summary=")
	builder.WriteString(s.Summary)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(s.Date)
	builder.WriteString(", ")
	builder.WriteString("name_cn=")
	builder.WriteString(s.NameCn)
	builder.WriteString(", ")
	builder.WriteString("on_hold=")
	builder.WriteString(fmt.Sprintf("%v", s.OnHold))
	builder.WriteString(", ")
	builder.WriteString("wish=")
	builder.WriteString(fmt.Sprintf("%v", s.Wish))
	builder.WriteString(", ")
	builder.WriteString("doing=")
	builder.WriteString(fmt.Sprintf("%v", s.Doing))
	builder.WriteString(", ")
	builder.WriteString("subject_type=")
	builder.WriteString(fmt.Sprintf("%v", s.SubjectType))
	builder.WriteString(", ")
	builder.WriteString("collect=")
	builder.WriteString(fmt.Sprintf("%v", s.Collect))
	builder.WriteByte(')')
	return builder.String()
}

// Subjects is a parsable slice of Subject.
type Subjects []*Subject

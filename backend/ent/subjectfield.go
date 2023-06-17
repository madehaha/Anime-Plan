// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/subject"
	"backend/ent/subjectfield"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// SubjectField is the model entity for the SubjectField schema.
type SubjectField struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Rate1 holds the value of the "rate_1" field.
	Rate1 uint32 `json:"rate_1,omitempty"`
	// Rate2 holds the value of the "rate_2" field.
	Rate2 uint32 `json:"rate_2,omitempty"`
	// Rate3 holds the value of the "rate_3" field.
	Rate3 uint32 `json:"rate_3,omitempty"`
	// Rate4 holds the value of the "rate_4" field.
	Rate4 uint32 `json:"rate_4,omitempty"`
	// Rate5 holds the value of the "rate_5" field.
	Rate5 uint32 `json:"rate_5,omitempty"`
	// Rate6 holds the value of the "rate_6" field.
	Rate6 uint32 `json:"rate_6,omitempty"`
	// Rate7 holds the value of the "rate_7" field.
	Rate7 uint32 `json:"rate_7,omitempty"`
	// Rate8 holds the value of the "rate_8" field.
	Rate8 uint32 `json:"rate_8,omitempty"`
	// Rate9 holds the value of the "rate_9" field.
	Rate9 uint32 `json:"rate_9,omitempty"`
	// Rate10 holds the value of the "rate_10" field.
	Rate10 uint32 `json:"rate_10,omitempty"`
	// AverageScore holds the value of the "average_score" field.
	AverageScore float64 `json:"average_score,omitempty"`
	// Rank holds the value of the "rank" field.
	Rank uint32 `json:"rank,omitempty"`
	// Year holds the value of the "year" field.
	Year uint32 `json:"year,omitempty"`
	// Month holds the value of the "month" field.
	Month uint8 `json:"month,omitempty"`
	// Date holds the value of the "date" field.
	Date uint8 `json:"date,omitempty"`
	// Weekday holds the value of the "weekday" field.
	Weekday uint8 `json:"weekday,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectFieldQuery when eager-loading is set.
	Edges                 SubjectFieldEdges `json:"edges"`
	subject_subject_field *uint32
	selectValues          sql.SelectValues
}

// SubjectFieldEdges holds the relations/edges for other nodes in the graph.
type SubjectFieldEdges struct {
	// Subject holds the value of the subject edge.
	Subject *Subject `json:"subject,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectFieldEdges) SubjectOrErr() (*Subject, error) {
	if e.loadedTypes[0] {
		if e.Subject == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: subject.Label}
		}
		return e.Subject, nil
	}
	return nil, &NotLoadedError{edge: "subject"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubjectField) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subjectfield.FieldAverageScore:
			values[i] = new(sql.NullFloat64)
		case subjectfield.FieldID, subjectfield.FieldRate1, subjectfield.FieldRate2, subjectfield.FieldRate3, subjectfield.FieldRate4, subjectfield.FieldRate5, subjectfield.FieldRate6, subjectfield.FieldRate7, subjectfield.FieldRate8, subjectfield.FieldRate9, subjectfield.FieldRate10, subjectfield.FieldRank, subjectfield.FieldYear, subjectfield.FieldMonth, subjectfield.FieldDate, subjectfield.FieldWeekday:
			values[i] = new(sql.NullInt64)
		case subjectfield.ForeignKeys[0]: // subject_subject_field
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubjectField fields.
func (sf *SubjectField) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subjectfield.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sf.ID = int(value.Int64)
		case subjectfield.FieldRate1:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_1", values[i])
			} else if value.Valid {
				sf.Rate1 = uint32(value.Int64)
			}
		case subjectfield.FieldRate2:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_2", values[i])
			} else if value.Valid {
				sf.Rate2 = uint32(value.Int64)
			}
		case subjectfield.FieldRate3:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_3", values[i])
			} else if value.Valid {
				sf.Rate3 = uint32(value.Int64)
			}
		case subjectfield.FieldRate4:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_4", values[i])
			} else if value.Valid {
				sf.Rate4 = uint32(value.Int64)
			}
		case subjectfield.FieldRate5:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_5", values[i])
			} else if value.Valid {
				sf.Rate5 = uint32(value.Int64)
			}
		case subjectfield.FieldRate6:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_6", values[i])
			} else if value.Valid {
				sf.Rate6 = uint32(value.Int64)
			}
		case subjectfield.FieldRate7:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_7", values[i])
			} else if value.Valid {
				sf.Rate7 = uint32(value.Int64)
			}
		case subjectfield.FieldRate8:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_8", values[i])
			} else if value.Valid {
				sf.Rate8 = uint32(value.Int64)
			}
		case subjectfield.FieldRate9:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_9", values[i])
			} else if value.Valid {
				sf.Rate9 = uint32(value.Int64)
			}
		case subjectfield.FieldRate10:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate_10", values[i])
			} else if value.Valid {
				sf.Rate10 = uint32(value.Int64)
			}
		case subjectfield.FieldAverageScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field average_score", values[i])
			} else if value.Valid {
				sf.AverageScore = value.Float64
			}
		case subjectfield.FieldRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				sf.Rank = uint32(value.Int64)
			}
		case subjectfield.FieldYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field year", values[i])
			} else if value.Valid {
				sf.Year = uint32(value.Int64)
			}
		case subjectfield.FieldMonth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field month", values[i])
			} else if value.Valid {
				sf.Month = uint8(value.Int64)
			}
		case subjectfield.FieldDate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				sf.Date = uint8(value.Int64)
			}
		case subjectfield.FieldWeekday:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weekday", values[i])
			} else if value.Valid {
				sf.Weekday = uint8(value.Int64)
			}
		case subjectfield.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field subject_subject_field", value)
			} else if value.Valid {
				sf.subject_subject_field = new(uint32)
				*sf.subject_subject_field = uint32(value.Int64)
			}
		default:
			sf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SubjectField.
// This includes values selected through modifiers, order, etc.
func (sf *SubjectField) Value(name string) (ent.Value, error) {
	return sf.selectValues.Get(name)
}

// QuerySubject queries the "subject" edge of the SubjectField entity.
func (sf *SubjectField) QuerySubject() *SubjectQuery {
	return NewSubjectFieldClient(sf.config).QuerySubject(sf)
}

// Update returns a builder for updating this SubjectField.
// Note that you need to call SubjectField.Unwrap() before calling this method if this SubjectField
// was returned from a transaction, and the transaction was committed or rolled back.
func (sf *SubjectField) Update() *SubjectFieldUpdateOne {
	return NewSubjectFieldClient(sf.config).UpdateOne(sf)
}

// Unwrap unwraps the SubjectField entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sf *SubjectField) Unwrap() *SubjectField {
	_tx, ok := sf.config.driver.(*txDriver)
	if !ok {
		panic("ent: SubjectField is not a transactional entity")
	}
	sf.config.driver = _tx.drv
	return sf
}

// String implements the fmt.Stringer.
func (sf *SubjectField) String() string {
	var builder strings.Builder
	builder.WriteString("SubjectField(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sf.ID))
	builder.WriteString("rate_1=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate1))
	builder.WriteString(", ")
	builder.WriteString("rate_2=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate2))
	builder.WriteString(", ")
	builder.WriteString("rate_3=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate3))
	builder.WriteString(", ")
	builder.WriteString("rate_4=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate4))
	builder.WriteString(", ")
	builder.WriteString("rate_5=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate5))
	builder.WriteString(", ")
	builder.WriteString("rate_6=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate6))
	builder.WriteString(", ")
	builder.WriteString("rate_7=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate7))
	builder.WriteString(", ")
	builder.WriteString("rate_8=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate8))
	builder.WriteString(", ")
	builder.WriteString("rate_9=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate9))
	builder.WriteString(", ")
	builder.WriteString("rate_10=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rate10))
	builder.WriteString(", ")
	builder.WriteString("average_score=")
	builder.WriteString(fmt.Sprintf("%v", sf.AverageScore))
	builder.WriteString(", ")
	builder.WriteString("rank=")
	builder.WriteString(fmt.Sprintf("%v", sf.Rank))
	builder.WriteString(", ")
	builder.WriteString("year=")
	builder.WriteString(fmt.Sprintf("%v", sf.Year))
	builder.WriteString(", ")
	builder.WriteString("month=")
	builder.WriteString(fmt.Sprintf("%v", sf.Month))
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(fmt.Sprintf("%v", sf.Date))
	builder.WriteString(", ")
	builder.WriteString("weekday=")
	builder.WriteString(fmt.Sprintf("%v", sf.Weekday))
	builder.WriteByte(')')
	return builder.String()
}

// SubjectFields is a parsable slice of SubjectField.
type SubjectFields []*SubjectField

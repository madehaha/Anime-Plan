// Code generated by ent, DO NOT EDIT.

package subjectfield

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subjectfield type in the database.
	Label = "subject_field"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRate1 holds the string denoting the rate_1 field in the database.
	FieldRate1 = "rate_1"
	// FieldRate2 holds the string denoting the rate_2 field in the database.
	FieldRate2 = "rate_2"
	// FieldRate3 holds the string denoting the rate_3 field in the database.
	FieldRate3 = "rate_3"
	// FieldRate4 holds the string denoting the rate_4 field in the database.
	FieldRate4 = "rate_4"
	// FieldRate5 holds the string denoting the rate_5 field in the database.
	FieldRate5 = "rate_5"
	// FieldRate6 holds the string denoting the rate_6 field in the database.
	FieldRate6 = "rate_6"
	// FieldRate7 holds the string denoting the rate_7 field in the database.
	FieldRate7 = "rate_7"
	// FieldRate8 holds the string denoting the rate_8 field in the database.
	FieldRate8 = "rate_8"
	// FieldRate9 holds the string denoting the rate_9 field in the database.
	FieldRate9 = "rate_9"
	// FieldRate10 holds the string denoting the rate_10 field in the database.
	FieldRate10 = "rate_10"
	// FieldAverageScore holds the string denoting the average_score field in the database.
	FieldAverageScore = "average_score"
	// FieldRank holds the string denoting the rank field in the database.
	FieldRank = "rank"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldMonth holds the string denoting the month field in the database.
	FieldMonth = "month"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldWeekday holds the string denoting the weekday field in the database.
	FieldWeekday = "weekday"
	// EdgeSubject holds the string denoting the subject edge name in mutations.
	EdgeSubject = "subject"
	// Table holds the table name of the subjectfield in the database.
	Table = "subject_fields"
	// SubjectTable is the table that holds the subject relation/edge.
	SubjectTable = "subject_fields"
	// SubjectInverseTable is the table name for the Subject entity.
	// It exists in this package in order to avoid circular dependency with the "subject" package.
	SubjectInverseTable = "subjects"
	// SubjectColumn is the table column denoting the subject relation/edge.
	SubjectColumn = "subject_subject_field"
)

// Columns holds all SQL columns for subjectfield fields.
var Columns = []string{
	FieldID,
	FieldRate1,
	FieldRate2,
	FieldRate3,
	FieldRate4,
	FieldRate5,
	FieldRate6,
	FieldRate7,
	FieldRate8,
	FieldRate9,
	FieldRate10,
	FieldAverageScore,
	FieldRank,
	FieldYear,
	FieldMonth,
	FieldDate,
	FieldWeekday,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "subject_fields"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"subject_subject_field",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRate1 holds the default value on creation for the "rate_1" field.
	DefaultRate1 uint32
	// DefaultRate2 holds the default value on creation for the "rate_2" field.
	DefaultRate2 uint32
	// DefaultRate3 holds the default value on creation for the "rate_3" field.
	DefaultRate3 uint32
	// DefaultRate4 holds the default value on creation for the "rate_4" field.
	DefaultRate4 uint32
	// DefaultRate5 holds the default value on creation for the "rate_5" field.
	DefaultRate5 uint32
	// DefaultRate6 holds the default value on creation for the "rate_6" field.
	DefaultRate6 uint32
	// DefaultRate7 holds the default value on creation for the "rate_7" field.
	DefaultRate7 uint32
	// DefaultRate8 holds the default value on creation for the "rate_8" field.
	DefaultRate8 uint32
	// DefaultRate9 holds the default value on creation for the "rate_9" field.
	DefaultRate9 uint32
	// DefaultRate10 holds the default value on creation for the "rate_10" field.
	DefaultRate10 uint32
	// DefaultAverageScore holds the default value on creation for the "average_score" field.
	DefaultAverageScore float64
	// AverageScoreValidator is a validator for the "average_score" field. It is called by the builders before save.
	AverageScoreValidator func(float64) error
	// DefaultRank holds the default value on creation for the "rank" field.
	DefaultRank uint32
	// MonthValidator is a validator for the "month" field. It is called by the builders before save.
	MonthValidator func(uint8) error
	// DateValidator is a validator for the "date" field. It is called by the builders before save.
	DateValidator func(uint8) error
	// WeekdayValidator is a validator for the "weekday" field. It is called by the builders before save.
	WeekdayValidator func(uint8) error
)

// OrderOption defines the ordering options for the SubjectField queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRate1 orders the results by the rate_1 field.
func ByRate1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate1, opts...).ToFunc()
}

// ByRate2 orders the results by the rate_2 field.
func ByRate2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate2, opts...).ToFunc()
}

// ByRate3 orders the results by the rate_3 field.
func ByRate3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate3, opts...).ToFunc()
}

// ByRate4 orders the results by the rate_4 field.
func ByRate4(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate4, opts...).ToFunc()
}

// ByRate5 orders the results by the rate_5 field.
func ByRate5(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate5, opts...).ToFunc()
}

// ByRate6 orders the results by the rate_6 field.
func ByRate6(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate6, opts...).ToFunc()
}

// ByRate7 orders the results by the rate_7 field.
func ByRate7(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate7, opts...).ToFunc()
}

// ByRate8 orders the results by the rate_8 field.
func ByRate8(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate8, opts...).ToFunc()
}

// ByRate9 orders the results by the rate_9 field.
func ByRate9(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate9, opts...).ToFunc()
}

// ByRate10 orders the results by the rate_10 field.
func ByRate10(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate10, opts...).ToFunc()
}

// ByAverageScore orders the results by the average_score field.
func ByAverageScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAverageScore, opts...).ToFunc()
}

// ByRank orders the results by the rank field.
func ByRank(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRank, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByMonth orders the results by the month field.
func ByMonth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMonth, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByWeekday orders the results by the weekday field.
func ByWeekday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeekday, opts...).ToFunc()
}

// BySubjectField orders the results by subject field.
func BySubjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubjectStep(), sql.OrderByField(field, opts...))
	}
}
func newSubjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, SubjectTable, SubjectColumn),
	)
}
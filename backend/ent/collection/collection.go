// Code generated by ent, DO NOT EDIT.

package collection

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the collection type in the database.
	Label = "collection"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldHasComment holds the string denoting the has_comment field in the database.
	FieldHasComment = "has_comment"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldAddTime holds the string denoting the add_time field in the database.
	FieldAddTime = "add_time"
	// FieldEpStatus holds the string denoting the ep_status field in the database.
	FieldEpStatus = "ep_status"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldSubjectID holds the string denoting the subject_id field in the database.
	FieldSubjectID = "subject_id"
	// EdgeMember holds the string denoting the member edge name in mutations.
	EdgeMember = "member"
	// EdgeSubject holds the string denoting the subject edge name in mutations.
	EdgeSubject = "subject"
	// MembersFieldID holds the string denoting the ID field of the Members.
	MembersFieldID = "uid"
	// Table holds the table name of the collection in the database.
	Table = "collections"
	// MemberTable is the table that holds the member relation/edge.
	MemberTable = "collections"
	// MemberInverseTable is the table name for the Members entity.
	// It exists in this package in order to avoid circular dependency with the "members" package.
	MemberInverseTable = "members"
	// MemberColumn is the table column denoting the member relation/edge.
	MemberColumn = "member_id"
	// SubjectTable is the table that holds the subject relation/edge.
	SubjectTable = "collections"
	// SubjectInverseTable is the table name for the Subject entity.
	// It exists in this package in order to avoid circular dependency with the "subject" package.
	SubjectInverseTable = "subjects"
	// SubjectColumn is the table column denoting the subject relation/edge.
	SubjectColumn = "subject_id"
)

// Columns holds all SQL columns for collection fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldHasComment,
	FieldComment,
	FieldScore,
	FieldAddTime,
	FieldEpStatus,
	FieldMemberID,
	FieldSubjectID,
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
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(uint8) error
	// DefaultHasComment holds the default value on creation for the "has_comment" field.
	DefaultHasComment bool
	// DefaultComment holds the default value on creation for the "comment" field.
	DefaultComment string
	// CommentValidator is a validator for the "comment" field. It is called by the builders before save.
	CommentValidator func(string) error
	// DefaultScore holds the default value on creation for the "score" field.
	DefaultScore uint8
	// DefaultAddTime holds the default value on creation for the "add_time" field.
	DefaultAddTime string
	// DefaultEpStatus holds the default value on creation for the "ep_status" field.
	DefaultEpStatus uint8
)

// OrderOption defines the ordering options for the Collection queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByHasComment orders the results by the has_comment field.
func ByHasComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasComment, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByAddTime orders the results by the add_time field.
func ByAddTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddTime, opts...).ToFunc()
}

// ByEpStatus orders the results by the ep_status field.
func ByEpStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEpStatus, opts...).ToFunc()
}

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// BySubjectID orders the results by the subject_id field.
func BySubjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubjectID, opts...).ToFunc()
}

// ByMemberField orders the results by member field.
func ByMemberField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMemberStep(), sql.OrderByField(field, opts...))
	}
}

// BySubjectField orders the results by subject field.
func BySubjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubjectStep(), sql.OrderByField(field, opts...))
	}
}
func newMemberStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MemberInverseTable, MembersFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
	)
}
func newSubjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
	)
}

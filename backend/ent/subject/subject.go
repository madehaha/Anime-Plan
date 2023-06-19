// Code generated by ent, DO NOT EDIT.

package subject

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subject type in the database.
	Label = "subject"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNameCn holds the string denoting the name_cn field in the database.
	FieldNameCn = "name_cn"
	// FieldEpisodes holds the string denoting the episodes field in the database.
	FieldEpisodes = "episodes"
	// FieldWish holds the string denoting the wish field in the database.
	FieldWish = "wish"
	// FieldDoing holds the string denoting the doing field in the database.
	FieldDoing = "doing"
	// FieldWatched holds the string denoting the watched field in the database.
	FieldWatched = "watched"
	// FieldOnHold holds the string denoting the on_hold field in the database.
	FieldOnHold = "on_hold"
	// FieldDropped holds the string denoting the dropped field in the database.
	FieldDropped = "dropped"
	// EdgeCollections holds the string denoting the collections edge name in mutations.
	EdgeCollections = "collections"
	// EdgeSubjectField holds the string denoting the subject_field edge name in mutations.
	EdgeSubjectField = "subject_field"
	// Table holds the table name of the subject in the database.
	Table = "subjects"
	// CollectionsTable is the table that holds the collections relation/edge.
	CollectionsTable = "collections"
	// CollectionsInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionsInverseTable = "collections"
	// CollectionsColumn is the table column denoting the collections relation/edge.
	CollectionsColumn = "subject_id"
	// SubjectFieldTable is the table that holds the subject_field relation/edge.
	SubjectFieldTable = "subject_fields"
	// SubjectFieldInverseTable is the table name for the SubjectField entity.
	// It exists in this package in order to avoid circular dependency with the "subjectfield" package.
	SubjectFieldInverseTable = "subject_fields"
	// SubjectFieldColumn is the table column denoting the subject_field relation/edge.
	SubjectFieldColumn = "subject_subject_field"
)

// Columns holds all SQL columns for subject fields.
var Columns = []string{
	FieldID,
	FieldImage,
	FieldSummary,
	FieldName,
	FieldNameCn,
	FieldEpisodes,
	FieldWish,
	FieldDoing,
	FieldWatched,
	FieldOnHold,
	FieldDropped,
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
	// DefaultImage holds the default value on creation for the "image" field.
	DefaultImage string
	// ImageValidator is a validator for the "image" field. It is called by the builders before save.
	ImageValidator func(string) error
	// DefaultSummary holds the default value on creation for the "summary" field.
	DefaultSummary string
	// SummaryValidator is a validator for the "summary" field. It is called by the builders before save.
	SummaryValidator func(string) error
	// DefaultWish holds the default value on creation for the "wish" field.
	DefaultWish uint32
	// DefaultDoing holds the default value on creation for the "doing" field.
	DefaultDoing uint32
	// DefaultWatched holds the default value on creation for the "watched" field.
	DefaultWatched uint32
	// DefaultOnHold holds the default value on creation for the "on_hold" field.
	DefaultOnHold uint32
	// DefaultDropped holds the default value on creation for the "dropped" field.
	DefaultDropped uint32
)

// OrderOption defines the ordering options for the Subject queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// BySummary orders the results by the summary field.
func BySummary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSummary, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByNameCn orders the results by the name_cn field.
func ByNameCn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNameCn, opts...).ToFunc()
}

// ByEpisodes orders the results by the episodes field.
func ByEpisodes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEpisodes, opts...).ToFunc()
}

// ByWish orders the results by the wish field.
func ByWish(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWish, opts...).ToFunc()
}

// ByDoing orders the results by the doing field.
func ByDoing(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDoing, opts...).ToFunc()
}

// ByWatched orders the results by the watched field.
func ByWatched(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWatched, opts...).ToFunc()
}

// ByOnHold orders the results by the on_hold field.
func ByOnHold(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOnHold, opts...).ToFunc()
}

// ByDropped orders the results by the dropped field.
func ByDropped(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDropped, opts...).ToFunc()
}

// ByCollectionsCount orders the results by collections count.
func ByCollectionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCollectionsStep(), opts...)
	}
}

// ByCollections orders the results by collections terms.
func ByCollections(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCollectionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubjectFieldField orders the results by subject_field field.
func BySubjectFieldField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubjectFieldStep(), sql.OrderByField(field, opts...))
	}
}
func newCollectionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CollectionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CollectionsTable, CollectionsColumn),
	)
}
func newSubjectFieldStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubjectFieldInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SubjectFieldTable, SubjectFieldColumn),
	)
}

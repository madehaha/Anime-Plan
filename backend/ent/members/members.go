// Code generated by ent, DO NOT EDIT.

package members

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the members type in the database.
	Label = "members"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldGid holds the string denoting the gid field in the database.
	FieldGid = "gid"
	// FieldRegisterTime holds the string denoting the register_time field in the database.
	FieldRegisterTime = "register_time"
	// EdgeCollections holds the string denoting the collections edge name in mutations.
	EdgeCollections = "collections"
	// CollectionFieldID holds the string denoting the ID field of the Collection.
	CollectionFieldID = "id"
	// Table holds the table name of the members in the database.
	Table = "members"
	// CollectionsTable is the table that holds the collections relation/edge.
	CollectionsTable = "collections"
	// CollectionsInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionsInverseTable = "collections"
	// CollectionsColumn is the table column denoting the collections relation/edge.
	CollectionsColumn = "members_collections"
)

// Columns holds all SQL columns for members fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEmail,
	FieldPassword,
	FieldNickname,
	FieldAvatar,
	FieldGid,
	FieldRegisterTime,
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// DefaultAvatar holds the default value on creation for the "avatar" field.
	DefaultAvatar string
	// AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	AvatarValidator func(string) error
	// DefaultGid holds the default value on creation for the "gid" field.
	DefaultGid uint8
)

// OrderOption defines the ordering options for the Members queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByGid orders the results by the gid field.
func ByGid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGid, opts...).ToFunc()
}

// ByRegisterTime orders the results by the register_time field.
func ByRegisterTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegisterTime, opts...).ToFunc()
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
func newCollectionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CollectionsInverseTable, CollectionFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CollectionsTable, CollectionsColumn),
	)
}

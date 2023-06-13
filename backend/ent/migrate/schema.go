// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MembersColumns holds the columns for the "members" table.
	MembersColumns = []*schema.Column{
		{Name: "uid", Type: field.TypeUint32, Increment: true},
		{Name: "username", Type: field.TypeString, Size: 30, Default: ""},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "password", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString, Size: 30},
		{Name: "avatar", Type: field.TypeString, Size: 255, Default: "https://lain.bgm.tv/pic/user/l/icon.jpg"},
		{Name: "gid", Type: field.TypeUint8, Default: 0},
		{Name: "register_time", Type: field.TypeString},
	}
	// MembersTable holds the schema information for the "members" table.
	MembersTable = &schema.Table{
		Name:       "members",
		Columns:    MembersColumns,
		PrimaryKey: []*schema.Column{MembersColumns[0]},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "uid", Type: field.TypeUint32, Increment: true},
		{Name: "image", Type: field.TypeString, Default: "https://lain.bgm.tv/pic/user/l/icon.jpg"},
		{Name: "summary", Type: field.TypeString, Size: 300},
		{Name: "name", Type: field.TypeString},
		{Name: "date", Type: field.TypeString},
		{Name: "name_cn", Type: field.TypeString},
		{Name: "on_hold", Type: field.TypeUint32},
		{Name: "wish", Type: field.TypeUint32},
		{Name: "doing", Type: field.TypeUint32},
		{Name: "subject_type", Type: field.TypeUint8, Default: 0},
		{Name: "collect", Type: field.TypeUint32, Default: 0},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MembersTable,
		SubjectsTable,
	}
)

func init() {
	SubjectsTable.Annotation = &entsql.Annotation{
		Table: "subjects",
	}
}

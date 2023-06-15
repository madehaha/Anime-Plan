// Code generated by ent, DO NOT EDIT.

package subject

import (
	"backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldID, id))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldImage, v))
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldSummary, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldName, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldDate, v))
}

// NameCn applies equality check predicate on the "name_cn" field. It's identical to NameCnEQ.
func NameCn(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldNameCn, v))
}

// OnHold applies equality check predicate on the "on_hold" field. It's identical to OnHoldEQ.
func OnHold(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldOnHold, v))
}

// Wish applies equality check predicate on the "wish" field. It's identical to WishEQ.
func Wish(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldWish, v))
}

// Doing applies equality check predicate on the "doing" field. It's identical to DoingEQ.
func Doing(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldDoing, v))
}

// SubjectType applies equality check predicate on the "subject_type" field. It's identical to SubjectTypeEQ.
func SubjectType(v uint8) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldSubjectType, v))
}

// Collect applies equality check predicate on the "collect" field. It's identical to CollectEQ.
func Collect(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldCollect, v))
}

// Drop applies equality check predicate on the "drop" field. It's identical to DropEQ.
func Drop(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldDrop, v))
}

// Watched applies equality check predicate on the "watched" field. It's identical to WatchedEQ.
func Watched(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldWatched, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasSuffix(FieldImage, v))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContainsFold(FieldImage, v))
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldSummary, v))
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldSummary, v))
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldSummary, vs...))
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldSummary, vs...))
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldSummary, v))
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldSummary, v))
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldSummary, v))
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldSummary, v))
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContains(FieldSummary, v))
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasPrefix(FieldSummary, v))
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasSuffix(FieldSummary, v))
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEqualFold(FieldSummary, v))
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContainsFold(FieldSummary, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContainsFold(FieldName, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldDate, v))
}

// DateContains applies the Contains predicate on the "date" field.
func DateContains(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContains(FieldDate, v))
}

// DateHasPrefix applies the HasPrefix predicate on the "date" field.
func DateHasPrefix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasPrefix(FieldDate, v))
}

// DateHasSuffix applies the HasSuffix predicate on the "date" field.
func DateHasSuffix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasSuffix(FieldDate, v))
}

// DateEqualFold applies the EqualFold predicate on the "date" field.
func DateEqualFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEqualFold(FieldDate, v))
}

// DateContainsFold applies the ContainsFold predicate on the "date" field.
func DateContainsFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContainsFold(FieldDate, v))
}

// NameCnEQ applies the EQ predicate on the "name_cn" field.
func NameCnEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldNameCn, v))
}

// NameCnNEQ applies the NEQ predicate on the "name_cn" field.
func NameCnNEQ(v string) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldNameCn, v))
}

// NameCnIn applies the In predicate on the "name_cn" field.
func NameCnIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldNameCn, vs...))
}

// NameCnNotIn applies the NotIn predicate on the "name_cn" field.
func NameCnNotIn(vs ...string) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldNameCn, vs...))
}

// NameCnGT applies the GT predicate on the "name_cn" field.
func NameCnGT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldNameCn, v))
}

// NameCnGTE applies the GTE predicate on the "name_cn" field.
func NameCnGTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldNameCn, v))
}

// NameCnLT applies the LT predicate on the "name_cn" field.
func NameCnLT(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldNameCn, v))
}

// NameCnLTE applies the LTE predicate on the "name_cn" field.
func NameCnLTE(v string) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldNameCn, v))
}

// NameCnContains applies the Contains predicate on the "name_cn" field.
func NameCnContains(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContains(FieldNameCn, v))
}

// NameCnHasPrefix applies the HasPrefix predicate on the "name_cn" field.
func NameCnHasPrefix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasPrefix(FieldNameCn, v))
}

// NameCnHasSuffix applies the HasSuffix predicate on the "name_cn" field.
func NameCnHasSuffix(v string) predicate.Subject {
	return predicate.Subject(sql.FieldHasSuffix(FieldNameCn, v))
}

// NameCnEqualFold applies the EqualFold predicate on the "name_cn" field.
func NameCnEqualFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldEqualFold(FieldNameCn, v))
}

// NameCnContainsFold applies the ContainsFold predicate on the "name_cn" field.
func NameCnContainsFold(v string) predicate.Subject {
	return predicate.Subject(sql.FieldContainsFold(FieldNameCn, v))
}

// OnHoldEQ applies the EQ predicate on the "on_hold" field.
func OnHoldEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldOnHold, v))
}

// OnHoldNEQ applies the NEQ predicate on the "on_hold" field.
func OnHoldNEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldOnHold, v))
}

// OnHoldIn applies the In predicate on the "on_hold" field.
func OnHoldIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldOnHold, vs...))
}

// OnHoldNotIn applies the NotIn predicate on the "on_hold" field.
func OnHoldNotIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldOnHold, vs...))
}

// OnHoldGT applies the GT predicate on the "on_hold" field.
func OnHoldGT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldOnHold, v))
}

// OnHoldGTE applies the GTE predicate on the "on_hold" field.
func OnHoldGTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldOnHold, v))
}

// OnHoldLT applies the LT predicate on the "on_hold" field.
func OnHoldLT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldOnHold, v))
}

// OnHoldLTE applies the LTE predicate on the "on_hold" field.
func OnHoldLTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldOnHold, v))
}

// WishEQ applies the EQ predicate on the "wish" field.
func WishEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldWish, v))
}

// WishNEQ applies the NEQ predicate on the "wish" field.
func WishNEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldWish, v))
}

// WishIn applies the In predicate on the "wish" field.
func WishIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldWish, vs...))
}

// WishNotIn applies the NotIn predicate on the "wish" field.
func WishNotIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldWish, vs...))
}

// WishGT applies the GT predicate on the "wish" field.
func WishGT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldWish, v))
}

// WishGTE applies the GTE predicate on the "wish" field.
func WishGTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldWish, v))
}

// WishLT applies the LT predicate on the "wish" field.
func WishLT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldWish, v))
}

// WishLTE applies the LTE predicate on the "wish" field.
func WishLTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldWish, v))
}

// DoingEQ applies the EQ predicate on the "doing" field.
func DoingEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldDoing, v))
}

// DoingNEQ applies the NEQ predicate on the "doing" field.
func DoingNEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldDoing, v))
}

// DoingIn applies the In predicate on the "doing" field.
func DoingIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldDoing, vs...))
}

// DoingNotIn applies the NotIn predicate on the "doing" field.
func DoingNotIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldDoing, vs...))
}

// DoingGT applies the GT predicate on the "doing" field.
func DoingGT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldDoing, v))
}

// DoingGTE applies the GTE predicate on the "doing" field.
func DoingGTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldDoing, v))
}

// DoingLT applies the LT predicate on the "doing" field.
func DoingLT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldDoing, v))
}

// DoingLTE applies the LTE predicate on the "doing" field.
func DoingLTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldDoing, v))
}

// SubjectTypeEQ applies the EQ predicate on the "subject_type" field.
func SubjectTypeEQ(v uint8) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldSubjectType, v))
}

// SubjectTypeNEQ applies the NEQ predicate on the "subject_type" field.
func SubjectTypeNEQ(v uint8) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldSubjectType, v))
}

// SubjectTypeIn applies the In predicate on the "subject_type" field.
func SubjectTypeIn(vs ...uint8) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldSubjectType, vs...))
}

// SubjectTypeNotIn applies the NotIn predicate on the "subject_type" field.
func SubjectTypeNotIn(vs ...uint8) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldSubjectType, vs...))
}

// SubjectTypeGT applies the GT predicate on the "subject_type" field.
func SubjectTypeGT(v uint8) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldSubjectType, v))
}

// SubjectTypeGTE applies the GTE predicate on the "subject_type" field.
func SubjectTypeGTE(v uint8) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldSubjectType, v))
}

// SubjectTypeLT applies the LT predicate on the "subject_type" field.
func SubjectTypeLT(v uint8) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldSubjectType, v))
}

// SubjectTypeLTE applies the LTE predicate on the "subject_type" field.
func SubjectTypeLTE(v uint8) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldSubjectType, v))
}

// CollectEQ applies the EQ predicate on the "collect" field.
func CollectEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldCollect, v))
}

// CollectNEQ applies the NEQ predicate on the "collect" field.
func CollectNEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldCollect, v))
}

// CollectIn applies the In predicate on the "collect" field.
func CollectIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldCollect, vs...))
}

// CollectNotIn applies the NotIn predicate on the "collect" field.
func CollectNotIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldCollect, vs...))
}

// CollectGT applies the GT predicate on the "collect" field.
func CollectGT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldCollect, v))
}

// CollectGTE applies the GTE predicate on the "collect" field.
func CollectGTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldCollect, v))
}

// CollectLT applies the LT predicate on the "collect" field.
func CollectLT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldCollect, v))
}

// CollectLTE applies the LTE predicate on the "collect" field.
func CollectLTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldCollect, v))
}

// DropEQ applies the EQ predicate on the "drop" field.
func DropEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldDrop, v))
}

// DropNEQ applies the NEQ predicate on the "drop" field.
func DropNEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldDrop, v))
}

// DropIn applies the In predicate on the "drop" field.
func DropIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldDrop, vs...))
}

// DropNotIn applies the NotIn predicate on the "drop" field.
func DropNotIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldDrop, vs...))
}

// DropGT applies the GT predicate on the "drop" field.
func DropGT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldDrop, v))
}

// DropGTE applies the GTE predicate on the "drop" field.
func DropGTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldDrop, v))
}

// DropLT applies the LT predicate on the "drop" field.
func DropLT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldDrop, v))
}

// DropLTE applies the LTE predicate on the "drop" field.
func DropLTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldDrop, v))
}

// WatchedEQ applies the EQ predicate on the "watched" field.
func WatchedEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldEQ(FieldWatched, v))
}

// WatchedNEQ applies the NEQ predicate on the "watched" field.
func WatchedNEQ(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNEQ(FieldWatched, v))
}

// WatchedIn applies the In predicate on the "watched" field.
func WatchedIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldIn(FieldWatched, vs...))
}

// WatchedNotIn applies the NotIn predicate on the "watched" field.
func WatchedNotIn(vs ...uint32) predicate.Subject {
	return predicate.Subject(sql.FieldNotIn(FieldWatched, vs...))
}

// WatchedGT applies the GT predicate on the "watched" field.
func WatchedGT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGT(FieldWatched, v))
}

// WatchedGTE applies the GTE predicate on the "watched" field.
func WatchedGTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldGTE(FieldWatched, v))
}

// WatchedLT applies the LT predicate on the "watched" field.
func WatchedLT(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLT(FieldWatched, v))
}

// WatchedLTE applies the LTE predicate on the "watched" field.
func WatchedLTE(v uint32) predicate.Subject {
	return predicate.Subject(sql.FieldLTE(FieldWatched, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Subject) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Subject) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Subject) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		p(s.Not())
	})
}

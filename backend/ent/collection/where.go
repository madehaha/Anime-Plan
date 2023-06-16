// Code generated by ent, DO NOT EDIT.

package collection

import (
	"backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldType, v))
}

// HasComment applies equality check predicate on the "has_comment" field. It's identical to HasCommentEQ.
func HasComment(v bool) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldHasComment, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldComment, v))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldScore, v))
}

// AddTime applies equality check predicate on the "add_time" field. It's identical to AddTimeEQ.
func AddTime(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldAddTime, v))
}

// EpStatus applies equality check predicate on the "ep_status" field. It's identical to EpStatusEQ.
func EpStatus(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldEpStatus, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...uint8) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...uint8) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldType, v))
}

// HasCommentEQ applies the EQ predicate on the "has_comment" field.
func HasCommentEQ(v bool) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldHasComment, v))
}

// HasCommentNEQ applies the NEQ predicate on the "has_comment" field.
func HasCommentNEQ(v bool) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldHasComment, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldComment, v))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldComment, v))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...uint8) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...uint8) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldScore, v))
}

// AddTimeEQ applies the EQ predicate on the "add_time" field.
func AddTimeEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldAddTime, v))
}

// AddTimeNEQ applies the NEQ predicate on the "add_time" field.
func AddTimeNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldAddTime, v))
}

// AddTimeIn applies the In predicate on the "add_time" field.
func AddTimeIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldAddTime, vs...))
}

// AddTimeNotIn applies the NotIn predicate on the "add_time" field.
func AddTimeNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldAddTime, vs...))
}

// AddTimeGT applies the GT predicate on the "add_time" field.
func AddTimeGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldAddTime, v))
}

// AddTimeGTE applies the GTE predicate on the "add_time" field.
func AddTimeGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldAddTime, v))
}

// AddTimeLT applies the LT predicate on the "add_time" field.
func AddTimeLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldAddTime, v))
}

// AddTimeLTE applies the LTE predicate on the "add_time" field.
func AddTimeLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldAddTime, v))
}

// AddTimeContains applies the Contains predicate on the "add_time" field.
func AddTimeContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldAddTime, v))
}

// AddTimeHasPrefix applies the HasPrefix predicate on the "add_time" field.
func AddTimeHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldAddTime, v))
}

// AddTimeHasSuffix applies the HasSuffix predicate on the "add_time" field.
func AddTimeHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldAddTime, v))
}

// AddTimeEqualFold applies the EqualFold predicate on the "add_time" field.
func AddTimeEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldAddTime, v))
}

// AddTimeContainsFold applies the ContainsFold predicate on the "add_time" field.
func AddTimeContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldAddTime, v))
}

// EpStatusEQ applies the EQ predicate on the "ep_status" field.
func EpStatusEQ(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldEpStatus, v))
}

// EpStatusNEQ applies the NEQ predicate on the "ep_status" field.
func EpStatusNEQ(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldEpStatus, v))
}

// EpStatusIn applies the In predicate on the "ep_status" field.
func EpStatusIn(vs ...uint8) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldEpStatus, vs...))
}

// EpStatusNotIn applies the NotIn predicate on the "ep_status" field.
func EpStatusNotIn(vs ...uint8) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldEpStatus, vs...))
}

// EpStatusGT applies the GT predicate on the "ep_status" field.
func EpStatusGT(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldEpStatus, v))
}

// EpStatusGTE applies the GTE predicate on the "ep_status" field.
func EpStatusGTE(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldEpStatus, v))
}

// EpStatusLT applies the LT predicate on the "ep_status" field.
func EpStatusLT(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldEpStatus, v))
}

// EpStatusLTE applies the LTE predicate on the "ep_status" field.
func EpStatusLTE(v uint8) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldEpStatus, v))
}

// HasMember applies the HasEdge predicate on the "member" edge.
func HasMember() predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberWith applies the HasEdge predicate on the "member" edge with a given conditions (other predicates).
func HasMemberWith(preds ...predicate.Members) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := newMemberStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubject applies the HasEdge predicate on the "subject" edge.
func HasSubject() predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectWith applies the HasEdge predicate on the "subject" edge with a given conditions (other predicates).
func HasSubjectWith(preds ...predicate.Subject) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := newSubjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Collection) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Collection) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
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
func Not(p predicate.Collection) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		p(s.Not())
	})
}

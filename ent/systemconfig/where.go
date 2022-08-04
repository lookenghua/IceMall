// Code generated by entc, DO NOT EDIT.

package systemconfig

import (
	"ice-mall/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Rule applies equality check predicate on the "rule" field. It's identical to RuleEQ.
func Rule(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRule), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// RuleEQ applies the EQ predicate on the "rule" field.
func RuleEQ(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRule), v))
	})
}

// RuleNEQ applies the NEQ predicate on the "rule" field.
func RuleNEQ(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRule), v))
	})
}

// RuleIn applies the In predicate on the "rule" field.
func RuleIn(vs ...string) predicate.SystemConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SystemConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRule), v...))
	})
}

// RuleNotIn applies the NotIn predicate on the "rule" field.
func RuleNotIn(vs ...string) predicate.SystemConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SystemConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRule), v...))
	})
}

// RuleGT applies the GT predicate on the "rule" field.
func RuleGT(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRule), v))
	})
}

// RuleGTE applies the GTE predicate on the "rule" field.
func RuleGTE(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRule), v))
	})
}

// RuleLT applies the LT predicate on the "rule" field.
func RuleLT(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRule), v))
	})
}

// RuleLTE applies the LTE predicate on the "rule" field.
func RuleLTE(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRule), v))
	})
}

// RuleContains applies the Contains predicate on the "rule" field.
func RuleContains(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRule), v))
	})
}

// RuleHasPrefix applies the HasPrefix predicate on the "rule" field.
func RuleHasPrefix(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRule), v))
	})
}

// RuleHasSuffix applies the HasSuffix predicate on the "rule" field.
func RuleHasSuffix(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRule), v))
	})
}

// RuleEqualFold applies the EqualFold predicate on the "rule" field.
func RuleEqualFold(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRule), v))
	})
}

// RuleContainsFold applies the ContainsFold predicate on the "rule" field.
func RuleContainsFold(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRule), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.SystemConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SystemConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.SystemConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SystemConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldValue), v))
	})
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldValue), v))
	})
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldValue), v))
	})
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldValue), v))
	})
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldValue), v))
	})
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemark), v))
	})
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.SystemConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SystemConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRemark), v...))
	})
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.SystemConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SystemConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRemark), v...))
	})
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemark), v))
	})
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemark), v))
	})
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemark), v))
	})
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemark), v))
	})
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemark), v))
	})
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemark), v))
	})
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemark), v))
	})
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemark), v))
	})
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemark), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SystemConfig) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SystemConfig) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
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
func Not(p predicate.SystemConfig) predicate.SystemConfig {
	return predicate.SystemConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}
// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package customtype

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv1/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
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
func IDGT(id int) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Custom applies equality check predicate on the "custom" field. It's identical to CustomEQ.
func Custom(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustom), v))
	})
}

// CustomEQ applies the EQ predicate on the "custom" field.
func CustomEQ(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustom), v))
	})
}

// CustomNEQ applies the NEQ predicate on the "custom" field.
func CustomNEQ(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCustom), v))
	})
}

// CustomIn applies the In predicate on the "custom" field.
func CustomIn(vs ...string) predicate.CustomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCustom), v...))
	})
}

// CustomNotIn applies the NotIn predicate on the "custom" field.
func CustomNotIn(vs ...string) predicate.CustomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCustom), v...))
	})
}

// CustomGT applies the GT predicate on the "custom" field.
func CustomGT(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCustom), v))
	})
}

// CustomGTE applies the GTE predicate on the "custom" field.
func CustomGTE(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCustom), v))
	})
}

// CustomLT applies the LT predicate on the "custom" field.
func CustomLT(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCustom), v))
	})
}

// CustomLTE applies the LTE predicate on the "custom" field.
func CustomLTE(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCustom), v))
	})
}

// CustomContains applies the Contains predicate on the "custom" field.
func CustomContains(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCustom), v))
	})
}

// CustomHasPrefix applies the HasPrefix predicate on the "custom" field.
func CustomHasPrefix(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCustom), v))
	})
}

// CustomHasSuffix applies the HasSuffix predicate on the "custom" field.
func CustomHasSuffix(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCustom), v))
	})
}

// CustomIsNil applies the IsNil predicate on the "custom" field.
func CustomIsNil() predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCustom)))
	})
}

// CustomNotNil applies the NotNil predicate on the "custom" field.
func CustomNotNil() predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCustom)))
	})
}

// CustomEqualFold applies the EqualFold predicate on the "custom" field.
func CustomEqualFold(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCustom), v))
	})
}

// CustomContainsFold applies the ContainsFold predicate on the "custom" field.
func CustomContainsFold(v string) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCustom), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CustomType) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CustomType) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
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
func Not(p predicate.CustomType) predicate.CustomType {
	return predicate.CustomType(func(s *sql.Selector) {
		p(s.Not())
	})
}
// Code generated by entc, DO NOT EDIT.

package goodsspecs

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/a20070322/shop-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
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
func IDGT(id int) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// SpecsName applies equality check predicate on the "specs_name" field. It's identical to SpecsNameEQ.
func SpecsName(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpecsName), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GoodsSpecs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GoodsSpecs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GoodsSpecs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GoodsSpecs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GoodsSpecs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.GoodsSpecs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// SpecsNameEQ applies the EQ predicate on the "specs_name" field.
func SpecsNameEQ(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpecsName), v))
	})
}

// SpecsNameNEQ applies the NEQ predicate on the "specs_name" field.
func SpecsNameNEQ(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpecsName), v))
	})
}

// SpecsNameIn applies the In predicate on the "specs_name" field.
func SpecsNameIn(vs ...string) predicate.GoodsSpecs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSpecsName), v...))
	})
}

// SpecsNameNotIn applies the NotIn predicate on the "specs_name" field.
func SpecsNameNotIn(vs ...string) predicate.GoodsSpecs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpecsName), v...))
	})
}

// SpecsNameGT applies the GT predicate on the "specs_name" field.
func SpecsNameGT(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpecsName), v))
	})
}

// SpecsNameGTE applies the GTE predicate on the "specs_name" field.
func SpecsNameGTE(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpecsName), v))
	})
}

// SpecsNameLT applies the LT predicate on the "specs_name" field.
func SpecsNameLT(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpecsName), v))
	})
}

// SpecsNameLTE applies the LTE predicate on the "specs_name" field.
func SpecsNameLTE(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpecsName), v))
	})
}

// SpecsNameContains applies the Contains predicate on the "specs_name" field.
func SpecsNameContains(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSpecsName), v))
	})
}

// SpecsNameHasPrefix applies the HasPrefix predicate on the "specs_name" field.
func SpecsNameHasPrefix(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSpecsName), v))
	})
}

// SpecsNameHasSuffix applies the HasSuffix predicate on the "specs_name" field.
func SpecsNameHasSuffix(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSpecsName), v))
	})
}

// SpecsNameEqualFold applies the EqualFold predicate on the "specs_name" field.
func SpecsNameEqualFold(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSpecsName), v))
	})
}

// SpecsNameContainsFold applies the ContainsFold predicate on the "specs_name" field.
func SpecsNameContainsFold(v string) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSpecsName), v))
	})
}

// HasGoodsSpecsOption applies the HasEdge predicate on the "goods_specs_option" edge.
func HasGoodsSpecsOption() predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSpecsOptionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GoodsSpecsOptionTable, GoodsSpecsOptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGoodsSpecsOptionWith applies the HasEdge predicate on the "goods_specs_option" edge with a given conditions (other predicates).
func HasGoodsSpecsOptionWith(preds ...predicate.GoodsSpecsOption) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSpecsOptionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GoodsSpecsOptionTable, GoodsSpecsOptionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodsSpecs) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodsSpecs) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
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
func Not(p predicate.GoodsSpecs) predicate.GoodsSpecs {
	return predicate.GoodsSpecs(func(s *sql.Selector) {
		p(s.Not())
	})
}

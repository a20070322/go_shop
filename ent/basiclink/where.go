// Code generated by entc, DO NOT EDIT.

package basiclink

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/a20070322/shop-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// LinkName applies equality check predicate on the "link_name" field. It's identical to LinkNameEQ.
func LinkName(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkName), v))
	})
}

// LinkType applies equality check predicate on the "link_type" field. It's identical to LinkTypeEQ.
func LinkType(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkType), v))
	})
}

// LinkAddress applies equality check predicate on the "link_address" field. It's identical to LinkAddressEQ.
func LinkAddress(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkAddress), v))
	})
}

// Appid applies equality check predicate on the "appid" field. It's identical to AppidEQ.
func Appid(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppid), v))
	})
}

// IsRegister applies equality check predicate on the "is_register" field. It's identical to IsRegisterEQ.
func IsRegister(v bool) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsRegister), v))
	})
}

// Remarks applies equality check predicate on the "remarks" field. It's identical to RemarksEQ.
func Remarks(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemarks), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// LinkNameEQ applies the EQ predicate on the "link_name" field.
func LinkNameEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkName), v))
	})
}

// LinkNameNEQ applies the NEQ predicate on the "link_name" field.
func LinkNameNEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLinkName), v))
	})
}

// LinkNameIn applies the In predicate on the "link_name" field.
func LinkNameIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLinkName), v...))
	})
}

// LinkNameNotIn applies the NotIn predicate on the "link_name" field.
func LinkNameNotIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLinkName), v...))
	})
}

// LinkNameGT applies the GT predicate on the "link_name" field.
func LinkNameGT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLinkName), v))
	})
}

// LinkNameGTE applies the GTE predicate on the "link_name" field.
func LinkNameGTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLinkName), v))
	})
}

// LinkNameLT applies the LT predicate on the "link_name" field.
func LinkNameLT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLinkName), v))
	})
}

// LinkNameLTE applies the LTE predicate on the "link_name" field.
func LinkNameLTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLinkName), v))
	})
}

// LinkNameContains applies the Contains predicate on the "link_name" field.
func LinkNameContains(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLinkName), v))
	})
}

// LinkNameHasPrefix applies the HasPrefix predicate on the "link_name" field.
func LinkNameHasPrefix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLinkName), v))
	})
}

// LinkNameHasSuffix applies the HasSuffix predicate on the "link_name" field.
func LinkNameHasSuffix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLinkName), v))
	})
}

// LinkNameEqualFold applies the EqualFold predicate on the "link_name" field.
func LinkNameEqualFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLinkName), v))
	})
}

// LinkNameContainsFold applies the ContainsFold predicate on the "link_name" field.
func LinkNameContainsFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLinkName), v))
	})
}

// LinkTypeEQ applies the EQ predicate on the "link_type" field.
func LinkTypeEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkType), v))
	})
}

// LinkTypeNEQ applies the NEQ predicate on the "link_type" field.
func LinkTypeNEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLinkType), v))
	})
}

// LinkTypeIn applies the In predicate on the "link_type" field.
func LinkTypeIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLinkType), v...))
	})
}

// LinkTypeNotIn applies the NotIn predicate on the "link_type" field.
func LinkTypeNotIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLinkType), v...))
	})
}

// LinkTypeGT applies the GT predicate on the "link_type" field.
func LinkTypeGT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLinkType), v))
	})
}

// LinkTypeGTE applies the GTE predicate on the "link_type" field.
func LinkTypeGTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLinkType), v))
	})
}

// LinkTypeLT applies the LT predicate on the "link_type" field.
func LinkTypeLT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLinkType), v))
	})
}

// LinkTypeLTE applies the LTE predicate on the "link_type" field.
func LinkTypeLTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLinkType), v))
	})
}

// LinkTypeContains applies the Contains predicate on the "link_type" field.
func LinkTypeContains(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLinkType), v))
	})
}

// LinkTypeHasPrefix applies the HasPrefix predicate on the "link_type" field.
func LinkTypeHasPrefix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLinkType), v))
	})
}

// LinkTypeHasSuffix applies the HasSuffix predicate on the "link_type" field.
func LinkTypeHasSuffix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLinkType), v))
	})
}

// LinkTypeEqualFold applies the EqualFold predicate on the "link_type" field.
func LinkTypeEqualFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLinkType), v))
	})
}

// LinkTypeContainsFold applies the ContainsFold predicate on the "link_type" field.
func LinkTypeContainsFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLinkType), v))
	})
}

// LinkAddressEQ applies the EQ predicate on the "link_address" field.
func LinkAddressEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressNEQ applies the NEQ predicate on the "link_address" field.
func LinkAddressNEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressIn applies the In predicate on the "link_address" field.
func LinkAddressIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLinkAddress), v...))
	})
}

// LinkAddressNotIn applies the NotIn predicate on the "link_address" field.
func LinkAddressNotIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLinkAddress), v...))
	})
}

// LinkAddressGT applies the GT predicate on the "link_address" field.
func LinkAddressGT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressGTE applies the GTE predicate on the "link_address" field.
func LinkAddressGTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressLT applies the LT predicate on the "link_address" field.
func LinkAddressLT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressLTE applies the LTE predicate on the "link_address" field.
func LinkAddressLTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressContains applies the Contains predicate on the "link_address" field.
func LinkAddressContains(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressHasPrefix applies the HasPrefix predicate on the "link_address" field.
func LinkAddressHasPrefix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressHasSuffix applies the HasSuffix predicate on the "link_address" field.
func LinkAddressHasSuffix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressEqualFold applies the EqualFold predicate on the "link_address" field.
func LinkAddressEqualFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLinkAddress), v))
	})
}

// LinkAddressContainsFold applies the ContainsFold predicate on the "link_address" field.
func LinkAddressContainsFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLinkAddress), v))
	})
}

// AppidEQ applies the EQ predicate on the "appid" field.
func AppidEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppid), v))
	})
}

// AppidNEQ applies the NEQ predicate on the "appid" field.
func AppidNEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppid), v))
	})
}

// AppidIn applies the In predicate on the "appid" field.
func AppidIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppid), v...))
	})
}

// AppidNotIn applies the NotIn predicate on the "appid" field.
func AppidNotIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppid), v...))
	})
}

// AppidGT applies the GT predicate on the "appid" field.
func AppidGT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppid), v))
	})
}

// AppidGTE applies the GTE predicate on the "appid" field.
func AppidGTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppid), v))
	})
}

// AppidLT applies the LT predicate on the "appid" field.
func AppidLT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppid), v))
	})
}

// AppidLTE applies the LTE predicate on the "appid" field.
func AppidLTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppid), v))
	})
}

// AppidContains applies the Contains predicate on the "appid" field.
func AppidContains(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppid), v))
	})
}

// AppidHasPrefix applies the HasPrefix predicate on the "appid" field.
func AppidHasPrefix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppid), v))
	})
}

// AppidHasSuffix applies the HasSuffix predicate on the "appid" field.
func AppidHasSuffix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppid), v))
	})
}

// AppidIsNil applies the IsNil predicate on the "appid" field.
func AppidIsNil() predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppid)))
	})
}

// AppidNotNil applies the NotNil predicate on the "appid" field.
func AppidNotNil() predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppid)))
	})
}

// AppidEqualFold applies the EqualFold predicate on the "appid" field.
func AppidEqualFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppid), v))
	})
}

// AppidContainsFold applies the ContainsFold predicate on the "appid" field.
func AppidContainsFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppid), v))
	})
}

// IsRegisterEQ applies the EQ predicate on the "is_register" field.
func IsRegisterEQ(v bool) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsRegister), v))
	})
}

// IsRegisterNEQ applies the NEQ predicate on the "is_register" field.
func IsRegisterNEQ(v bool) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsRegister), v))
	})
}

// RemarksEQ applies the EQ predicate on the "remarks" field.
func RemarksEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemarks), v))
	})
}

// RemarksNEQ applies the NEQ predicate on the "remarks" field.
func RemarksNEQ(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemarks), v))
	})
}

// RemarksIn applies the In predicate on the "remarks" field.
func RemarksIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRemarks), v...))
	})
}

// RemarksNotIn applies the NotIn predicate on the "remarks" field.
func RemarksNotIn(vs ...string) predicate.BasicLink {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BasicLink(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRemarks), v...))
	})
}

// RemarksGT applies the GT predicate on the "remarks" field.
func RemarksGT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemarks), v))
	})
}

// RemarksGTE applies the GTE predicate on the "remarks" field.
func RemarksGTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemarks), v))
	})
}

// RemarksLT applies the LT predicate on the "remarks" field.
func RemarksLT(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemarks), v))
	})
}

// RemarksLTE applies the LTE predicate on the "remarks" field.
func RemarksLTE(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemarks), v))
	})
}

// RemarksContains applies the Contains predicate on the "remarks" field.
func RemarksContains(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemarks), v))
	})
}

// RemarksHasPrefix applies the HasPrefix predicate on the "remarks" field.
func RemarksHasPrefix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemarks), v))
	})
}

// RemarksHasSuffix applies the HasSuffix predicate on the "remarks" field.
func RemarksHasSuffix(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemarks), v))
	})
}

// RemarksIsNil applies the IsNil predicate on the "remarks" field.
func RemarksIsNil() predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRemarks)))
	})
}

// RemarksNotNil applies the NotNil predicate on the "remarks" field.
func RemarksNotNil() predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRemarks)))
	})
}

// RemarksEqualFold applies the EqualFold predicate on the "remarks" field.
func RemarksEqualFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemarks), v))
	})
}

// RemarksContainsFold applies the ContainsFold predicate on the "remarks" field.
func RemarksContainsFold(v string) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemarks), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// HasBasicBanner applies the HasEdge predicate on the "basic_banner" edge.
func HasBasicBanner() predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BasicBannerTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BasicBannerTable, BasicBannerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBasicBannerWith applies the HasEdge predicate on the "basic_banner" edge with a given conditions (other predicates).
func HasBasicBannerWith(preds ...predicate.BasicBanner) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BasicBannerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BasicBannerTable, BasicBannerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BasicLink) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BasicLink) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
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
func Not(p predicate.BasicLink) predicate.BasicLink {
	return predicate.BasicLink(func(s *sql.Selector) {
		p(s.Not())
	})
}

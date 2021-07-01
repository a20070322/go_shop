// Code generated by entc, DO NOT EDIT.

package orderinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/a20070322/shop-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// OrderNumber applies equality check predicate on the "order_number" field. It's identical to OrderNumberEQ.
func OrderNumber(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderNumber), v))
	})
}

// PrepayID applies equality check predicate on the "prepay_id" field. It's identical to PrepayIDEQ.
func PrepayID(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrepayID), v))
	})
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// DeliveryStatus applies equality check predicate on the "delivery_status" field. It's identical to DeliveryStatusEQ.
func DeliveryStatus(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeliveryStatus), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// OrderNumberEQ applies the EQ predicate on the "order_number" field.
func OrderNumberEQ(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberNEQ applies the NEQ predicate on the "order_number" field.
func OrderNumberNEQ(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberIn applies the In predicate on the "order_number" field.
func OrderNumberIn(vs ...string) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrderNumber), v...))
	})
}

// OrderNumberNotIn applies the NotIn predicate on the "order_number" field.
func OrderNumberNotIn(vs ...string) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrderNumber), v...))
	})
}

// OrderNumberGT applies the GT predicate on the "order_number" field.
func OrderNumberGT(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberGTE applies the GTE predicate on the "order_number" field.
func OrderNumberGTE(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberLT applies the LT predicate on the "order_number" field.
func OrderNumberLT(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberLTE applies the LTE predicate on the "order_number" field.
func OrderNumberLTE(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberContains applies the Contains predicate on the "order_number" field.
func OrderNumberContains(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberHasPrefix applies the HasPrefix predicate on the "order_number" field.
func OrderNumberHasPrefix(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberHasSuffix applies the HasSuffix predicate on the "order_number" field.
func OrderNumberHasSuffix(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberEqualFold applies the EqualFold predicate on the "order_number" field.
func OrderNumberEqualFold(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrderNumber), v))
	})
}

// OrderNumberContainsFold applies the ContainsFold predicate on the "order_number" field.
func OrderNumberContainsFold(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrderNumber), v))
	})
}

// PrepayIDEQ applies the EQ predicate on the "prepay_id" field.
func PrepayIDEQ(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrepayID), v))
	})
}

// PrepayIDNEQ applies the NEQ predicate on the "prepay_id" field.
func PrepayIDNEQ(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrepayID), v))
	})
}

// PrepayIDIn applies the In predicate on the "prepay_id" field.
func PrepayIDIn(vs ...string) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrepayID), v...))
	})
}

// PrepayIDNotIn applies the NotIn predicate on the "prepay_id" field.
func PrepayIDNotIn(vs ...string) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrepayID), v...))
	})
}

// PrepayIDGT applies the GT predicate on the "prepay_id" field.
func PrepayIDGT(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrepayID), v))
	})
}

// PrepayIDGTE applies the GTE predicate on the "prepay_id" field.
func PrepayIDGTE(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrepayID), v))
	})
}

// PrepayIDLT applies the LT predicate on the "prepay_id" field.
func PrepayIDLT(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrepayID), v))
	})
}

// PrepayIDLTE applies the LTE predicate on the "prepay_id" field.
func PrepayIDLTE(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrepayID), v))
	})
}

// PrepayIDContains applies the Contains predicate on the "prepay_id" field.
func PrepayIDContains(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrepayID), v))
	})
}

// PrepayIDHasPrefix applies the HasPrefix predicate on the "prepay_id" field.
func PrepayIDHasPrefix(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrepayID), v))
	})
}

// PrepayIDHasSuffix applies the HasSuffix predicate on the "prepay_id" field.
func PrepayIDHasSuffix(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrepayID), v))
	})
}

// PrepayIDIsNil applies the IsNil predicate on the "prepay_id" field.
func PrepayIDIsNil() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPrepayID)))
	})
}

// PrepayIDNotNil applies the NotNil predicate on the "prepay_id" field.
func PrepayIDNotNil() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPrepayID)))
	})
}

// PrepayIDEqualFold applies the EqualFold predicate on the "prepay_id" field.
func PrepayIDEqualFold(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrepayID), v))
	})
}

// PrepayIDContainsFold applies the ContainsFold predicate on the "prepay_id" field.
func PrepayIDContainsFold(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrepayID), v))
	})
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemark), v))
	})
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func RemarkNotIn(vs ...string) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func RemarkGT(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemark), v))
	})
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemark), v))
	})
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemark), v))
	})
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemark), v))
	})
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemark), v))
	})
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemark), v))
	})
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemark), v))
	})
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRemark)))
	})
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRemark)))
	})
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemark), v))
	})
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemark), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// DeliveryStatusEQ applies the EQ predicate on the "delivery_status" field.
func DeliveryStatusEQ(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeliveryStatus), v))
	})
}

// DeliveryStatusNEQ applies the NEQ predicate on the "delivery_status" field.
func DeliveryStatusNEQ(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeliveryStatus), v))
	})
}

// DeliveryStatusIn applies the In predicate on the "delivery_status" field.
func DeliveryStatusIn(vs ...int8) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeliveryStatus), v...))
	})
}

// DeliveryStatusNotIn applies the NotIn predicate on the "delivery_status" field.
func DeliveryStatusNotIn(vs ...int8) predicate.OrderInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeliveryStatus), v...))
	})
}

// DeliveryStatusGT applies the GT predicate on the "delivery_status" field.
func DeliveryStatusGT(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeliveryStatus), v))
	})
}

// DeliveryStatusGTE applies the GTE predicate on the "delivery_status" field.
func DeliveryStatusGTE(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeliveryStatus), v))
	})
}

// DeliveryStatusLT applies the LT predicate on the "delivery_status" field.
func DeliveryStatusLT(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeliveryStatus), v))
	})
}

// DeliveryStatusLTE applies the LTE predicate on the "delivery_status" field.
func DeliveryStatusLTE(v int8) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeliveryStatus), v))
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderGoodsSku applies the HasEdge predicate on the "order_goods_sku" edge.
func HasOrderGoodsSku() predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderGoodsSkuTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrderGoodsSkuTable, OrderGoodsSkuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderGoodsSkuWith applies the HasEdge predicate on the "order_goods_sku" edge with a given conditions (other predicates).
func HasOrderGoodsSkuWith(preds ...predicate.OrderGoodsSku) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderGoodsSkuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrderGoodsSkuTable, OrderGoodsSkuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderInfo) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderInfo) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
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
func Not(p predicate.OrderInfo) predicate.OrderInfo {
	return predicate.OrderInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}

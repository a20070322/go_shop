// Code generated by entc, DO NOT EDIT.

package goodssku

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/a20070322/shop-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
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
func IDGT(id int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// SkuName applies equality check predicate on the "sku_name" field. It's identical to SkuNameEQ.
func SkuName(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuName), v))
	})
}

// SkuCode applies equality check predicate on the "sku_code" field. It's identical to SkuCodeEQ.
func SkuCode(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuCode), v))
	})
}

// StockNum applies equality check predicate on the "stock_num" field. It's identical to StockNumEQ.
func StockNum(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStockNum), v))
	})
}

// SalesNum applies equality check predicate on the "sales_num" field. It's identical to SalesNumEQ.
func SalesNum(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalesNum), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// SkuNameEQ applies the EQ predicate on the "sku_name" field.
func SkuNameEQ(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuName), v))
	})
}

// SkuNameNEQ applies the NEQ predicate on the "sku_name" field.
func SkuNameNEQ(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkuName), v))
	})
}

// SkuNameIn applies the In predicate on the "sku_name" field.
func SkuNameIn(vs ...string) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSkuName), v...))
	})
}

// SkuNameNotIn applies the NotIn predicate on the "sku_name" field.
func SkuNameNotIn(vs ...string) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSkuName), v...))
	})
}

// SkuNameGT applies the GT predicate on the "sku_name" field.
func SkuNameGT(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSkuName), v))
	})
}

// SkuNameGTE applies the GTE predicate on the "sku_name" field.
func SkuNameGTE(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSkuName), v))
	})
}

// SkuNameLT applies the LT predicate on the "sku_name" field.
func SkuNameLT(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSkuName), v))
	})
}

// SkuNameLTE applies the LTE predicate on the "sku_name" field.
func SkuNameLTE(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSkuName), v))
	})
}

// SkuNameContains applies the Contains predicate on the "sku_name" field.
func SkuNameContains(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSkuName), v))
	})
}

// SkuNameHasPrefix applies the HasPrefix predicate on the "sku_name" field.
func SkuNameHasPrefix(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSkuName), v))
	})
}

// SkuNameHasSuffix applies the HasSuffix predicate on the "sku_name" field.
func SkuNameHasSuffix(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSkuName), v))
	})
}

// SkuNameEqualFold applies the EqualFold predicate on the "sku_name" field.
func SkuNameEqualFold(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSkuName), v))
	})
}

// SkuNameContainsFold applies the ContainsFold predicate on the "sku_name" field.
func SkuNameContainsFold(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSkuName), v))
	})
}

// SkuCodeEQ applies the EQ predicate on the "sku_code" field.
func SkuCodeEQ(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuCode), v))
	})
}

// SkuCodeNEQ applies the NEQ predicate on the "sku_code" field.
func SkuCodeNEQ(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkuCode), v))
	})
}

// SkuCodeIn applies the In predicate on the "sku_code" field.
func SkuCodeIn(vs ...string) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSkuCode), v...))
	})
}

// SkuCodeNotIn applies the NotIn predicate on the "sku_code" field.
func SkuCodeNotIn(vs ...string) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSkuCode), v...))
	})
}

// SkuCodeGT applies the GT predicate on the "sku_code" field.
func SkuCodeGT(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSkuCode), v))
	})
}

// SkuCodeGTE applies the GTE predicate on the "sku_code" field.
func SkuCodeGTE(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSkuCode), v))
	})
}

// SkuCodeLT applies the LT predicate on the "sku_code" field.
func SkuCodeLT(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSkuCode), v))
	})
}

// SkuCodeLTE applies the LTE predicate on the "sku_code" field.
func SkuCodeLTE(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSkuCode), v))
	})
}

// SkuCodeContains applies the Contains predicate on the "sku_code" field.
func SkuCodeContains(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSkuCode), v))
	})
}

// SkuCodeHasPrefix applies the HasPrefix predicate on the "sku_code" field.
func SkuCodeHasPrefix(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSkuCode), v))
	})
}

// SkuCodeHasSuffix applies the HasSuffix predicate on the "sku_code" field.
func SkuCodeHasSuffix(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSkuCode), v))
	})
}

// SkuCodeEqualFold applies the EqualFold predicate on the "sku_code" field.
func SkuCodeEqualFold(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSkuCode), v))
	})
}

// SkuCodeContainsFold applies the ContainsFold predicate on the "sku_code" field.
func SkuCodeContainsFold(v string) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSkuCode), v))
	})
}

// StockNumEQ applies the EQ predicate on the "stock_num" field.
func StockNumEQ(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStockNum), v))
	})
}

// StockNumNEQ applies the NEQ predicate on the "stock_num" field.
func StockNumNEQ(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStockNum), v))
	})
}

// StockNumIn applies the In predicate on the "stock_num" field.
func StockNumIn(vs ...int) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStockNum), v...))
	})
}

// StockNumNotIn applies the NotIn predicate on the "stock_num" field.
func StockNumNotIn(vs ...int) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStockNum), v...))
	})
}

// StockNumGT applies the GT predicate on the "stock_num" field.
func StockNumGT(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStockNum), v))
	})
}

// StockNumGTE applies the GTE predicate on the "stock_num" field.
func StockNumGTE(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStockNum), v))
	})
}

// StockNumLT applies the LT predicate on the "stock_num" field.
func StockNumLT(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStockNum), v))
	})
}

// StockNumLTE applies the LTE predicate on the "stock_num" field.
func StockNumLTE(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStockNum), v))
	})
}

// SalesNumEQ applies the EQ predicate on the "sales_num" field.
func SalesNumEQ(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalesNum), v))
	})
}

// SalesNumNEQ applies the NEQ predicate on the "sales_num" field.
func SalesNumNEQ(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSalesNum), v))
	})
}

// SalesNumIn applies the In predicate on the "sales_num" field.
func SalesNumIn(vs ...int) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSalesNum), v...))
	})
}

// SalesNumNotIn applies the NotIn predicate on the "sales_num" field.
func SalesNumNotIn(vs ...int) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSalesNum), v...))
	})
}

// SalesNumGT applies the GT predicate on the "sales_num" field.
func SalesNumGT(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSalesNum), v))
	})
}

// SalesNumGTE applies the GTE predicate on the "sales_num" field.
func SalesNumGTE(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSalesNum), v))
	})
}

// SalesNumLT applies the LT predicate on the "sales_num" field.
func SalesNumLT(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSalesNum), v))
	})
}

// SalesNumLTE applies the LTE predicate on the "sales_num" field.
func SalesNumLTE(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSalesNum), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.GoodsSku {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSku(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// HasGoodsSpu applies the HasEdge predicate on the "goods_spu" edge.
func HasGoodsSpu() predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSpuTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GoodsSpuTable, GoodsSpuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGoodsSpuWith applies the HasEdge predicate on the "goods_spu" edge with a given conditions (other predicates).
func HasGoodsSpuWith(preds ...predicate.GoodsSpu) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSpuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GoodsSpuTable, GoodsSpuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGoodsSpecsOption applies the HasEdge predicate on the "goods_specs_option" edge.
func HasGoodsSpecsOption() predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSpecsOptionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GoodsSpecsOptionTable, GoodsSpecsOptionPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGoodsSpecsOptionWith applies the HasEdge predicate on the "goods_specs_option" edge with a given conditions (other predicates).
func HasGoodsSpecsOptionWith(preds ...predicate.GoodsSpecsOption) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSpecsOptionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GoodsSpecsOptionTable, GoodsSpecsOptionPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodsSku) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodsSku) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
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
func Not(p predicate.GoodsSku) predicate.GoodsSku {
	return predicate.GoodsSku(func(s *sql.Selector) {
		p(s.Not())
	})
}
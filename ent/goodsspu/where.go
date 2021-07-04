// Code generated by entc, DO NOT EDIT.

package goodsspu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/a20070322/shop-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func IDGT(id int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// SpuName applies equality check predicate on the "spu_name" field. It's identical to SpuNameEQ.
func SpuName(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuName), v))
	})
}

// SpuCode applies equality check predicate on the "spu_code" field. It's identical to SpuCodeEQ.
func SpuCode(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuCode), v))
	})
}

// SpuHeadImg applies equality check predicate on the "spu_head_img" field. It's identical to SpuHeadImgEQ.
func SpuHeadImg(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuHeadImg), v))
	})
}

// SalesNum applies equality check predicate on the "sales_num" field. It's identical to SalesNumEQ.
func SalesNum(v int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalesNum), v))
	})
}

// SpuDesc applies equality check predicate on the "spu_desc" field. It's identical to SpuDescEQ.
func SpuDesc(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuDesc), v))
	})
}

// SpuDetails applies equality check predicate on the "spu_details" field. It's identical to SpuDetailsEQ.
func SpuDetails(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuDetails), v))
	})
}

// IsCustomSku applies equality check predicate on the "is_custom_sku" field. It's identical to IsCustomSkuEQ.
func IsCustomSku(v bool) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsCustomSku), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// SpuNameEQ applies the EQ predicate on the "spu_name" field.
func SpuNameEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuName), v))
	})
}

// SpuNameNEQ applies the NEQ predicate on the "spu_name" field.
func SpuNameNEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpuName), v))
	})
}

// SpuNameIn applies the In predicate on the "spu_name" field.
func SpuNameIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSpuName), v...))
	})
}

// SpuNameNotIn applies the NotIn predicate on the "spu_name" field.
func SpuNameNotIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpuName), v...))
	})
}

// SpuNameGT applies the GT predicate on the "spu_name" field.
func SpuNameGT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpuName), v))
	})
}

// SpuNameGTE applies the GTE predicate on the "spu_name" field.
func SpuNameGTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpuName), v))
	})
}

// SpuNameLT applies the LT predicate on the "spu_name" field.
func SpuNameLT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpuName), v))
	})
}

// SpuNameLTE applies the LTE predicate on the "spu_name" field.
func SpuNameLTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpuName), v))
	})
}

// SpuNameContains applies the Contains predicate on the "spu_name" field.
func SpuNameContains(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSpuName), v))
	})
}

// SpuNameHasPrefix applies the HasPrefix predicate on the "spu_name" field.
func SpuNameHasPrefix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSpuName), v))
	})
}

// SpuNameHasSuffix applies the HasSuffix predicate on the "spu_name" field.
func SpuNameHasSuffix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSpuName), v))
	})
}

// SpuNameEqualFold applies the EqualFold predicate on the "spu_name" field.
func SpuNameEqualFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSpuName), v))
	})
}

// SpuNameContainsFold applies the ContainsFold predicate on the "spu_name" field.
func SpuNameContainsFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSpuName), v))
	})
}

// SpuCodeEQ applies the EQ predicate on the "spu_code" field.
func SpuCodeEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuCode), v))
	})
}

// SpuCodeNEQ applies the NEQ predicate on the "spu_code" field.
func SpuCodeNEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpuCode), v))
	})
}

// SpuCodeIn applies the In predicate on the "spu_code" field.
func SpuCodeIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSpuCode), v...))
	})
}

// SpuCodeNotIn applies the NotIn predicate on the "spu_code" field.
func SpuCodeNotIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpuCode), v...))
	})
}

// SpuCodeGT applies the GT predicate on the "spu_code" field.
func SpuCodeGT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpuCode), v))
	})
}

// SpuCodeGTE applies the GTE predicate on the "spu_code" field.
func SpuCodeGTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpuCode), v))
	})
}

// SpuCodeLT applies the LT predicate on the "spu_code" field.
func SpuCodeLT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpuCode), v))
	})
}

// SpuCodeLTE applies the LTE predicate on the "spu_code" field.
func SpuCodeLTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpuCode), v))
	})
}

// SpuCodeContains applies the Contains predicate on the "spu_code" field.
func SpuCodeContains(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSpuCode), v))
	})
}

// SpuCodeHasPrefix applies the HasPrefix predicate on the "spu_code" field.
func SpuCodeHasPrefix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSpuCode), v))
	})
}

// SpuCodeHasSuffix applies the HasSuffix predicate on the "spu_code" field.
func SpuCodeHasSuffix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSpuCode), v))
	})
}

// SpuCodeEqualFold applies the EqualFold predicate on the "spu_code" field.
func SpuCodeEqualFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSpuCode), v))
	})
}

// SpuCodeContainsFold applies the ContainsFold predicate on the "spu_code" field.
func SpuCodeContainsFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSpuCode), v))
	})
}

// SpuHeadImgEQ applies the EQ predicate on the "spu_head_img" field.
func SpuHeadImgEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgNEQ applies the NEQ predicate on the "spu_head_img" field.
func SpuHeadImgNEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgIn applies the In predicate on the "spu_head_img" field.
func SpuHeadImgIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSpuHeadImg), v...))
	})
}

// SpuHeadImgNotIn applies the NotIn predicate on the "spu_head_img" field.
func SpuHeadImgNotIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpuHeadImg), v...))
	})
}

// SpuHeadImgGT applies the GT predicate on the "spu_head_img" field.
func SpuHeadImgGT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgGTE applies the GTE predicate on the "spu_head_img" field.
func SpuHeadImgGTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgLT applies the LT predicate on the "spu_head_img" field.
func SpuHeadImgLT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgLTE applies the LTE predicate on the "spu_head_img" field.
func SpuHeadImgLTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgContains applies the Contains predicate on the "spu_head_img" field.
func SpuHeadImgContains(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgHasPrefix applies the HasPrefix predicate on the "spu_head_img" field.
func SpuHeadImgHasPrefix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgHasSuffix applies the HasSuffix predicate on the "spu_head_img" field.
func SpuHeadImgHasSuffix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgIsNil applies the IsNil predicate on the "spu_head_img" field.
func SpuHeadImgIsNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSpuHeadImg)))
	})
}

// SpuHeadImgNotNil applies the NotNil predicate on the "spu_head_img" field.
func SpuHeadImgNotNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSpuHeadImg)))
	})
}

// SpuHeadImgEqualFold applies the EqualFold predicate on the "spu_head_img" field.
func SpuHeadImgEqualFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSpuHeadImg), v))
	})
}

// SpuHeadImgContainsFold applies the ContainsFold predicate on the "spu_head_img" field.
func SpuHeadImgContainsFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSpuHeadImg), v))
	})
}

// SalesNumEQ applies the EQ predicate on the "sales_num" field.
func SalesNumEQ(v int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalesNum), v))
	})
}

// SalesNumNEQ applies the NEQ predicate on the "sales_num" field.
func SalesNumNEQ(v int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSalesNum), v))
	})
}

// SalesNumIn applies the In predicate on the "sales_num" field.
func SalesNumIn(vs ...int) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func SalesNumNotIn(vs ...int) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func SalesNumGT(v int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSalesNum), v))
	})
}

// SalesNumGTE applies the GTE predicate on the "sales_num" field.
func SalesNumGTE(v int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSalesNum), v))
	})
}

// SalesNumLT applies the LT predicate on the "sales_num" field.
func SalesNumLT(v int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSalesNum), v))
	})
}

// SalesNumLTE applies the LTE predicate on the "sales_num" field.
func SalesNumLTE(v int) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSalesNum), v))
	})
}

// SalesNumIsNil applies the IsNil predicate on the "sales_num" field.
func SalesNumIsNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSalesNum)))
	})
}

// SalesNumNotNil applies the NotNil predicate on the "sales_num" field.
func SalesNumNotNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSalesNum)))
	})
}

// SpuDescEQ applies the EQ predicate on the "spu_desc" field.
func SpuDescEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuDesc), v))
	})
}

// SpuDescNEQ applies the NEQ predicate on the "spu_desc" field.
func SpuDescNEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpuDesc), v))
	})
}

// SpuDescIn applies the In predicate on the "spu_desc" field.
func SpuDescIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSpuDesc), v...))
	})
}

// SpuDescNotIn applies the NotIn predicate on the "spu_desc" field.
func SpuDescNotIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpuDesc), v...))
	})
}

// SpuDescGT applies the GT predicate on the "spu_desc" field.
func SpuDescGT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpuDesc), v))
	})
}

// SpuDescGTE applies the GTE predicate on the "spu_desc" field.
func SpuDescGTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpuDesc), v))
	})
}

// SpuDescLT applies the LT predicate on the "spu_desc" field.
func SpuDescLT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpuDesc), v))
	})
}

// SpuDescLTE applies the LTE predicate on the "spu_desc" field.
func SpuDescLTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpuDesc), v))
	})
}

// SpuDescContains applies the Contains predicate on the "spu_desc" field.
func SpuDescContains(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSpuDesc), v))
	})
}

// SpuDescHasPrefix applies the HasPrefix predicate on the "spu_desc" field.
func SpuDescHasPrefix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSpuDesc), v))
	})
}

// SpuDescHasSuffix applies the HasSuffix predicate on the "spu_desc" field.
func SpuDescHasSuffix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSpuDesc), v))
	})
}

// SpuDescIsNil applies the IsNil predicate on the "spu_desc" field.
func SpuDescIsNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSpuDesc)))
	})
}

// SpuDescNotNil applies the NotNil predicate on the "spu_desc" field.
func SpuDescNotNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSpuDesc)))
	})
}

// SpuDescEqualFold applies the EqualFold predicate on the "spu_desc" field.
func SpuDescEqualFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSpuDesc), v))
	})
}

// SpuDescContainsFold applies the ContainsFold predicate on the "spu_desc" field.
func SpuDescContainsFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSpuDesc), v))
	})
}

// SpuDetailsEQ applies the EQ predicate on the "spu_details" field.
func SpuDetailsEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsNEQ applies the NEQ predicate on the "spu_details" field.
func SpuDetailsNEQ(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsIn applies the In predicate on the "spu_details" field.
func SpuDetailsIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSpuDetails), v...))
	})
}

// SpuDetailsNotIn applies the NotIn predicate on the "spu_details" field.
func SpuDetailsNotIn(vs ...string) predicate.GoodsSpu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodsSpu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSpuDetails), v...))
	})
}

// SpuDetailsGT applies the GT predicate on the "spu_details" field.
func SpuDetailsGT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsGTE applies the GTE predicate on the "spu_details" field.
func SpuDetailsGTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsLT applies the LT predicate on the "spu_details" field.
func SpuDetailsLT(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsLTE applies the LTE predicate on the "spu_details" field.
func SpuDetailsLTE(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsContains applies the Contains predicate on the "spu_details" field.
func SpuDetailsContains(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsHasPrefix applies the HasPrefix predicate on the "spu_details" field.
func SpuDetailsHasPrefix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsHasSuffix applies the HasSuffix predicate on the "spu_details" field.
func SpuDetailsHasSuffix(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsIsNil applies the IsNil predicate on the "spu_details" field.
func SpuDetailsIsNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSpuDetails)))
	})
}

// SpuDetailsNotNil applies the NotNil predicate on the "spu_details" field.
func SpuDetailsNotNil() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSpuDetails)))
	})
}

// SpuDetailsEqualFold applies the EqualFold predicate on the "spu_details" field.
func SpuDetailsEqualFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSpuDetails), v))
	})
}

// SpuDetailsContainsFold applies the ContainsFold predicate on the "spu_details" field.
func SpuDetailsContainsFold(v string) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSpuDetails), v))
	})
}

// IsCustomSkuEQ applies the EQ predicate on the "is_custom_sku" field.
func IsCustomSkuEQ(v bool) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsCustomSku), v))
	})
}

// IsCustomSkuNEQ applies the NEQ predicate on the "is_custom_sku" field.
func IsCustomSkuNEQ(v bool) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsCustomSku), v))
	})
}

// HasGoodsClassify applies the HasEdge predicate on the "goods_classify" edge.
func HasGoodsClassify() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsClassifyTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GoodsClassifyTable, GoodsClassifyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGoodsClassifyWith applies the HasEdge predicate on the "goods_classify" edge with a given conditions (other predicates).
func HasGoodsClassifyWith(preds ...predicate.GoodsClassify) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsClassifyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GoodsClassifyTable, GoodsClassifyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGoodsSku applies the HasEdge predicate on the "goods_sku" edge.
func HasGoodsSku() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSkuTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GoodsSkuTable, GoodsSkuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGoodsSkuWith applies the HasEdge predicate on the "goods_sku" edge with a given conditions (other predicates).
func HasGoodsSkuWith(preds ...predicate.GoodsSku) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSkuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GoodsSkuTable, GoodsSkuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderGoodsSku applies the HasEdge predicate on the "order_goods_sku" edge.
func HasOrderGoodsSku() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderGoodsSkuTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrderGoodsSkuTable, OrderGoodsSkuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderGoodsSkuWith applies the HasEdge predicate on the "order_goods_sku" edge with a given conditions (other predicates).
func HasOrderGoodsSkuWith(preds ...predicate.OrderGoodsSku) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
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

// HasGoodsSpuImgs applies the HasEdge predicate on the "goods_spu_imgs" edge.
func HasGoodsSpuImgs() predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSpuImgsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GoodsSpuImgsTable, GoodsSpuImgsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGoodsSpuImgsWith applies the HasEdge predicate on the "goods_spu_imgs" edge with a given conditions (other predicates).
func HasGoodsSpuImgsWith(preds ...predicate.GoodsSpuImgs) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GoodsSpuImgsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GoodsSpuImgsTable, GoodsSpuImgsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodsSpu) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodsSpu) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
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
func Not(p predicate.GoodsSpu) predicate.GoodsSpu {
	return predicate.GoodsSpu(func(s *sql.Selector) {
		p(s.Not())
	})
}

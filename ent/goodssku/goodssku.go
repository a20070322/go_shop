// Code generated by entc, DO NOT EDIT.

package goodssku

import (
	"time"
)

const (
	// Label holds the string label denoting the goodssku type in the database.
	Label = "goods_sku"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldSkuName holds the string denoting the sku_name field in the database.
	FieldSkuName = "sku_name"
	// FieldSkuCode holds the string denoting the sku_code field in the database.
	FieldSkuCode = "sku_code"
	// FieldStockNum holds the string denoting the stock_num field in the database.
	FieldStockNum = "stock_num"
	// FieldSalesNum holds the string denoting the sales_num field in the database.
	FieldSalesNum = "sales_num"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// EdgeGoodsSpu holds the string denoting the goods_spu edge name in mutations.
	EdgeGoodsSpu = "goods_spu"
	// EdgeGoodsSpecsOption holds the string denoting the goods_specs_option edge name in mutations.
	EdgeGoodsSpecsOption = "goods_specs_option"
	// Table holds the table name of the goodssku in the database.
	Table = "goods_skus"
	// GoodsSpuTable is the table the holds the goods_spu relation/edge.
	GoodsSpuTable = "goods_skus"
	// GoodsSpuInverseTable is the table name for the GoodsSpu entity.
	// It exists in this package in order to avoid circular dependency with the "goodsspu" package.
	GoodsSpuInverseTable = "goods_spus"
	// GoodsSpuColumn is the table column denoting the goods_spu relation/edge.
	GoodsSpuColumn = "goods_spu_goods_sku"
	// GoodsSpecsOptionTable is the table the holds the goods_specs_option relation/edge. The primary key declared below.
	GoodsSpecsOptionTable = "goods_specs_option_goods_sku"
	// GoodsSpecsOptionInverseTable is the table name for the GoodsSpecsOption entity.
	// It exists in this package in order to avoid circular dependency with the "goodsspecsoption" package.
	GoodsSpecsOptionInverseTable = "goods_specs_options"
)

// Columns holds all SQL columns for goodssku fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldSkuName,
	FieldSkuCode,
	FieldStockNum,
	FieldSalesNum,
	FieldPrice,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "goods_skus"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"goods_spu_goods_sku",
}

var (
	// GoodsSpecsOptionPrimaryKey and GoodsSpecsOptionColumn2 are the table columns denoting the
	// primary key for the goods_specs_option relation (M2M).
	GoodsSpecsOptionPrimaryKey = []string{"goods_specs_option_id", "goods_sku_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStockNum holds the default value on creation for the "stock_num" field.
	DefaultStockNum int
	// DefaultSalesNum holds the default value on creation for the "sales_num" field.
	DefaultSalesNum int
)

// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BasicBannersColumns holds the columns for the "basic_banners" table.
	BasicBannersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "banner_name", Type: field.TypeString},
		{Name: "banner_img", Type: field.TypeString},
		{Name: "order", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "status", Type: field.TypeBool, Default: true},
		{Name: "basic_banner_position_basic_banner", Type: field.TypeInt, Nullable: true},
		{Name: "basic_link_basic_banner", Type: field.TypeInt, Nullable: true},
	}
	// BasicBannersTable holds the schema information for the "basic_banners" table.
	BasicBannersTable = &schema.Table{
		Name:       "basic_banners",
		Columns:    BasicBannersColumns,
		PrimaryKey: []*schema.Column{BasicBannersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "basic_banners_basic_banner_positions_basic_banner",
				Columns:    []*schema.Column{BasicBannersColumns[8]},
				RefColumns: []*schema.Column{BasicBannerPositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "basic_banners_basic_links_basic_banner",
				Columns:    []*schema.Column{BasicBannersColumns[9]},
				RefColumns: []*schema.Column{BasicLinksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BasicBannerPositionsColumns holds the columns for the "basic_banner_positions" table.
	BasicBannerPositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "position_name", Type: field.TypeString},
		{Name: "remarks", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeBool, Default: true},
	}
	// BasicBannerPositionsTable holds the schema information for the "basic_banner_positions" table.
	BasicBannerPositionsTable = &schema.Table{
		Name:        "basic_banner_positions",
		Columns:     BasicBannerPositionsColumns,
		PrimaryKey:  []*schema.Column{BasicBannerPositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// BasicLinksColumns holds the columns for the "basic_links" table.
	BasicLinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "link_name", Type: field.TypeString},
		{Name: "link_type", Type: field.TypeString},
		{Name: "link_address", Type: field.TypeString},
		{Name: "appid", Type: field.TypeString, Nullable: true},
		{Name: "is_register", Type: field.TypeBool, Default: false},
		{Name: "remarks", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeBool, Default: true},
	}
	// BasicLinksTable holds the schema information for the "basic_links" table.
	BasicLinksTable = &schema.Table{
		Name:        "basic_links",
		Columns:     BasicLinksColumns,
		PrimaryKey:  []*schema.Column{BasicLinksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "mini_openid", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "wechat_openid", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "union_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "is_disable", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:        "customers",
		Columns:     CustomersColumns,
		PrimaryKey:  []*schema.Column{CustomersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CustomerAddressesColumns holds the columns for the "customer_addresses" table.
	CustomerAddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "province", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "area", Type: field.TypeString},
		{Name: "detailed", Type: field.TypeString},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "is_default", Type: field.TypeBool, Default: false},
		{Name: "customer_address", Type: field.TypeInt, Nullable: true},
	}
	// CustomerAddressesTable holds the schema information for the "customer_addresses" table.
	CustomerAddressesTable = &schema.Table{
		Name:       "customer_addresses",
		Columns:    CustomerAddressesColumns,
		PrimaryKey: []*schema.Column{CustomerAddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "customer_addresses_customers_address",
				Columns:    []*schema.Column{CustomerAddressesColumns[12]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GoodsClassifiesColumns holds the columns for the "goods_classifies" table.
	GoodsClassifiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "classify_name", Type: field.TypeString, Nullable: true},
		{Name: "classify_code", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "pid", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "level", Type: field.TypeInt},
		{Name: "sort", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "icon", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "is_disable", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// GoodsClassifiesTable holds the schema information for the "goods_classifies" table.
	GoodsClassifiesTable = &schema.Table{
		Name:        "goods_classifies",
		Columns:     GoodsClassifiesColumns,
		PrimaryKey:  []*schema.Column{GoodsClassifiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GoodsSkusColumns holds the columns for the "goods_skus" table.
	GoodsSkusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "sku_name", Type: field.TypeString},
		{Name: "sku_code", Type: field.TypeString, Unique: true},
		{Name: "stock_num", Type: field.TypeInt, Default: 0},
		{Name: "sales_num", Type: field.TypeInt, Default: 0},
		{Name: "price", Type: field.TypeInt},
		{Name: "goods_spu_goods_sku", Type: field.TypeInt, Nullable: true},
	}
	// GoodsSkusTable holds the schema information for the "goods_skus" table.
	GoodsSkusTable = &schema.Table{
		Name:       "goods_skus",
		Columns:    GoodsSkusColumns,
		PrimaryKey: []*schema.Column{GoodsSkusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "goods_skus_goods_spus_goods_sku",
				Columns:    []*schema.Column{GoodsSkusColumns[9]},
				RefColumns: []*schema.Column{GoodsSpusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GoodsSpecsColumns holds the columns for the "goods_specs" table.
	GoodsSpecsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "specs_name", Type: field.TypeString, Unique: true},
	}
	// GoodsSpecsTable holds the schema information for the "goods_specs" table.
	GoodsSpecsTable = &schema.Table{
		Name:        "goods_specs",
		Columns:     GoodsSpecsColumns,
		PrimaryKey:  []*schema.Column{GoodsSpecsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GoodsSpecsOptionsColumns holds the columns for the "goods_specs_options" table.
	GoodsSpecsOptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "specs_option_value", Type: field.TypeString, Unique: true},
		{Name: "goods_specs_goods_specs_option", Type: field.TypeInt, Nullable: true},
	}
	// GoodsSpecsOptionsTable holds the schema information for the "goods_specs_options" table.
	GoodsSpecsOptionsTable = &schema.Table{
		Name:       "goods_specs_options",
		Columns:    GoodsSpecsOptionsColumns,
		PrimaryKey: []*schema.Column{GoodsSpecsOptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "goods_specs_options_goods_specs_goods_specs_option",
				Columns:    []*schema.Column{GoodsSpecsOptionsColumns[5]},
				RefColumns: []*schema.Column{GoodsSpecsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GoodsSpusColumns holds the columns for the "goods_spus" table.
	GoodsSpusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "spu_name", Type: field.TypeString},
		{Name: "spu_code", Type: field.TypeString, Unique: true},
		{Name: "spu_head_img", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "sales_num", Type: field.TypeInt, Nullable: true},
		{Name: "spu_desc", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "spu_details", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "is_custom_sku", Type: field.TypeBool},
		{Name: "goods_classify_goods_spu", Type: field.TypeInt, Nullable: true},
	}
	// GoodsSpusTable holds the schema information for the "goods_spus" table.
	GoodsSpusTable = &schema.Table{
		Name:       "goods_spus",
		Columns:    GoodsSpusColumns,
		PrimaryKey: []*schema.Column{GoodsSpusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "goods_spus_goods_classifies_goods_spu",
				Columns:    []*schema.Column{GoodsSpusColumns[11]},
				RefColumns: []*schema.Column{GoodsClassifiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GoodsSpuImgsColumns holds the columns for the "goods_spu_imgs" table.
	GoodsSpuImgsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "img_name", Type: field.TypeString, Nullable: true},
		{Name: "img_path", Type: field.TypeString, Nullable: true},
		{Name: "goods_spu_goods_spu_imgs", Type: field.TypeInt, Nullable: true},
	}
	// GoodsSpuImgsTable holds the schema information for the "goods_spu_imgs" table.
	GoodsSpuImgsTable = &schema.Table{
		Name:       "goods_spu_imgs",
		Columns:    GoodsSpuImgsColumns,
		PrimaryKey: []*schema.Column{GoodsSpuImgsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "goods_spu_imgs_goods_spus_goods_spu_imgs",
				Columns:    []*schema.Column{GoodsSpuImgsColumns[6]},
				RefColumns: []*schema.Column{GoodsSpusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderAddressesColumns holds the columns for the "order_addresses" table.
	OrderAddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "province", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "area", Type: field.TypeString},
		{Name: "detailed", Type: field.TypeString},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "order_info_order_address", Type: field.TypeInt, Nullable: true},
	}
	// OrderAddressesTable holds the schema information for the "order_addresses" table.
	OrderAddressesTable = &schema.Table{
		Name:       "order_addresses",
		Columns:    OrderAddressesColumns,
		PrimaryKey: []*schema.Column{OrderAddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_addresses_order_infos_order_address",
				Columns:    []*schema.Column{OrderAddressesColumns[8]},
				RefColumns: []*schema.Column{OrderInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderGoodsSkusColumns holds the columns for the "order_goods_skus" table.
	OrderGoodsSkusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "spu_name", Type: field.TypeString},
		{Name: "spu_code", Type: field.TypeString},
		{Name: "spu_head_img", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "spu_desc", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "spu_details", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "is_custom_sku", Type: field.TypeBool},
		{Name: "specs_option", Type: field.TypeJSON},
		{Name: "sku_id", Type: field.TypeInt},
		{Name: "sku_name", Type: field.TypeString},
		{Name: "sku_code", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt},
		{Name: "amount", Type: field.TypeInt},
		{Name: "goods_spu_order_goods_sku", Type: field.TypeInt, Nullable: true},
		{Name: "order_info_order_goods_sku", Type: field.TypeInt, Nullable: true},
	}
	// OrderGoodsSkusTable holds the schema information for the "order_goods_skus" table.
	OrderGoodsSkusTable = &schema.Table{
		Name:       "order_goods_skus",
		Columns:    OrderGoodsSkusColumns,
		PrimaryKey: []*schema.Column{OrderGoodsSkusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_goods_skus_goods_spus_order_goods_sku",
				Columns:    []*schema.Column{OrderGoodsSkusColumns[16]},
				RefColumns: []*schema.Column{GoodsSpusColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "order_goods_skus_order_infos_order_goods_sku",
				Columns:    []*schema.Column{OrderGoodsSkusColumns[17]},
				RefColumns: []*schema.Column{OrderInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrderInfosColumns holds the columns for the "order_infos" table.
	OrderInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "order_number", Type: field.TypeString, Unique: true},
		{Name: "pay_method", Type: field.TypeInt8, Nullable: true},
		{Name: "pay_money", Type: field.TypeInt},
		{Name: "remark", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "delivery_status", Type: field.TypeInt8, Default: 0},
		{Name: "customer_order_info", Type: field.TypeInt, Nullable: true},
	}
	// OrderInfosTable holds the schema information for the "order_infos" table.
	OrderInfosTable = &schema.Table{
		Name:       "order_infos",
		Columns:    OrderInfosColumns,
		PrimaryKey: []*schema.Column{OrderInfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_infos_customers_order_info",
				Columns:    []*schema.Column{OrderInfosColumns[10]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WeChatPaysColumns holds the columns for the "we_chat_pays" table.
	WeChatPaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "out_trade_no", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "transaction_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "trade_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"JSAPI", "NATIVE", "APP"}},
		{Name: "bank_type", Type: field.TypeString, Nullable: true},
		{Name: "success_time", Type: field.TypeTime, Nullable: true},
		{Name: "payer_currency", Type: field.TypeString, Nullable: true},
		{Name: "payer_total", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "trade_state", Type: field.TypeInt8, Nullable: true, Default: 0},
		{Name: "order_info_wechat_pay", Type: field.TypeInt, Nullable: true},
	}
	// WeChatPaysTable holds the schema information for the "we_chat_pays" table.
	WeChatPaysTable = &schema.Table{
		Name:       "we_chat_pays",
		Columns:    WeChatPaysColumns,
		PrimaryKey: []*schema.Column{WeChatPaysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "we_chat_pays_order_infos_wechat_pay",
				Columns:    []*schema.Column{WeChatPaysColumns[12]},
				RefColumns: []*schema.Column{OrderInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GoodsSpecsOptionGoodsSkuColumns holds the columns for the "goods_specs_option_goods_sku" table.
	GoodsSpecsOptionGoodsSkuColumns = []*schema.Column{
		{Name: "goods_specs_option_id", Type: field.TypeInt},
		{Name: "goods_sku_id", Type: field.TypeInt},
	}
	// GoodsSpecsOptionGoodsSkuTable holds the schema information for the "goods_specs_option_goods_sku" table.
	GoodsSpecsOptionGoodsSkuTable = &schema.Table{
		Name:       "goods_specs_option_goods_sku",
		Columns:    GoodsSpecsOptionGoodsSkuColumns,
		PrimaryKey: []*schema.Column{GoodsSpecsOptionGoodsSkuColumns[0], GoodsSpecsOptionGoodsSkuColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "goods_specs_option_goods_sku_goods_specs_option_id",
				Columns:    []*schema.Column{GoodsSpecsOptionGoodsSkuColumns[0]},
				RefColumns: []*schema.Column{GoodsSpecsOptionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "goods_specs_option_goods_sku_goods_sku_id",
				Columns:    []*schema.Column{GoodsSpecsOptionGoodsSkuColumns[1]},
				RefColumns: []*schema.Column{GoodsSkusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BasicBannersTable,
		BasicBannerPositionsTable,
		BasicLinksTable,
		CustomersTable,
		CustomerAddressesTable,
		GoodsClassifiesTable,
		GoodsSkusTable,
		GoodsSpecsTable,
		GoodsSpecsOptionsTable,
		GoodsSpusTable,
		GoodsSpuImgsTable,
		OrderAddressesTable,
		OrderGoodsSkusTable,
		OrderInfosTable,
		WeChatPaysTable,
		GoodsSpecsOptionGoodsSkuTable,
	}
)

func init() {
	BasicBannersTable.ForeignKeys[0].RefTable = BasicBannerPositionsTable
	BasicBannersTable.ForeignKeys[1].RefTable = BasicLinksTable
	CustomerAddressesTable.ForeignKeys[0].RefTable = CustomersTable
	GoodsSkusTable.ForeignKeys[0].RefTable = GoodsSpusTable
	GoodsSpecsOptionsTable.ForeignKeys[0].RefTable = GoodsSpecsTable
	GoodsSpusTable.ForeignKeys[0].RefTable = GoodsClassifiesTable
	GoodsSpuImgsTable.ForeignKeys[0].RefTable = GoodsSpusTable
	OrderAddressesTable.ForeignKeys[0].RefTable = OrderInfosTable
	OrderGoodsSkusTable.ForeignKeys[0].RefTable = GoodsSpusTable
	OrderGoodsSkusTable.ForeignKeys[1].RefTable = OrderInfosTable
	OrderInfosTable.ForeignKeys[0].RefTable = CustomersTable
	WeChatPaysTable.ForeignKeys[0].RefTable = OrderInfosTable
	GoodsSpecsOptionGoodsSkuTable.ForeignKeys[0].RefTable = GoodsSpecsOptionsTable
	GoodsSpecsOptionGoodsSkuTable.ForeignKeys[1].RefTable = GoodsSkusTable
}

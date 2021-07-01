// Code generated by entc, DO NOT EDIT.

package basiclink

import (
	"time"
)

const (
	// Label holds the string label denoting the basiclink type in the database.
	Label = "basic_link"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldLinkName holds the string denoting the link_name field in the database.
	FieldLinkName = "link_name"
	// FieldLinkType holds the string denoting the link_type field in the database.
	FieldLinkType = "link_type"
	// FieldLinkAddress holds the string denoting the link_address field in the database.
	FieldLinkAddress = "link_address"
	// FieldAppid holds the string denoting the appid field in the database.
	FieldAppid = "appid"
	// FieldIsRegister holds the string denoting the is_register field in the database.
	FieldIsRegister = "is_register"
	// FieldRemarks holds the string denoting the remarks field in the database.
	FieldRemarks = "remarks"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeBasicBanner holds the string denoting the basic_banner edge name in mutations.
	EdgeBasicBanner = "basic_banner"
	// Table holds the table name of the basiclink in the database.
	Table = "basic_links"
	// BasicBannerTable is the table the holds the basic_banner relation/edge.
	BasicBannerTable = "basic_banners"
	// BasicBannerInverseTable is the table name for the BasicBanner entity.
	// It exists in this package in order to avoid circular dependency with the "basicbanner" package.
	BasicBannerInverseTable = "basic_banners"
	// BasicBannerColumn is the table column denoting the basic_banner relation/edge.
	BasicBannerColumn = "basic_link_basic_banner"
)

// Columns holds all SQL columns for basiclink fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldLinkName,
	FieldLinkType,
	FieldLinkAddress,
	FieldAppid,
	FieldIsRegister,
	FieldRemarks,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// DefaultIsRegister holds the default value on creation for the "is_register" field.
	DefaultIsRegister bool
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
)

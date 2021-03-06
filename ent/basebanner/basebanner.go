// Code generated by entc, DO NOT EDIT.

package basebanner

import (
	"time"
)

const (
	// Label holds the string label denoting the basebanner type in the database.
	Label = "base_banner"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldBannerName holds the string denoting the banner_name field in the database.
	FieldBannerName = "banner_name"
	// FieldBannerImg holds the string denoting the banner_img field in the database.
	FieldBannerImg = "banner_img"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the basebanner in the database.
	Table = "base_banners"
)

// Columns holds all SQL columns for basebanner fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldBannerName,
	FieldBannerImg,
	FieldOrder,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
)

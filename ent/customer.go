// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/a20070322/shop-go/ent/customer"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"-"`
	// MiniOpenid holds the value of the "mini_openid" field.
	// 小程序openid
	MiniOpenid string `json:"mini_openid,omitempty"`
	// WechatOpenid holds the value of the "wechat_openid" field.
	// 公众号openid
	WechatOpenid string `json:"wechat_openid,omitempty"`
	// UnionID holds the value of the "union_id" field.
	// 开放平台
	UnionID string `json:"union_id,omitempty"`
	// Phone holds the value of the "phone" field.
	// 手机号
	Phone string `json:"phone,omitempty"`
	// Avatar holds the value of the "avatar" field.
	// 头像
	Avatar string `json:"avatar,omitempty"`
	// IsDisable holds the value of the "is_disable" field.
	// 是否禁用
	IsDisable bool `json:"is_disable,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges CustomerEdges `json:"edges"`
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Address holds the value of the address edge.
	Address []*CustomerAddress `json:"address,omitempty"`
	// OrderInfo holds the value of the order_info edge.
	OrderInfo []*OrderInfo `json:"order_info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AddressOrErr returns the Address value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) AddressOrErr() ([]*CustomerAddress, error) {
	if e.loadedTypes[0] {
		return e.Address, nil
	}
	return nil, &NotLoadedError{edge: "address"}
}

// OrderInfoOrErr returns the OrderInfo value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) OrderInfoOrErr() ([]*OrderInfo, error) {
	if e.loadedTypes[1] {
		return e.OrderInfo, nil
	}
	return nil, &NotLoadedError{edge: "order_info"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldIsDisable:
			values[i] = new(sql.NullBool)
		case customer.FieldID:
			values[i] = new(sql.NullInt64)
		case customer.FieldMiniOpenid, customer.FieldWechatOpenid, customer.FieldUnionID, customer.FieldPhone, customer.FieldAvatar:
			values[i] = new(sql.NullString)
		case customer.FieldCreatedAt, customer.FieldUpdatedAt, customer.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Customer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case customer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case customer.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case customer.FieldMiniOpenid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mini_openid", values[i])
			} else if value.Valid {
				c.MiniOpenid = value.String
			}
		case customer.FieldWechatOpenid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wechat_openid", values[i])
			} else if value.Valid {
				c.WechatOpenid = value.String
			}
		case customer.FieldUnionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field union_id", values[i])
			} else if value.Valid {
				c.UnionID = value.String
			}
		case customer.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case customer.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				c.Avatar = value.String
			}
		case customer.FieldIsDisable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_disable", values[i])
			} else if value.Valid {
				c.IsDisable = value.Bool
			}
		}
	}
	return nil
}

// QueryAddress queries the "address" edge of the Customer entity.
func (c *Customer) QueryAddress() *CustomerAddressQuery {
	return (&CustomerClient{config: c.config}).QueryAddress(c)
}

// QueryOrderInfo queries the "order_info" edge of the Customer entity.
func (c *Customer) QueryOrderInfo() *OrderInfoQuery {
	return (&CustomerClient{config: c.config}).QueryOrderInfo(c)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", mini_openid=")
	builder.WriteString(c.MiniOpenid)
	builder.WriteString(", wechat_openid=")
	builder.WriteString(c.WechatOpenid)
	builder.WriteString(", union_id=")
	builder.WriteString(c.UnionID)
	builder.WriteString(", phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", avatar=")
	builder.WriteString(c.Avatar)
	builder.WriteString(", is_disable=")
	builder.WriteString(fmt.Sprintf("%v", c.IsDisable))
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}

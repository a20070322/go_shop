// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/a20070322/shop-go/ent/goodsclassify"
)

// GoodsClassify is the model entity for the GoodsClassify schema.
type GoodsClassify struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"-"`
	// ClassifyName holds the value of the "classify_name" field.
	// 分类名称
	ClassifyName string `json:"classify_name,omitempty"`
	// ClassifyCode holds the value of the "classify_code" field.
	// 分类编码
	ClassifyCode string `json:"classify_code,omitempty"`
	// Pid holds the value of the "pid" field.
	// 父级id
	Pid int `json:"pid,omitempty"`
	// Level holds the value of the "level" field.
	// 级别
	Level int `json:"level,omitempty"`
	// Sort holds the value of the "sort" field.
	// 排序
	Sort int `json:"sort,omitempty"`
	// Icon holds the value of the "icon" field.
	// icon图标
	Icon string `json:"icon,omitempty"`
	// IsDisable holds the value of the "is_disable" field.
	// 是否禁用
	IsDisable bool `json:"is_disable,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GoodsClassifyQuery when eager-loading is set.
	Edges GoodsClassifyEdges `json:"edges"`
}

// GoodsClassifyEdges holds the relations/edges for other nodes in the graph.
type GoodsClassifyEdges struct {
	// GoodsSpu holds the value of the goods_spu edge.
	GoodsSpu []*GoodsSpu `json:"goods_spu,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GoodsSpuOrErr returns the GoodsSpu value or an error if the edge
// was not loaded in eager-loading.
func (e GoodsClassifyEdges) GoodsSpuOrErr() ([]*GoodsSpu, error) {
	if e.loadedTypes[0] {
		return e.GoodsSpu, nil
	}
	return nil, &NotLoadedError{edge: "goods_spu"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodsClassify) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodsclassify.FieldIsDisable:
			values[i] = new(sql.NullBool)
		case goodsclassify.FieldID, goodsclassify.FieldPid, goodsclassify.FieldLevel, goodsclassify.FieldSort:
			values[i] = new(sql.NullInt64)
		case goodsclassify.FieldClassifyName, goodsclassify.FieldClassifyCode, goodsclassify.FieldIcon:
			values[i] = new(sql.NullString)
		case goodsclassify.FieldCreatedAt, goodsclassify.FieldUpdatedAt, goodsclassify.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodsClassify", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodsClassify fields.
func (gc *GoodsClassify) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodsclassify.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gc.ID = int(value.Int64)
		case goodsclassify.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gc.CreatedAt = value.Time
			}
		case goodsclassify.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gc.UpdatedAt = value.Time
			}
		case goodsclassify.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gc.DeletedAt = value.Time
			}
		case goodsclassify.FieldClassifyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field classify_name", values[i])
			} else if value.Valid {
				gc.ClassifyName = value.String
			}
		case goodsclassify.FieldClassifyCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field classify_code", values[i])
			} else if value.Valid {
				gc.ClassifyCode = value.String
			}
		case goodsclassify.FieldPid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				gc.Pid = int(value.Int64)
			}
		case goodsclassify.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				gc.Level = int(value.Int64)
			}
		case goodsclassify.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				gc.Sort = int(value.Int64)
			}
		case goodsclassify.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				gc.Icon = value.String
			}
		case goodsclassify.FieldIsDisable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_disable", values[i])
			} else if value.Valid {
				gc.IsDisable = value.Bool
			}
		}
	}
	return nil
}

// QueryGoodsSpu queries the "goods_spu" edge of the GoodsClassify entity.
func (gc *GoodsClassify) QueryGoodsSpu() *GoodsSpuQuery {
	return (&GoodsClassifyClient{config: gc.config}).QueryGoodsSpu(gc)
}

// Update returns a builder for updating this GoodsClassify.
// Note that you need to call GoodsClassify.Unwrap() before calling this method if this GoodsClassify
// was returned from a transaction, and the transaction was committed or rolled back.
func (gc *GoodsClassify) Update() *GoodsClassifyUpdateOne {
	return (&GoodsClassifyClient{config: gc.config}).UpdateOne(gc)
}

// Unwrap unwraps the GoodsClassify entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gc *GoodsClassify) Unwrap() *GoodsClassify {
	tx, ok := gc.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoodsClassify is not a transactional entity")
	}
	gc.config.driver = tx.drv
	return gc
}

// String implements the fmt.Stringer.
func (gc *GoodsClassify) String() string {
	var builder strings.Builder
	builder.WriteString("GoodsClassify(")
	builder.WriteString(fmt.Sprintf("id=%v", gc.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(gc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(gc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(gc.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", classify_name=")
	builder.WriteString(gc.ClassifyName)
	builder.WriteString(", classify_code=")
	builder.WriteString(gc.ClassifyCode)
	builder.WriteString(", pid=")
	builder.WriteString(fmt.Sprintf("%v", gc.Pid))
	builder.WriteString(", level=")
	builder.WriteString(fmt.Sprintf("%v", gc.Level))
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", gc.Sort))
	builder.WriteString(", icon=")
	builder.WriteString(gc.Icon)
	builder.WriteString(", is_disable=")
	builder.WriteString(fmt.Sprintf("%v", gc.IsDisable))
	builder.WriteByte(')')
	return builder.String()
}

// GoodsClassifies is a parsable slice of GoodsClassify.
type GoodsClassifies []*GoodsClassify

func (gc GoodsClassifies) config(cfg config) {
	for _i := range gc {
		gc[_i].config = cfg
	}
}
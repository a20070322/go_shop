package customer_model

import (
	"context"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/customer"
	"github.com/a20070322/shop-go/ent/customeraddress"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/pkg/model_utils"
	"time"
)

type AddressFromType struct {
	model_utils.PageOptions
	Name      string `json:"name" binding:"required"`       //姓名
	Phone     string `json:"phone" binding:"required"`      //联系方式
	Province  string `json:"province" binding:"required"`   //省
	City      string `json:"city" binding:"required"`       //市
	Area      string `json:"area" binding:"required"`       //区
	Detailed  string `json:"detailed" binding:"required"`   //详细地址
	Remark    string `json:"remark"`                        //备注
	IsDefault bool   `json:"is_default"` //是否默认
}

type AddressRepList struct {
	model_utils.PageOptions
	Data  []*ent.CustomerAddress `json:"data"`
	Total int                    `json:"total"`
}

func AddressInit(ctx context.Context) *Curd {
	art := &Curd{}
	art.db = global.Db.CustomerAddress
	art.ctx = ctx
	return art
}

type Curd struct {
	db  *ent.CustomerAddressClient
	ctx context.Context
}

func (m *Curd) List(form *model_utils.PageOptions, customerId int) (*AddressRepList, error) {
	repList := &AddressRepList{}
	model_utils.PipePagerFn(form)
	db := m.db.
		Query().
		Where(
			customeraddress.DeletedAtIsNil(),
			customeraddress.HasCustomerWith(customer.IDEQ(customerId)),
		)

	total, err2 := db.Count(m.ctx)
	if err2 != nil {
		return repList, err2
	}
	repList.Total = total
	model_utils.PipeLimitFn(db, form)

	list, err := db.All(m.ctx)
	if err != nil {
		return repList, err
	}
	repList.Data = list
	return repList, nil
}

func (m *Curd) Add(form *AddressFromType, customerId int) (*ent.CustomerAddress, error) {
	if form.IsDefault{
		m.db.Update().Where(customeraddress.HasCustomerWith(customer.ID(customerId))).SetIsDefault(false).Save(m.ctx)
	}
	db := m.db.
		Create().
		SetCustomerID(customerId).
		SetName(form.Name).
		SetPhone(form.Phone).
		SetProvince(form.Province).
		SetCity(form.City).
		SetArea(form.Area).
		SetDetailed(form.Detailed).
		SetRemark(form.Remark).
		SetIsDefault(form.IsDefault)
	rep, err := db.Save(m.ctx)
	if err != nil {
		return nil, err
	}
	return rep, nil
}

func (m *Curd) Update(uid int, form *AddressFromType, customerId int) (*ent.CustomerAddress, error) {
	if form.IsDefault{
		m.db.Update().Where(customeraddress.HasCustomerWith(customer.ID(customerId))).SetIsDefault(false).Save(m.ctx)
	}
	db, err2 := m.db.
		Query().
		Where(
			customeraddress.IDEQ(uid),
			customeraddress.DeletedAtIsNil(),
			customeraddress.HasCustomerWith(customer.IDEQ(customerId)),
		).
		First(m.ctx)
	if err2 != nil {
		return nil, err2
	}
	up := db.Update().
		SetName(form.Name).
		SetPhone(form.Phone).
		SetProvince(form.Province).
		SetCity(form.City).
		SetArea(form.Area).
		SetDetailed(form.Detailed).
		SetRemark(form.Remark).
		SetIsDefault(form.IsDefault)

	rep, err := up.Save(m.ctx)
	if err != nil {
		return nil, err
	}
	return rep, nil
}

func (m *Curd) Delete(uid int, customerId int) error {
	db, err2 := m.db.
		Query().
		Where(
			customeraddress.IDEQ(uid),
			customeraddress.DeletedAtIsNil(),
			customeraddress.HasCustomerWith(customer.IDEQ(customerId)),
		).
		First(m.ctx)
	if err2 != nil {
		return err2
	}
	_, err := db.Update().SetDeletedAt(time.Now()).Save(m.ctx)
	return err
}

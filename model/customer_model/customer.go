package customer_model

import (
	"context"
	"errors"
	"fmt"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/customer"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/pkg/utils/jwt"
)

func Init(ctx context.Context) *Customer {
	art := &Customer{}
	art.db = global.Db.Customer
	art.ctx = ctx
	return art
}

type Customer struct {
	db  *ent.CustomerClient
	ctx context.Context
}

func (m *Customer) FindOrCreateUserFormType(form *CreateUserFormType) (*ent.Customer, error) {
	if form.MiniOpenid == "" && form.UnionId == "" && form.WechatOpenid == "" {
		return nil, fmt.Errorf("用户必须存在唯一ID")
	}
	user, _ := m.db.
		Query().
		Where(
			customer.Or(
				customer.MiniOpenidEQ(form.MiniOpenid),
				customer.WechatOpenidEQ(form.WechatOpenid),
				customer.UnionIDEQ(form.UnionId),
			),
		).First(m.ctx)
	if user != nil {
		return user, nil
	}
	user2, err2 := m.CustomerCreate(form)
	if err2 != nil {
		return nil, err2
	}
	return user2, nil
}

type CreateUserFormType struct {
	MiniOpenid   string `json:"mini_openid"`
	WechatOpenid string `json:"wechat_openid"`
	UnionId      string `json:"union_id"`
	Phone        string `json:"phone"`
	Avatar       string `json:"avatar"`
}

func (m *Customer) CustomerCreate(form *CreateUserFormType) (*ent.Customer, error) {

	db := m.db.Create()
	if form.MiniOpenid != "" {
		db.SetMiniOpenid(form.MiniOpenid)
	}
	if form.WechatOpenid != "" {
		db.SetWechatOpenid(form.WechatOpenid)
	}
	if form.UnionId != "" {
		db.SetUnionID(form.UnionId)
	}
	if form.Phone != "" {
		db.SetPhone(form.Phone)
	}
	if form.Avatar != "" {
		db.SetAvatar(form.Avatar)
	}
	result, err := db.Save(m.ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type LoginUserRep struct {
	JwtData *jwt.AuthReturn `json:"jwt_data"`
	User    *ent.Customer   `json:"user"`
}

func (m *Customer) CustomerLogin(user *ent.Customer ) (*LoginUserRep, error) {
	var rep LoginUserRep
	rep.User = user
	j, jErr := jwt.GenToken(&jwt.Claims{
		UserGroup:  jwt.UserGroupCustomer,
		UserID:     user.ID,
	})
	if jErr != nil {
		return nil, errors.New("token 生成异常")
	}
	rep.JwtData = j
	return &rep, nil
}

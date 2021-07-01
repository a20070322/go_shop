package auth

import (
	"github.com/a20070322/shop-go/ent/customer"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/global/ws"
	"github.com/a20070322/shop-go/model/customer_model"
	"github.com/a20070322/shop-go/pkg/utils/jwt"
	"github.com/a20070322/shop-go/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
}

type FormMinChatLogin struct {
	JsCode string `json:"js_code" binding:"required"`
}

/**微信小程序登录*/
func (c Auth) MinChatLogin(ctx *gin.Context) {
	var form FormMinChatLogin
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	result, err2 := ws.MiniChat.Code2Session(form.JsCode)
	if err2 != nil {
		response.Fail(ctx, http.StatusInternalServerError, err2.Error(), nil)
		return
	}
	customerModel := customer_model.Init(ctx)
	user, err3 := customerModel.FindOrCreateUserFormType(&customer_model.CreateUserFormType{
		MiniOpenid: result.Openid,
		UnionId:    result.Unionid,
	})
	if err3 != nil && user != nil {
		response.Fail(ctx, http.StatusInternalServerError, err2.Error(), nil)
		return
	}
	rep, err4 := customerModel.CustomerLogin(user)
	if err4 != nil {
		response.Fail(ctx, http.StatusBadRequest, err4.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

type FormRefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// 刷新token
func (c Auth) RefreshToken(ctx *gin.Context) {
	var form FormRefreshToken
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	claims, err2 := jwt.VerifyAction(form.RefreshToken)
	if err2 != nil || claims.IsRefreshToken != true {
		response.Fail(ctx, http.StatusUnprocessableEntity, "refresh_token格式错误", nil)
		return
	}
	isCheck, err3 := jwt.CheckTokenCatch(claims, form.RefreshToken, true)
	if err3 != nil  {
		response.Fail(ctx, http.StatusBadGateway, "服务器内部错误", nil)
		return
	}
	if isCheck != true{
		response.Fail(ctx, http.StatusBadRequest, "refresh_token已失效", nil)
		return
	}
	user, err3 := global.Db.Customer.Query().Where(customer.IDEQ(claims.UserID)).First(ctx)
	if err3 != nil && user != nil {
		response.Fail(ctx, http.StatusInternalServerError, err2.Error(), nil)
		return
	}
	rep, err4 := customer_model.Init(ctx).CustomerLogin(user)
	if err4 != nil {
		response.Fail(ctx, http.StatusBadRequest, err4.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)

}

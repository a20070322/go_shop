package customer

import (
	"github.com/a20070322/shop-go/model/customer_model"
	"github.com/a20070322/shop-go/pkg/model_utils"
	"github.com/a20070322/shop-go/pkg/utils/jwt"
	"github.com/a20070322/shop-go/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Address struct {
}

func (a Address) List(ctx *gin.Context) {
	var form model_utils.PageOptions
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	uid, errJwt := jwt.GetTokenId(ctx)
	if errJwt != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errJwt.Error(), nil)
		return
	}
	rep, err := customer_model.AddressInit(ctx).List(&form, uid)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (a Address) Create(ctx *gin.Context) {
	var form customer_model.AddressFromType
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	uid, errJwt := jwt.GetTokenId(ctx)
	if errJwt != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errJwt.Error(), nil)
		return
	}
	_, err := customer_model.AddressInit(ctx).Add(&form, uid)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", true)
}

func (a Address) Update(ctx *gin.Context) {
	var form customer_model.AddressFromType
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	uid, errJwt := jwt.GetTokenId(ctx)
	if errJwt != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errJwt.Error(), nil)
		return
	}
	pUid, errP := strconv.Atoi(ctx.Param("id"))
	if errP != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errP.Error(), nil)
		return
	}
	_, err := customer_model.AddressInit(ctx).Update(pUid, &form, uid)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", true)
}

func (a Address) Delete(ctx *gin.Context) {
	uid, errJwt := jwt.GetTokenId(ctx)
	if errJwt != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errJwt.Error(), nil)
		return
	}
	pUid, errP := strconv.Atoi(ctx.Param("id"))
	if errP != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errP.Error(), nil)
		return
	}
	err := customer_model.AddressInit(ctx).Delete(pUid, uid)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", true)
}

package goods

import (
	"github.com/a20070322/shop-go/model/goods_model"
	"github.com/a20070322/shop-go/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Spu struct {
}

func (t Spu) List(ctx *gin.Context) {
	var form goods_model.SpuListFormType
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	list, err := goods_model.SpuInit(ctx).List(&form)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", list)
}

func (t Spu) GetID(ctx *gin.Context) {
	var form goods_model.SpuGetIdFormType
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	Item, err := goods_model.SpuInit(ctx).GetId(&form)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", Item)
}

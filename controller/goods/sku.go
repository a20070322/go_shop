package goods

import (
	"github.com/a20070322/shop-go/model/goods_model"
	"github.com/a20070322/shop-go/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)
type Sku struct {
}
//GetSku

func (t Sku) GetSku(ctx *gin.Context) {
	var form goods_model.SkuGetFormType
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	Item, err := goods_model.SkuInit(ctx).GetSku(&form)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", Item)
}

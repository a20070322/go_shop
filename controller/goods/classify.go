package goods

import (
	"github.com/a20070322/shop-go/model/goods_model"
	"github.com/a20070322/shop-go/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Classify struct {

}


type ClassifyConsumerGetType struct {
	ID        int   `json:"id" form:"id"`
}

func (t Classify) ConsumerGet(ctx *gin.Context) {
	var form ClassifyConsumerGetType
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	if form.ID == 0 {
		list, err := goods_model.ClassifyInit(ctx).ConsumerGet()
		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		response.Success(ctx, "ok", list)
		return
	}else{
		list, err := goods_model.ClassifyInit(ctx).ConsumerGetChildren(form.ID)
		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		response.Success(ctx, "ok", list)
		return

	}
}
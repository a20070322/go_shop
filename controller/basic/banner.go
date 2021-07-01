package basic

import (
	"github.com/a20070322/shop-go/model/basic_model"
	"github.com/a20070322/shop-go/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Banner struct {
}

// banner图列表
func (t Banner) List(ctx *gin.Context) {
	pid, errP := strconv.Atoi(ctx.Param("id"))
	if errP != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errP.Error(), nil)
		return
	}
	rep, err := basic_model.BannerInit(ctx).List(pid)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

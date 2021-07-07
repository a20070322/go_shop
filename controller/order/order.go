package order

import (
	"github.com/a20070322/shop-go/model/order"
	"github.com/a20070322/shop-go/pkg/utils/jwt"
	"github.com/a20070322/shop-go/pkg/utils/response"
	"github.com/a20070322/shop-go/queue_jobs"
	"github.com/a20070322/shop-go/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Order struct {
}

// 创建订单
func (t Order) Create(ctx *gin.Context) {
	var form order.CreateFormType
	uid, errJwt := jwt.GetTokenId(ctx)
	if errJwt != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errJwt.Error(), nil)
		return
	}
	form.CustomerId = uid
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	code, err := order.InfoInit(ctx).CreateOrder(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	//异步队列处理库存及校验订单 ， 优先级最高为10
	queue_jobs.OrderPublish(&types.QueuePayloadType{
		MethodName: "OrderProcessing",
		PayLoad:    code,
	}, 10)
	response.Success(ctx, "ok", code)
}

// 预提交订单
func (t Order) PreCreateOrder(ctx *gin.Context)  {
	var form order.PreCreateFormType
	uid, errJwt := jwt.GetTokenId(ctx)
	if errJwt != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errJwt.Error(), nil)
		return
	}
	form.CustomerId = uid
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := order.InfoInit(ctx).PreCreateOrder(&form)

	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), rep)
		return
	}
	response.Success(ctx, "ok", rep)
}

// 查看订单状态
func (t Order) StateOrder(ctx *gin.Context)  {
	var form order.StateOrderFormType
	uid, errJwt := jwt.GetTokenId(ctx)
	if errJwt != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errJwt.Error(), nil)
		return
	}
	form.CustomerId = uid
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	state, err := order.InfoInit(ctx).StateOrder(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", state)
}
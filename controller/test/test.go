package test

import (
	"fmt"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/goodsclassify"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct {
}

type ItemFormType struct {
	ImgPath []string `json:"img_path"`
	Content string   `json:"content"`
	Price   float32  `json:"price"`
	Title   string   `json:"title"`
	Img     string   `json:"img"`
	SkuId   string   `json:"sku_id"`
	ClassId int      `json:"class_id"`
}

// 商品对接nodeJs 数据导入
func (t Test) Test(ctx *gin.Context) {
	var form ItemFormType
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	sku, _ := global.Db.GoodsSku.Create().
		SetSkuName("默认").
		SetSkuCode(form.SkuId).
		SetStockNum(1000).
		SetSalesNum(0).
		//SetPrice(form.Price).
		Save(ctx)

	var bulk []*ent.GoodsSpuImgsCreate
	for _, imgPath := range form.ImgPath {
		bulk = append(
			bulk,
			global.Db.
				GoodsSpuImgs.
				Create().
				SetImgName("").
				SetImgPath(imgPath),
		)
	}
	imgs, _ := global.Db.GoodsSpuImgs.CreateBulk(bulk...).Save(ctx)
	fmt.Println(imgs)
	global.Db.GoodsSpu.Create().
		SetSpuName(form.Title).
		SetSpuCode(form.SkuId).
		SetSpuHeadImg(form.Img).
		SetSpuDesc("").
		SetSpuDetails(form.Content).
		SetIsCustomSku(false).
		SetGoodsClassifyID(form.ClassId).
		AddGoodsSku(sku).
		AddGoodsSpuImgs(imgs...).
		Save(ctx)

	response.Success(ctx, "ok", "")
}

type ItemFormType2 struct {
	Title string `json:"title"`
	Code  string `json:"code"`
	Icon  string `json:"icon"`
	Id    int    `json:"id"`
	Pid   int    `json:"pid"`
	Level int    `json:"level"`
}

// 分类对接nodejs 数据导入
func (t Test) Test2(ctx *gin.Context) {

	//var form []ItemFormType2
	//if err := ctx.ShouldBind(&form); err != nil {
	//	response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
	//	return
	//}
	//// 批量创建所有的一级分类
	//var bulk []*ent.GoodsClassifyCreate
	//for _, item := range form {
	//	bulk = append(
	//		bulk,
	//		global.Db.
	//			GoodsClassify.
	//			Create().
	//			//SetID(item.Id).
	//			SetClassifyName(item.Title).
	//			SetClassifyCode(item.Code).
	//			SetIcon(item.Icon).
	//			SetPid(item.Pid).
	//			SetLevel(item.Level).
	//			SetSort(0),
	//	)
	//}
	//global.Db.GoodsClassify.CreateBulk(bulk...).Save(ctx)


	global.Db.GoodsSku.Update().Where().SetStockNum(100).SetSalesNum(10).Save(ctx)
	response.Success(ctx, "ok", "")
}

func (t Test) Test3(ctx *gin.Context) {
	val, _ := global.Db.GoodsClassify.Query().Where(goodsclassify.LevelEQ(3)).All(ctx)
	response.Success(ctx, "ok", val)
}

func (t Test) Test4(ctx *gin.Context) {
	//oyCvV5Op3NNJWwuSZDUdXM_D67YI
	//wechat.PayInit(ctx).JsapiPrepay(&wechat.FormJsapiPrepay{
	//	OrderNumber: "1234567890",
	//	OpenID:      "oyCvV5Op3NNJWwuSZDUdXM_D67YI",
	//	Total:       10,
	//})

	//list,_ := global.Db.GoodsSpu.Query().Where(goodsspu.IDEQ(1)).WithGoodsSku(func(query *ent.GoodsSkuQuery) {
	//	query.GroupBy(goodssku.FieldSalesNum).Aggregate(ent.Sum(goodssku.FieldSalesNum)).Int(ctx)
	//}).First(ctx)
	//response.Success(ctx, "ok", list)
	//list,_ := global.Db.GoodsSku.Query().All(ctx)
	//for _,v := range list{
	//	v.Update().SetPrice(v.Price*100).Save(ctx)
	//}
	str := "12345"
	str2 := string([]rune(str)[:len(str)-2])
	str3 :=  string([]rune(str)[len(str)-2:len(str)])

	response.Success(ctx, "ok", str2+"."+str3)
}

type ItemFormType5 struct {
	Code string `json:"code"`
	Icon string `json:"icon"`
}

func (t Test) Test5(ctx *gin.Context) {
	//var form []ItemFormType5
	//if err := ctx.ShouldBind(&form); err != nil {
	//	response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
	//	return
	//}
	//for _, v := range form {
	//	global.Db.GoodsClassify.Update().Where(goodsclassify.ClassifyCodeEQ(v.Code)).SetIcon(v.Icon).Save(ctx)
	//}
	global.Db.GoodsClassify.Update().Where(goodsclassify.IconHasPrefix("data:image/png;base64,")).SetIcon("https://m.360buyimg.com/babel/jfs/t1/163408/25/5973/28466/6022424eEbf7f71ee/7193d0f9430e5515.png").Save(ctx)

	response.Success(ctx, "ok", true)
}

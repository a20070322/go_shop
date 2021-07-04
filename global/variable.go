package global

import (
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/types"
	"github.com/a20070322/shop-go/types/config"
	"github.com/go-redis/redis/v8"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger
var Db *ent.Client
var Rdb *redis.Client
var AppSetting *config.AppConfigure
var RabbitMq *types.RabbitMqType
var PayClient *core.Client

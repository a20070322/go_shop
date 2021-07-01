package basic_model

import (
	"context"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/basicbanner"
	"github.com/a20070322/shop-go/ent/basicbannerposition"
	"github.com/a20070322/shop-go/global"
)

func BannerInit(ctx context.Context) *Banner {
	art := &Banner{}
	art.db = global.Db.BasicBanner
	art.ctx = ctx
	return art
}

type Banner struct {
	db  *ent.BasicBannerClient
	ctx context.Context
}

func (m *Banner) List(positionId int) ([]*ent.BasicBanner, error) {
	return m.db.Query().Where(
		basicbanner.HasBasicBannerPositionWith(basicbannerposition.IDEQ(positionId)),
		basicbanner.StatusEQ(true),
	).
		Order(ent.Desc(basicbanner.FieldOrder)).
		WithBasicLink().
		All(m.ctx)
}

package repo

import (
	"context"
	"core-server/app/model/gormx/entity"
	"core-server/pb"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// DemoSet 注入Demo
var DemoSet = wire.NewSet(wire.Struct(new(Demo), "*"))

// Demo 示例储存
type Demo struct {
	DB *gorm.DB
}

func (a *Demo) getQueryOption(opts ...pb.DemoQueryOptions) pb.DemoQueryOptions {
	var opt pb.DemoQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Demo) Query(ctx context.Context, params *pb.DemoQueryParam, opts ...pb.DemoQueryOptions) (*pb.DemoQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetDemoDB(ctx, a.DB)
	if v := params.Code; v != "" {
		db = db.Where("code=?", v)
	}
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("name=?", v)
	}

	opt.OrderFields = append(opt.OrderFields, &pb.OrderField{Key: "id", Direction: pb.OrderDirection_OrderByDESC})
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Demos
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, err
	}
	qr := &pb.DemoQueryResult{
		PageResult: pr,
		Data:       list.ToPBDemos(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Demo) Get(ctx context.Context, id string) (*pb.Demo, error) {
	db := entity.GetDemoDB(ctx, a.DB).Where("id=?", id)
	var item entity.Demo
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToPBDemo(), nil
}

// Create 创建数据
func (a *Demo) Create(ctx context.Context, item *pb.Demo) error {
	eItem := entity.PBDemo(*item).ToDemo()
	result := entity.GetDemoDB(ctx, a.DB).Create(eItem)
	return errors.WithStack(result.Error)
}

// Update 更新数据
func (a *Demo) Update(ctx context.Context, id string, item *pb.Demo) error {
	eItem := entity.PBDemo(*item).ToDemo()
	result := entity.GetDemoDB(ctx, a.DB).Where("id=?", id).Updates(eItem)
	return errors.WithStack(result.Error)
}

// Delete 删除数据
func (a *Demo) Delete(ctx context.Context, id string) error {
	result := entity.GetDemoDB(ctx, a.DB).Where("id", id).Delete(&entity.Demo{})
	return errors.WithStack(result.Error)
}

// UpdateStatus 更新状态
func (a *Demo) UpdateStatus(ctx context.Context, id string, status int) error {
	result := entity.GetDemoDB(ctx, a.DB).Where("id=?", id).Update("status", status)
	return errors.WithStack(result.Error)
}

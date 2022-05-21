package repo

import (
	"context"
	"core-server/app/model/gormx/entity"
	"core-server/pb"
	"github.com/google/wire"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户存储
type User struct {
	DB *gorm.DB
}

func (a *User) getQueryOption(opts ...pb.UserQueryOptions) pb.UserQueryOptions {
	var opt pb.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *User) Query(ctx context.Context, params *pb.UserQueryParam, opts ...pb.UserQueryOptions) (*pb.UserQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetUserDB(ctx, a.DB)
	if v := params.UserName; v != "" {
		db = db.Where("user_name=?", v)
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}
	if v := params.UserIDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("user_name LIKE ? OR real_name LIKE ? OR phone LIKE ? OR email LIKE ?", v, v, v, v)
	}
	if v := params.PreciseSearch; v != "" {
		db = db.Where("user_name = ? OR phone = ?", v, v)
	}

	opt.OrderFields = append(opt.OrderFields, &pb.OrderField{Key: "id", Direction: pb.OrderDirection_OrderByDESC})
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.Users
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &pb.UserQueryResult{
		PageResult: pr,
		Data:       list.ToPBUsers(),
	}
	return qr, nil
}

// Get 查询指定数据
func (a *User) Get(ctx context.Context, id string) (*pb.User, error) {
	var item entity.User
	ok, err := FindOne(ctx, entity.GetUserDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToPBUser(), nil
}

// GetByUserName 使用用户名获取用户信息
func (a *User) GetByUserName(ctx context.Context, userName string) (*pb.User, error) {
	var item entity.User
	ok, err := FindOne(ctx, entity.GetUserDB(ctx, a.DB).Where("user_name=?", userName), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToPBUser(), nil
}

// Create 创建数据
func (a *User) Create(ctx context.Context, item *pb.User) error {
	eItem := entity.PBUser(*item).ToUser()
	result := entity.GetUserDB(ctx, a.DB).Create(eItem)
	return errors.WithStack(result.Error)
}

// Update 更新数据
func (a *User) Update(ctx context.Context, id string, item *pb.User) error {
	eItem := entity.PBUser(*item).ToUser()
	result := entity.GetUserDB(ctx, a.DB).Where("id=?", id).Updates(eItem)
	return errors.WithStack(result.Error)
}

// Delete 删除数据
func (a *User) Delete(ctx context.Context, id string) error {
	result := entity.GetUserDB(ctx, a.DB).Where("id=?", id).Delete(&entity.User{})
	return errors.WithStack(result.Error)
}

// UpdateStatus 更新状态
func (a *User) UpdateStatus(ctx context.Context, id string, status int) error {
	result := entity.GetUserDB(ctx, a.DB).Where("id=?", id).Update("status", status)
	return errors.WithStack(result.Error)
}

// UpdatePassword 更新密码
func (a *User) UpdatePassword(ctx context.Context, id, password string) error {
	result := entity.GetUserDB(ctx, a.DB).Where("id=?", id).Update("password", password)
	return errors.WithStack(result.Error)
}

package service

import (
	"context"
	"core-server/app/model/gormx/repo"
	"core-server/config"
	"core-server/pb"
	"core-server/util/hash"
	"core-server/util/uuid"
	"errors"

	"github.com/google/wire"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "TransModel", "UserModel"))

// User 用户管理
type User struct {
	TransModel *repo.Trans
	UserModel  *repo.User
	pb.UnimplementedUserGreeterServer
}

// Query 查询数据
func (a *User) Query(ctx context.Context, in *pb.UserQueryRequest) (*pb.UserQueryResult, error) {
	params := in.GetParams()
	opts := in.GetOptions()

	if params == nil || opts == nil {
		return nil, errors.New("请求参数不正确")
	}
	return a.UserModel.Query(ctx, params, *opts)
}

// Get 查询指定数据
func (a *User) Get(ctx context.Context, in *pb.UserGetRequest) (*pb.User, error) {
	id := in.GetID()
	if id == "" {
		return nil, errors.New("ID不能为空")
	}
	item, err := a.UserModel.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.New("未找到指定错误")
	}

	return item, nil
}

// Create 创建数据
func (a *User) Create(ctx context.Context, item *pb.User) (*pb.IDResult, error) {
	err := a.checkUserName(ctx, item)
	if err != nil {
		return nil, err
	}

	item.Password = hash.SHA1String(item.Password)
	item.ID = uuid.MustString()
	err = a.TransModel.Exec(ctx, func(ctx context.Context) error {
		return a.UserModel.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return &pb.IDResult{ID: item.ID}, nil
}

func (a *User) checkUserName(ctx context.Context, item *pb.User) error {
	if item.UserName == config.C.Root.UserName {
		return errors.New("用户名不合法")
	}

	result, err := a.UserModel.Query(ctx, &pb.UserQueryParam{
		PaginationParam: &pb.PaginationParam{OnlyCount: true},
		UserName:        item.UserName,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New("用户名已经存在")
	}
	return nil
}

// Update 更新数据
func (a *User) Update(ctx context.Context, item *pb.User) (*pb.StatusResult, error) {
	id := item.GetID()
	if id == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("ID不能为空")
	}

	oldItem, err := a.UserModel.Get(ctx, id)
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	} else if oldItem == nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("数据不存在")
	} else if oldItem.UserName != item.UserName {
		err := a.checkUserName(ctx, item)
		if err != nil {
			return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
		}
	}

	if item.Password != "" {
		item.Password = hash.SHA1String(item.Password)
	} else {
		item.Password = oldItem.Password
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt
	err = a.TransModel.Exec(ctx, func(ctx context.Context) error {
		return a.UserModel.Update(ctx, id, item)
	})
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	}

	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}

// Delete 删除数据
func (a *User) Delete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.StatusResult, error) {
	id := in.GetID()
	if id == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("ID不能为空")
	}

	oldItem, err := a.UserModel.Get(ctx, id)
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	} else if oldItem == nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("数据不存在")
	}

	err = a.TransModel.Exec(ctx, func(ctx context.Context) error {
		return a.UserModel.Delete(ctx, id)
	})
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	}

	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}

// UpdateStatus 更新状态
func (a *User) UpdateStatus(ctx context.Context, in *pb.UserUpdateStatusRequest) (*pb.StatusResult, error) {
	id := in.GetID()
	status := in.GetStatus()
	if id == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("ID不能为空")
	}
	if status != 1 && status != 2 {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("非法状态")
	}

	oldItem, err := a.UserModel.Get(ctx, id)
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	} else if oldItem == nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("数据不存在")
	}

	err = a.UserModel.UpdateStatus(ctx, id, (int)(status))
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	}
	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}

// UpdatePassword 更新密码
func (a *User) UpdatePassword(ctx context.Context, in *pb.UserUpdatePasswordRequest) (*pb.StatusResult, error) {
	id := in.GetID()
	password := in.GetPassword()
	if id == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("ID不能为空")
	}
	if password == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("密码不能为空")
	}

	oldItem, err := a.UserModel.Get(ctx, id)
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	} else if oldItem == nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("数据不存在")
	}

	err = a.UserModel.UpdatePassword(ctx, id, hash.SHA1String(password))
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	}
	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}

// Verify 验证用户
func (a *User) Verify(ctx context.Context, in *pb.UserVerifyRequest) (*pb.StatusResult, error) {
	if in.GetUserName() == "" || in.GetPassword() == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("用户名和密码不能为空")
	}

	// 先判断是否是root账号
	if in.GetUserName() == config.C.Root.UserName && in.GetPassword() == hash.MD5String(config.C.Root.Password) {
		return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
	}
	// 判断是否只需要验证root用户
	if in.GetIsRoot() {
		// 直接返回
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("管理员用户名和密码错误")
	}

	oldItem, err := a.UserModel.GetByUserName(ctx, in.GetUserName())
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	} else if oldItem == nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("用户不存在")
	}

	if hash.SHA1String(in.GetPassword()) != oldItem.Password {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("用户名密码错误")
	}

	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}

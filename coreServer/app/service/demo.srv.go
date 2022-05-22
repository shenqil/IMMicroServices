package service

import (
	"context"
	"core-server/app/model/gormx/repo"
	"core-server/pb"
	"core-server/util/uuid"
	"errors"
	"github.com/google/wire"
)

// DemoSet 注入Demo
var DemoSet = wire.NewSet(wire.Struct(new(Demo), "DemoModel"))

// Demo 示例程序
type Demo struct {
	DemoModel *repo.Demo
	pb.UnimplementedDemoGreeterServer
}

// Query 查询数据
func (a *Demo) Query(ctx context.Context, in *pb.DemoQueryRequest) (*pb.DemoQueryResult, error) {
	if in.GetOptions() == nil || in.GetOptions() == nil {
		return nil, errors.New("请求参数不正确")
	}
	return a.DemoModel.Query(ctx, in.GetParams(), *in.GetOptions())
}

// Get 查询指定数据
func (a *Demo) Get(ctx context.Context, in *pb.DemoGetRequest) (*pb.Demo, error) {
	if in.GetID() == "" {
		return nil, errors.New("ID不能为空")
	}
	item, err := a.DemoModel.Get(ctx, in.GetID())
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.New("未找到指定数据")
	}
	return item, nil
}

func (a *Demo) checkCode(ctx context.Context, code string) error {
	result, err := a.DemoModel.Query(ctx, &pb.DemoQueryParam{
		PaginationParam: &pb.PaginationParam{
			OnlyCount: true,
		},
		Code: code,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New("编号已经存在")
	}

	return nil
}

// Create 创建数据
func (a *Demo) Create(ctx context.Context, item *pb.Demo) (*pb.IDResult, error) {
	err := a.checkCode(ctx, item.Code)
	if err != nil {
		return nil, err
	}

	item.ID = uuid.MustString()
	err = a.DemoModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return &pb.IDResult{ID: item.ID}, nil
}

// Update 更新数据
func (a *Demo) Update(ctx context.Context, item *pb.Demo) (*pb.StatusResult, error) {
	if item.GetID() == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("ID不能为空")
	}

	oldItem, err := a.DemoModel.Get(ctx, item.GetID())
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	} else if oldItem == nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("数据不存在")
	} else if oldItem.Code != item.Code {
		if err := a.checkCode(ctx, item.Code); err != nil {
			return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
		}
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt
	item.Creator = oldItem.Creator

	err = a.DemoModel.Update(ctx, item.ID, item)
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	}

	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}

// Delete 删除数据
func (a *Demo) Delete(ctx context.Context, in *pb.DemoDeleteRequest) (*pb.StatusResult, error) {
	if in.GetID() == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("ID不能为空")
	}

	oldItem, err := a.DemoModel.Get(ctx, in.GetID())
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	} else if oldItem == nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("数据不存在")
	}

	err = a.DemoModel.Delete(ctx, in.GetID())
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	}
	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}

// UpdateStatus 更新状态
func (a *Demo) UpdateStatus(ctx context.Context, in *pb.DemoUpdateStatusRequest) (*pb.StatusResult, error) {
	if in.GetID() == "" {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("ID不能为空")
	}
	if in.GetStatus() != 1 && in.GetStatus() != 2 {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("非法状态")
	}
	oldItem, err := a.DemoModel.Get(ctx, in.GetID())
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	} else if oldItem == nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, errors.New("数据不存在")
	}

	err = a.DemoModel.UpdateStatus(ctx, in.GetID(), (int)(in.GetStatus()))
	if err != nil {
		return &pb.StatusResult{Status: pb.StatusText_ErrorStatus}, err
	}

	return &pb.StatusResult{Status: pb.StatusText_OKStatus}, nil
}

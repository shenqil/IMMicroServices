package entity

import (
	"context"
	"time"

	"core-server/pb"
	"core-server/util/structure"

	"gorm.io/gorm"
)

// GetDemoDB 获取demo储存
func GetDemoDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Demo))
}

// PBDemo demo对象
type PBDemo pb.Demo

// ToDemo 转换demo实体
func (a PBDemo) ToDemo() *Demo {
	item := new(Demo)
	structure.Copy(a, item)
	return item
}

// Demo demo实体
type Demo struct {
	ID        string         `gorm:"column:id;primary_key;size:36;"`
	Code      string         `gorm:"column:code;size:50;index;default:'';not null;"`
	Name      string         `gorm:"column:name;size:100;index;default:'';not null;"` // 名称
	Memo      *string        `gorm:"column:memo;size:200;"`                           // 备注
	Status    int            `gorm:"column:status;index;default:0;not null;"`         // 状态(1:启用 2:停用)
	Creator   string         `gorm:"column:creator;size:36;"`                         // 创建者
	CreatedAt time.Time      `gorm:"column:created_at;index;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;"`
}

// ToPBDemo 转换为demo对象
func (a Demo) ToPBDemo() *pb.Demo {
	item := new(pb.Demo)
	structure.Copy(a, item)
	return item
}

// Demos demo列表
type Demos []*Demo

// ToPBDemos 转换为demo对象列表
func (a Demos) ToPBDemos() []*pb.Demo {
	list := make([]*pb.Demo, len(a))
	for i, item := range a {
		list[i] = item.ToPBDemo()
	}
	return list
}

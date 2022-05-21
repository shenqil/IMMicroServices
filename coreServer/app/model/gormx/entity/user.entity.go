package entity

import (
	"context"
	"core-server/pb"
	"core-server/util/structure"
	"gorm.io/gorm"
	"time"
)

// GetUserDB 获取用户数据
func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(User))
}

// PBUser 用户对象
type PBUser pb.User

// ToUser 转换为用户实体
func (a PBUser) ToUser() *User {
	item := new(User)
	structure.Copy(a, item)
	return item
}

// User 用户实体
type User struct {
	ID        string         `gorm:"column:id;primaryKey;size:36;"`
	UserName  string         `gorm:"column:user_name;size:64;index;default:'';not null;"` // 用户名
	RealName  string         `gorm:"column:real_name;size:64;index;default:'';not null;"` // 真实姓名
	Password  string         `gorm:"column:password;size:40;default:'';not null;"`        // 密码(sha1(md5(明文))加密)
	Avatar    string         `gorm:"column:avatar;size:128;default:'';not null;"`         // 头像
	Email     *string        `gorm:"column:email;size:255;index;"`                        // 邮箱
	Phone     *string        `gorm:"column:phone;size:20;index;"`                         // 手机号
	Status    int            `gorm:"column:status;index;default:0;not null"`              // 状态(1:启用 2:停用)
	Creator   string         `gorm:"column:creator;size:36;"`                             // 创建者
	CreatedAt time.Time      `gorm:"column:created_at;index;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;"`
}

// ToPBUser 转换为用户对象
func (a User) ToPBUser() *pb.User {
	item := new(pb.User)
	structure.Copy(a, item)
	return item
}

// Users 用户实体列表
type Users []*User

// ToPBUsers 转换为用户对象列表
func (a Users) ToPBUsers() []*pb.User {
	list := make([]*pb.User, len(a))
	for i, item := range a {
		list[i] = item.ToPBUser()
	}
	return list
}

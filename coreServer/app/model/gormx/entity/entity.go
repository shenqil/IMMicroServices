package entity

import (
	"context"

	"gorm.io/gorm"

	"core-server/contextx"
)

// GetDB ...
func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := contextx.FromTrans(ctx)

	if ok && !contextx.FromNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if contextx.FromTransLock(ctx) {
				db = db.Set("gorm:query_option", "FOR UPDATE")
			}
			return db
		}
	}
	return defDB
}

// GetDBWithModel ...
func GetDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return GetDB(ctx, defDB).Model(m)
}

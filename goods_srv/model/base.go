package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	// 为什么ID要定义为int32，因为数据库的类型不同会出现错误。type为int类型基本够用，如果数据量大可以定义为bigint
	ID        int32     `gorm:"primarykey;type:int"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

package model

import "time"

// Counter 计数器模型
type Counter struct {
	Id          int32     `gorm:"column:id" json:"id"`
	Count       int32     `gorm:"column:count" json:"count"`
	GmtCreated  time.Time `gorm:"column:gmt_created" json:"gmtCreated"`
	GmtModified time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
}

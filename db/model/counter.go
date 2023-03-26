package model

import "time"

// CounterModel 计数器模型
type Counter struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Count     int32     `gorm:"column:count" json:"count"`
	gmtCreated time.Time `gorm:"column:gmt_created" json:"gmtCreated"`
	gmtModified time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
}

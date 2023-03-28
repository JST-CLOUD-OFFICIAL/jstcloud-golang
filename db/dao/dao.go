package dao

import (
	"jstcloud-golang/db"
	"jstcloud-golang/db/model"
)

const tableName = "counter"

// GetCounter 获取计数器
func (imp *CounterInterfaceImp) GetCounterById(id int32) (*model.Counter, error) {
	var err error
	var counter = new(model.Counter)

	cli := db.Get()
	err = cli.Table(tableName).Where("id = ?", id).First(counter).Error

	return counter, err
}

// GetCounter 获取最新的
func (imp *CounterInterfaceImp) GetCounter() (*model.Counter, error) {
	var err error
	var counter = new(model.Counter)

	cli := db.Get()
	err = cli.Table(tableName).Order("gmt_created DESC").Limit(1).First(counter).Error

	return counter, err
}

// UpdateCounter 更新计数器
func (imp *CounterInterfaceImp) UpdateCounter(counter *model.Counter) error {
	cli := db.Get()
	return cli.Table(tableName).Updates(counter).Error
}

// InsertCounter 插入计数器
func (imp *CounterInterfaceImp) InsertCounter(counter *model.Counter) error {
	cli := db.Get()
	return cli.Table(tableName).Create(counter).Error
}

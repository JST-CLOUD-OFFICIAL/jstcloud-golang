package dao

import (
	"jstcloud-golang/db/model"
)

// CounterInterface 计数器数据模型接口
type CounterInterface interface {
	GetCounterById(id int32) (*model.Counter, error)
	GetCounter() (*model.Counter, error)
	InsertCounter(counter *model.Counter) error
	UpdateCounter(counter *model.Counter) error
}

// CounterInterfaceImp 计数器数据模型实现
type CounterInterfaceImp struct{}

// Imp 实现实例
var Imp CounterInterface = &CounterInterfaceImp{}

package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"jstcloud-golang/db/dao"
	"jstcloud-golang/db/model"

	"gorm.io/gorm"
)

// JsonResult 返回结构
type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

// CounterHandler 计数器接口
func CounterHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	switch r.URL.Path {
	case "/counter":
		if r.Method == http.MethodGet {
			count, err := getCounter()
			if err != nil {
				res.Code = -1
				res.Message = err.Error()
			} else {
				res.Data = count
			}
		} else {
			res.Code = -1
			res.Message = fmt.Sprintf("请求方法 %s 不支持", r.Method)
		}
	case "/counter/increment":
		if r.Method == http.MethodGet {
			count, err := increaseCounter()
			if err != nil {
				res.Code = -1
				res.Message = err.Error()
			} else {
				res.Data = count
			}
		} else {
			res.Code = -1
			res.Message = fmt.Sprintf("请求方法 %s 不支持", r.Method)
		}
	case "/counter/decrement":
		if r.Method == http.MethodGet {
			count, err := decreaseCounter()
			if err != nil {
				res.Code = -1
				res.Message = err.Error()
			} else {
				res.Data = count
			}
		} else {
			res.Code = -1
			res.Message = fmt.Sprintf("请求方法 %s 不支持", r.Method)
		}
	default:
		res.Code = -1
		res.Message = "请求路径不存在"
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

// getCounter 获取计数器值
func getCounter() (int32, error) {
	counter, err := dao.Imp.GetCounterById(1)
	if err != nil {
		return 0, err
	}
	return counter.Count, nil
}

// increaseCounter 自增计数器
func increaseCounter() (int32, error) {
	currentCounter, err := dao.Imp.GetCounterById(1)
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	counter := &model.Counter{
		Id:          1,
		GmtModified: time.Now(),
	}
	if err == gorm.ErrRecordNotFound {
		counter.Count = 1
		counter.GmtCreated = time.Now()
		err = dao.Imp.InsertCounter(counter)
	} else {
		counter.Count = currentCounter.Count + 1
		err = dao.Imp.UpdateCounter(counter)
	}
	if err != nil {
		return 0, err
	}
	return counter.Count, nil
}

// decreaseCounter 自减计数器
func decreaseCounter() (int32, error) {
	currentCounter, err := dao.Imp.GetCounterById(1)
	if err != nil {
		return 0, err
	}
	counter := &model.Counter{
		Id:          1,
		GmtModified: time.Now(),
	}
	if err == gorm.ErrRecordNotFound {
		counter.Count = 1
		counter.GmtCreated = time.Now()
		err = dao.Imp.InsertCounter(counter)
	} else {
		counter.Count = currentCounter.Count - 1
		err = dao.Imp.UpdateCounter(counter)
	}
	if err != nil {
		return 0, err
	}
	return counter.Count, nil
}

package utils

import (
	"sync"
	"time"
)

var lock = &sync.Mutex{}
var baseTimeStamp, _ = time.Parse("2006/01/02 15:04:05", "2023/01/01 00:00:00")
var name = baseTimeStamp.Unix()
var preTimeStamp int64
var num int64

// GetUserId 仿照雪花算法生成id
func GetUserId() int64 {
	lock.Lock()
	TimeStamp := time.Now().Unix() - baseTimeStamp.Unix()
	if preTimeStamp < TimeStamp {
		preTimeStamp = TimeStamp
		num = 0
	} else {
		num++
	}
	id := TimeStamp<<14 | num
	lock.Unlock()
	return id

}

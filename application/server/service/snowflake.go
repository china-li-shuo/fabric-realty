package service

import (
	"fmt"
	"sync"
	"time"
)

const (
	timeLeft     = 22   // 时间戳左移的位数
	machineLeft  = 12   // 机器序号左移的位数
	maxWorkID    = 4095 // 每一毫秒内最大的工作id
	maxMachineID = 1023 // 最大的机器编号
	// 开始时间，2023-01-01 00:00:00的毫秒时间戳, 运行一段时间后如果修改该值可能导致生成重复id
	startTime = 1672502400000 //
)

type snowflake struct {
	seq       int64 // 步进号，每生成1个id自增1
	timeStamp int64 // 当前生成id的时间戳
	machineId int64 // 机器编号，范围只能在 [0:1024)
	mu        sync.Mutex
}

// NewSnowflake 输入机器编号，创建新的雪花算法
func NewSnowflake(id int64) (*snowflake, error) {
	if id < 0 || id > maxMachineID {
		return nil, fmt.Errorf("illegal machineID, the range should be between 0 and 1023")
	}

	return &snowflake{
		//timeStamp: getNowLeftTime(),
		timeStamp: time.Now().UnixNano() / 1e6,
		machineId: id,
	}, nil
}

// GetID 获取新的雪花id
func (s *snowflake) GetID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	var id int64
	curTime := time.Now().UnixNano() / 1e6

	// 新生成的雪花id时间戳必须大于等于当前时间的时间戳
	if s.timeStamp < curTime {
		s.timeStamp = curTime
		s.seq = 0
	}

	// 当前时间戳生成的雪花id已经达到最大数值后，必须进入下一毫秒才能继续生成
	if s.seq > maxWorkID {
		time.Sleep(time.Millisecond)
		s.timeStamp = time.Now().UnixNano() / 1e6
		s.seq = 0
	}
	id = ((s.timeStamp - startTime) << timeLeft) | (s.machineId << machineLeft) | s.seq
	s.seq++
	return id
}

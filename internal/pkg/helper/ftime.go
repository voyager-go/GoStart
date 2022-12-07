package helper

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 本方案为Gorm提供了一个格式化时间输入和输出的方法

type FTime struct {
	time.Time
}

func (t FTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

func (t FTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *FTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = FTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

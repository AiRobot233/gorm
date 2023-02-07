package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime struct {
	time.Time
}

// MarshalJSON 重组时间格式
func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

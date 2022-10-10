package sdk

import (
	"fmt"
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	dt.Time, err = time.Parse("2006-01-02T15:04:05", s)
	return
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", dt.Time.Format("2006-01-02T15:04:05"))), nil
}

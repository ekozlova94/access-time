package forms

import (
	"time"
)

type AccessTime struct {
	Time *time.Time `json:"time,omitempty"`
	Err  string     `json:"error,omitempty"`
}

func NewAccessTime(t *time.Time, err string) *AccessTime {
	return &AccessTime{
		Time: t,
		Err:  err,
	}
}

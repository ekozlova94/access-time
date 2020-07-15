package model

import (
	"time"
)

type AccessTime struct {
	Ip   string `json:"ip"`
	Time time.Time
}

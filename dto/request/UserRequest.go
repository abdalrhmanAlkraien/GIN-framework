package request

import (
	"time"
)

type UserRequest struct {
	Name         string
	Email        string // allow null
	Age          uint8
	Birthday     time.Time
	MemberNumber string
}

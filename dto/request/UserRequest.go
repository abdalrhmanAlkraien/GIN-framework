package request

import (
	"time"
)

type UserRequest struct {
	Name         string    `json:"name"`
	Email        string    `json:"email"` // allow null
	Age          uint8     `json:"age"`
	Birthday     time.Time `json:"birthday"`
	MemberNumber string    `json:"memberNumber"`
}

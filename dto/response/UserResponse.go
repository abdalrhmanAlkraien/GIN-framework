package response

import "time"

type UserResponse struct {
	Id           int
	Name         string
	Email        string // allow null
	Age          uint8
	Birthday     time.Time
	MemberNumber string
}

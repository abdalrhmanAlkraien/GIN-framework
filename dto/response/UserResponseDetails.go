package response

import "time"

type UserResponseDetails struct {
	Name         string
	Email        string // allow null
	Age          uint8
	Birthday     time.Time
	MemberNumber string
	todo         []TodoResponse
}

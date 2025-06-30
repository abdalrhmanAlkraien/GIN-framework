package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        *string // allow null
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString

	todo []Todo
}

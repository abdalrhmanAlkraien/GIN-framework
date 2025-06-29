package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           string `gorm:"type:uuid;primaryKey"`
	Name         string
	Email        *string // allow null
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime

	CreatedAt time.Time
	UpdatedAt time.Time
}

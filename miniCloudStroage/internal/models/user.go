package models

import "time"

type User struct {
	UserId       uint64    `gorm:"primaryKey;column:user_id"`
	UserName     string    `gorm:"column:username"`
	Email        string    `grom:"column:email"`
	PasswordHash string    `gorm:"column:password_hash"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

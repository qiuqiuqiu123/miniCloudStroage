package entity

import "time"

type FileInfo struct {
	ID          uint64 `gorm:"primary_key"`
	Name        string
	Size        int64
	Type        string
	StorageType string
	StoragePath string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	OwnerId     uint64
}

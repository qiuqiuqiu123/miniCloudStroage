package models

import "time"

type File struct {
	FileId      uint64    `grom:"primaryKey;column:file_id"`
	FileName    string    `gorm:"column:file_name"`
	FileSize    uint64    `gorm:"column:file_size"`
	FileType    string    `gorm:"column:fileType"`
	StoragePath string    `grom:"column:storage_path"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:update_at"`
}

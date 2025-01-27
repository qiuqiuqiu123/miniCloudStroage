package repository

import (
	"gorm.io/gorm"
	"miniCloudStroage/internal/models"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (fr *FileRepository) Create(file *models.File) error {
	return fr.db.Create(file).Error
}

func (fr *FileRepository) Update(file *models.File) error {
	return fr.db.Save(file).Error
}

func (fr *FileRepository) GetById(fileId uint64) (*models.File, error) {
	var file models.File
	if err := fr.db.First(&file, fileId).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (fr *)  {

}

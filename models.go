package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	UUID     string `gorm:"uniqueIndex"`
	FilePath string
}

func CreateFile(db *gorm.DB, filePath string) (string, error) {
	uuid := uuid.New().String()
	file := File{UUID: uuid, FilePath: filePath}
	if err := db.Create(&file).Error; err != nil {
		return "", err
	}
	return uuid, nil
}

func GetFilePath(db *gorm.DB, uuid string) (string, error) {
	var file File
	if err := db.Where("uuid = ?", uuid).First(&file).Error; err != nil {
		return "", err
	}
	return file.FilePath, nil
}

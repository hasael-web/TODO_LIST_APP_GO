package config

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadFileToLocalDir(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Tentukan lokasi penyimpanan file yang diunggah

	var uploadDir string

	if filepath.Ext(file.Filename) == ".pdf" {
		uploadDir = "./uploads"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return "", err
		}
	} else if filepath.Ext(file.Filename) == ".txt" {
		uploadDir = "./uploads"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return "", err
		}
	}

	// Generate nama file yang unik
	fileName := fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), filepath.Ext(file.Filename))
	dst, err := os.Create(filepath.Join(uploadDir, fileName))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Salin konten file ke file tujuan
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return filepath.Join(uploadDir, fileName), nil
}

func DeleteUploadedFile(filePath string) {
	if filePath != "" {
		_ = os.Remove(filePath)
	}
}

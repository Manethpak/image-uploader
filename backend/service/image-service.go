package service

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImageService interface {
	SaveImage(file multipart.File, image *multipart.FileHeader) (string, error)
}

type imageService struct {
}

func NewImageService() ImageService {
	return &imageService{}
}

// SaveImage accept an image and return the path
func (s *imageService) SaveImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	// write image to disk
	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath := "http://localhost:8000/public/" + filename

	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return filePath, nil
}

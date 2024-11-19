package service

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	gonanoid "github.com/matoous/go-nanoid/v2"
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

	id, err := gonanoid.New()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	fileoutput := "public/" + id + fileExt

	out, err := os.Create(fileoutput)
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

	return fileoutput, nil
}

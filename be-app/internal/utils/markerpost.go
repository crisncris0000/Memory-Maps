package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
)

func HandleImageConversion(image multipart.FileHeader) []byte {
	extension := filepath.Ext(image.Filename)
	newImageName := uuid.New().String() + extension

	fmt.Println(newImageName)

	return []byte(newImageName)
}

package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
)

func HandleMarkerPostConversion(latitudeStr, longitudeStr,
	userIDStr, descriptionStr,
	likesStr, visibilityIDStr string, imageFIle *multipart.FileHeader) (float32, float32, []byte, string, int, int, int, error) {

	latitude, err := strconv.ParseFloat(latitudeStr, 32)

	if err != nil {
		fmt.Println("Error converting to float32 for latitude")
		return -1, -1, nil, "", -1, -1, -1, err
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 32)

	if err != nil {
		fmt.Println("Error converting to float32 for longitude")
		return -1, -1, nil, "", -1, -1, -1, err
	}

	likes, err := strconv.Atoi(likesStr)

	if err != nil {
		fmt.Println("Error converting to int for likes")
		return -1, -1, nil, "", -1, -1, -1, err
	}

	visibilityID, err := strconv.Atoi(visibilityIDStr)

	if err != nil {
		fmt.Println("Error converting to int for visibility ID")
		return -1, -1, nil, "", -1, -1, -1, err
	}

	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		fmt.Println("Error converting to int for user ID")
		return -1, -1, nil, "", -1, -1, -1, err
	}

	image := HandleImageConversion(*imageFIle)

	return float32(latitude), float32(longitude), image, descriptionStr, likes, visibilityID, userID, nil
}

func HandleImageConversion(image multipart.FileHeader) []byte {
	extension := filepath.Ext(image.Filename)
	newImageName := uuid.New().String() + extension

	fmt.Println(newImageName)

	return []byte(newImageName)
}

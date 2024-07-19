package configs

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func ConnectCloudinary() (*cloudinary.Cloudinary, error) {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")

	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return nil, err
	}
	return cld, nil
}

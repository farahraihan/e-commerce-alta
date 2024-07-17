package configs

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func NewCloudinary() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary, %v", err)
	}

	return cld
}

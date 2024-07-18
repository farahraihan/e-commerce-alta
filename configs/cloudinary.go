// configs/cloudinary.go
package configs

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func ConnectCloudinary() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		return nil, err
	}
	return cld, nil
}

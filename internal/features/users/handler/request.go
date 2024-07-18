package handler

import (
	"TokoGadget/internal/features/users"
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type RegisterRequest struct {
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func ToModelUsers(r RegisterRequest) users.User {
	return users.User{
		Fullname: r.Fullname,
		Password: r.Password,
		Email:    r.Email,
	}
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	ProfilePicture string `form:"profile_picture"`
	Fullname       string `form:"fullname"`
	Email          string `form:"email"`
	Password       string `form:"password"`
	PhoneNumber    string `form:"phone_number"`
	Address        string `form:"address"`
}

func (u UpdateRequest) uploadToCloudinary(file io.Reader, filename string) (string, error) {
	// Konfigurasi Cloudinary
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", err
	}

	// Upload file ke Cloudinary
	uploadParams := uploader.UploadParams{
		Folder:   "user_pictures",
		PublicID: filename,
	}
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", err
	}

	// Ambil URL publik dari hasil unggah
	publicURL := uploadResult.URL
	return publicURL, nil
}

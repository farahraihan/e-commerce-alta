package handler

import (
	"TokoGadget/internal/features/users"
	"TokoGadget/internal/helper"
	"TokoGadget/internal/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	srv users.UserServices
}

func NewUserController(s users.UserServices) users.UserHandler {
	return &UserController{
		srv: s,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterRequest
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("register parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormatNonData(400, "input error", "error"))
		}

		err = uc.srv.Register(ToModelUsers(input))

		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak valid") {
				errCode = 400
			}
			return c.JSON(errCode, helper.ResponseFormatNonData(errCode, err.Error(), "error"))
		}

		return c.JSON(201, helper.ResponseFormatNonData(201, "success insert data", "success"))
	}
}


func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("login parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormatNonData(400, "input error", "error"))
		}

		result, token, err := uc.srv.Login(input.Email, input.Password)

		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				errCode = 400
			}
			return c.JSON(errCode, helper.ResponseFormat("error",errCode, err.Error(), nil))
		}

		return c.JSON(200, helper.ResponseFormat("success", 200, "user login successful", ToLoginReponse(result, token)))
	}
}


func (uc *UserController) Update(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID := utils.DecodeToken(token)

	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFormatNonData(http.StatusUnauthorized, "Unauthorized", "error"))
	}

	newUser := UpdateRequest{}
	if errBind := c.Bind(&newUser); errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormatNonData(http.StatusBadRequest, "Error binding data: "+errBind.Error(), "error"))
	}

	// Membaca file gambar pengguna (jika ada)
	file, err := c.FormFile("profile_picture")
	var imageURL string
	if err == nil {
		// Buka file
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormatNonData(http.StatusInternalServerError, "Gagal membuka file gambar: "+err.Error(), "error"))
		}
		defer src.Close()

		// Upload file ke Cloudinary
		imageURL, err = newUser.uploadToCloudinary(src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormatNonData(http.StatusInternalServerError, "Gagal mengunggah gambar: "+err.Error(), "error"))
		}
	}

	dataUser := users.User{
		ProfilePicture: imageURL,
		Fullname:       newUser.Fullname,
		Email:          newUser.Email,
		Password:       newUser.Password,
		PhoneNumber:    newUser.PhoneNumber,
		Address:        newUser.Address,
	}

	if errInsert := uc.srv.UpdateProfile(uint(userID), dataUser); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormatNonData(http.StatusBadRequest, "Failed to update account: "+errInsert.Error(), "error"))
		}
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormatNonData(http.StatusInternalServerError, "Failed to update account: "+errInsert.Error(), "error"))
	}

	return c.JSON(http.StatusOK, helper.ResponseFormatNonData(http.StatusOK, "Successfully updated account", "success"))
}

func (uc *UserController) GetProfile(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID := utils.DecodeToken(token)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFormatNonData(http.StatusUnauthorized, "Unauthorized", "error"))
	}

	profile, err := uc.srv.GetProfile(uint(userID))
	if err != nil {
		errMessage := "Get user profile failed: " + err.Error()
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormatNonData(http.StatusInternalServerError, errMessage, "error"))
	}
	// Mengonversi *users.User menjadi users.User menggunakan dereference (*)
	userResponse := ToGetUserResponse(*profile)

	return c.JSON(http.StatusOK, helper.ResponseFormat("success", http.StatusOK, "Get user profile successful", userResponse))
}

func (uc *UserController) Delete(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	userID := utils.DecodeToken(token)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFormatNonData(http.StatusUnauthorized,"Unauthorized", "error"))
	}

	if errDelete := uc.srv.DeleteAccount(uint(userID)); errDelete != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFormatNonData(http.StatusInternalServerError, "Failed to deleted account: "+errDelete.Error(), "Failed"))
	}

	return c.JSON(http.StatusOK, helper.ResponseFormatNonData(http.StatusOK, "Succesfully deleted account", "Success"))
}


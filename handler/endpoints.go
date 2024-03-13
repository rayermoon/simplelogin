package handler

import (
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/labstack/echo/v4"
)

func (s *Server) PostRegister(ctx echo.Context) error {
	var body generated.RegisterUserRequest

	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	createRequest := repository.CreateUserInput{
		Email:       string(*body.Email),
		FullName:    *body.FullName,
		Gender:      string(*body.Gender),
		Occupation:  *body.Occupation,
		Password:    *body.Password,
		PhoneNumber: *body.PhoneNumber,
	}
	err = s.Repository.CreateUser(ctx.Request().Context(), createRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, "ok")
}

func (s *Server) PostLogin(ctx echo.Context) error {
	var body generated.LoginRequest

	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	loginRequest := repository.LoginInput{
		PhoneNumber: *body.PhoneNumber,
		Password:    *body.Password,
	}
	token, err := s.Repository.Login(ctx.Request().Context(), loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	Message := fmt.Sprintf("User Token %s", token)
	return ctx.JSON(http.StatusOK, Message)
}

func (s *Server) GetMyProfile(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")

	data, err := utils.VerifyJWT(token)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "Not Logged in")
	}
	resp, err := s.Repository.GetUserById(ctx.Request().Context(), repository.GetUserByIdInput(data.UserID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PutMyProfile(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")

	data, err := utils.VerifyJWT(token)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "Not Logged in")
	}
	var body generated.UpdateUserProfileRequest

	err = ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	UpdateRequest := repository.UpdateUserInput{
		UserId:      data.UserID,
		FullName:    body.FullName,
		PhoneNumber: body.PhoneNumber,
	}

	err = s.Repository.UpdateUser(ctx.Request().Context(), UpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, "OK")
}

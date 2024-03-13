package handler

import (
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
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if !utils.IsValidEmail(string(*body.Email)) {
		return ctx.JSON(http.StatusBadRequest, "bad email")
	}

	createRequest := repository.CreateUserInput{
		Email:       string(*body.Email),
		FullName:    utils.StripXSS(*body.FullName),
		Gender:      string(*body.Gender),
		Occupation:  utils.StripXSS(*body.Occupation),
		Password:    *body.Password,
		PhoneNumber: *body.PhoneNumber,
	}

	err = s.Repository.CreateUser(ctx.Request().Context(), createRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "ok")
}

func (s *Server) PostLogin(ctx echo.Context) error {
	var body generated.LoginRequest

	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if !utils.IsValidEmail(string(*body.Email)) {
		return ctx.JSON(http.StatusBadRequest, "bad email")
	}

	loginRequest := repository.LoginInput{
		UserEmail: string(*body.Email),
		Password:  *body.Password,
	}
	token, err := s.Repository.Login(ctx.Request().Context(), loginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, token)
}

func (s *Server) GetMyProfile(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")

	data, err := utils.VerifyJWT(token)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "Not Logged in")
	}
	resp, err := s.Repository.GetUserById(ctx.Request().Context(), repository.GetUserByIdInput(data.UserID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
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
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	dataUser, err := s.Repository.GetUserById(ctx.Request().Context(), repository.GetUserByIdInput(data.UserID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	UpdateRequest := repository.UpdateUserInput{
		UserId: data.UserID,
	}
	if body.FullName != "" {
		UpdateRequest.FullName = body.FullName
	} else {
		UpdateRequest.FullName = dataUser.FullName
	}

	if body.PhoneNumber != "" {
		UpdateRequest.PhoneNumber = body.PhoneNumber
	} else {
		UpdateRequest.PhoneNumber = dataUser.PhoneNumber
	}

	err = s.Repository.UpdateUser(ctx.Request().Context(), UpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, "OK")
}

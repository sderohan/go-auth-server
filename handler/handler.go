package handler

import (
	"net/http"

	"github.com/sderohan/go-auth-server/pkg/config"
	"github.com/sderohan/go-auth-server/pkg/constants"
	"github.com/sderohan/go-auth-server/pkg/models"
	"github.com/sderohan/go-auth-server/pkg/utils"
	"github.com/sderohan/go-auth-server/service/jwtauth"
)

type IAuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignOut(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	jwtAuthService *jwtauth.JWTManager
}

func (auth AuthHandler) SignIn(w http.ResponseWriter, req *http.Request) {

	reqBody, err := utils.GetRequestBody(req)
	if err != nil {
		utils.WriteResponse(w, err.Error(), constants.ContentTypeKey, constants.ContentTypeValue, constants.INTERNAL_SERVER_ERROR, false)
		return
	}

	reqErrors := utils.ValidateRequestContent(reqBody)
	if reqErrors != nil {
		utils.WriteResponse(w, reqErrors, constants.ContentTypeKey, constants.ContentTypeValue, constants.BAD_REQUEST, false)
		return
	}

	// NOTE: BELOW content will be changed and for now it is used for testing purpose
	// look for the user if exists and get the user data via username and also validate the pasword present in db
	user := &models.User{
		Username: reqBody.UserName,
		Password: reqBody.Password,
		Role:     "test",
	}
	token, err := auth.jwtAuthService.GenerateToken(user)
	if err != nil {
		// handle error
	}
	tokenResponse := models.TokenResponse{
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}
	utils.WriteResponse(w, tokenResponse, constants.ContentTypeKey, constants.ContentTypeValue, constants.SUCCESS, false)
}

func (auth AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
}

func (auth AuthHandler) Validate(w http.ResponseWriter, r *http.Request) {
}

func NewAuthHandler() IAuthHandler {
	return AuthHandler{
		jwtAuthService: jwtauth.NewJWTManager(config.GetJWTConfig().SecretKey, config.GetJWTConfig().ExpiryDuration),
	}
}

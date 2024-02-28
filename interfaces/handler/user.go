package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/take-o20/layered-architecture-sample/config"
	"github.com/take-o20/layered-architecture-sample/domain"

	"github.com/take-o20/layered-architecture-sample/interfaces/response"
	"github.com/take-o20/layered-architecture-sample/usecase"
)

type UserHandler interface {
	HandleUserGet(http.ResponseWriter, *http.Request, httprouter.Params)
	HandleUserCreate(http.ResponseWriter, *http.Request, httprouter.Params)
	HandleUserList(http.ResponseWriter, *http.Request, httprouter.Params)
	HandleUserUpdate(http.ResponseWriter, *http.Request, httprouter.Params)
	HandleUserDelete(http.ResponseWriter, *http.Request, httprouter.Params)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

type userRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) HandleUserGet(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	const errMessage = "failed to got user"
	const successMessage = "got user"

	userID := params.ByName("id")

	user, err := uh.userUseCase.GetByUserID(config.DB, userID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}

	responseErr := response.UserResponse(writer, http.StatusOK, successMessage, []domain.User{*user})
	if responseErr != nil {
		response.Error(writer, http.StatusInternalServerError, responseErr, errMessage)
		return
	}
}
func (uh userHandler) HandleUserList(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	const errMessage = "failed to got users"
	const successMessage = "got users"

	users, err := uh.userUseCase.List(config.DB)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, errMessage)
		return
	}

	responseErr := response.UserResponse(writer, http.StatusOK, successMessage, users)
	if responseErr != nil {
		response.Error(writer, http.StatusInternalServerError, responseErr, errMessage)
		return
	}
}

func (uh userHandler) HandleUserCreate(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	const successMessage = "created user"
	const errMessage = "failed to create user"

	body, err := io.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err, errMessage)
		return
	}
	var requestBody userRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, errMessage)
		return
	}

	user, err := uh.userUseCase.Insert(config.DB, requestBody.Name, requestBody.Email)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, errMessage)
		return
	}

	responseErr := response.UserResponse(writer, http.StatusOK, successMessage, []domain.User{*user})
	if responseErr != nil {
		response.Error(writer, http.StatusInternalServerError, err, errMessage)
		return
	}
}

func (uh userHandler) HandleUserUpdate(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	const successMessage = "updated user"
	const errMessage = "failed to update user"

	userID := params.ByName("id")

	body, err := io.ReadAll(request.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err, errMessage)
		return
	}
	var requestBody userRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err, errMessage)
		return
	}

	user, err := uh.userUseCase.Update(config.DB, userID, requestBody.Name, requestBody.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err, errMessage)
		return
	}

	responseErr := response.UserResponse(w, http.StatusOK, successMessage, []domain.User{*user})
	if responseErr != nil {
		response.Error(w, http.StatusInternalServerError, err, errMessage)
		return
	}
}

func (uh userHandler) HandleUserDelete(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	const successMessage = "deleted user"
	const errMessage = "failed to delete user"

	userID := params.ByName("id")

	user, err := uh.userUseCase.Delete(config.DB, userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err, errMessage)
		return
	}

	responseErr := response.UserResponse(w, http.StatusOK, successMessage, []domain.User{*user})
	if responseErr != nil {
		response.Error(w, http.StatusInternalServerError, err, errMessage)
		return
	}
}

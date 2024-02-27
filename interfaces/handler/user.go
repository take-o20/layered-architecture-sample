package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/take-o20/layered-architecture-sample/config"

	"github.com/take-o20/layered-architecture-sample/interfaces/response"
	"github.com/take-o20/layered-architecture-sample/usecase"
)

type UserHandler interface {
	HandleUserGet(http.ResponseWriter, *http.Request, httprouter.Params)
	HandleUserSignup(http.ResponseWriter, *http.Request, httprouter.Params)
	HandleUserList(http.ResponseWriter, *http.Request, httprouter.Params)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) HandleUserGet(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userID := params.ByName("id")

	//usecaseレイヤを操作して、ユーザデータ取得
	user, err := uh.userUseCase.GetByUserID(config.DB, userID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}

	//レスポンスに必要な情報を詰めて返却
	response.JSON(writer, http.StatusOK, user)
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

func (uh userHandler) HandleUserSignup(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//リクエストボディを取得
	body, err := io.ReadAll(request.Body)
	if err != nil {
		response.Error(writer, http.StatusBadRequest, err, "Invalid Request Body")
		return
	}

	//リクエストボディのパース
	var requestBody userSignupRequest
	json.Unmarshal(body, &requestBody)

	//usecaseの呼び出し
	err = uh.userUseCase.Insert(config.DB, requestBody.Name, requestBody.Email)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}

	// レスポンスに必要な情報を詰めて返却
	response.JSON(writer, http.StatusOK, "")
}

type userSignupRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

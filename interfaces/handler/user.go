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
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) HandleUserGet(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Contextから認証済みのユーザIDを取得
	// ctx := request.Context()
	//  todo analyze userID from ctx
	userID := "" // context.GetUserIDFromContext(ctx)

	//usecaseレイヤを操作して、ユーザデータ取得
	user, err := uh.userUseCase.GetByUserID(config.DB, userID) // usecase.UserUsecase{}.SelectByPrimaryKey(config.DB, userID)
	if err != nil {
		response.Error(writer, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}

	//レスポンスに必要な情報を詰めて返却
	response.JSON(writer, http.StatusOK, user)
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

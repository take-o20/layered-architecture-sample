package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/take-o20/layered-architecture-sample/config"
	"github.com/take-o20/layered-architecture-sample/infrastructure/persistence"
	"github.com/take-o20/layered-architecture-sample/interfaces/handler"
	"github.com/take-o20/layered-architecture-sample/usecase"
)

func main() {
	// do not error handling
	// only use .env for development
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// config初期化
	config.Init()

	// 依存関係を注入（DI まではいきませんが一応注入っぽいことをしてる）
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	//ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v2/users/:id", userHandler.HandleUserGet)
	router.GET("/api/v2/users", userHandler.HandleUserList)
	router.POST("/api/v2/users", userHandler.HandleUserCreate)
	router.PATCH("/api/v2/users/:id", userHandler.HandleUserUpdate)

	// サーバ起動
	fmt.Println("Server Running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

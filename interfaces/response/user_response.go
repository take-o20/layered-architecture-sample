package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/take-o20/layered-architecture-sample/domain"
)

func JSON(w http.ResponseWriter, statusCode int, body interface{}) {
	json, err := json.Marshal(body)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "")
	}

	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(json))
}

func UserResponse(w http.ResponseWriter, statusCode int, message string, users []domain.User) error {
	json, err := json.Marshal(map[string]interface{}{
		"message": message,
		"users":   users,
	})

	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(json))
	return nil
}

func Error(w http.ResponseWriter, statusCode int, err error, message string) {
	log.Printf("Error: %s", err)

	json, err := json.Marshal(map[string]interface{}{
		"message": message,
	})

	if err != nil {
		log.Printf("Error: %s", err)
		w.WriteHeader(statusCode)
		return
	}

	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(json))
}

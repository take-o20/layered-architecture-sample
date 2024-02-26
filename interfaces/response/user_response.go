package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, body interface{}) {
	json, err := json.Marshal(body)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "")
	}

	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(json))
}

func Error(w http.ResponseWriter, statusCode int, err error, message string) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, err)
}

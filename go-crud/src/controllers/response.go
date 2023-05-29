package routes

import (
	"net/http"

	"go-crud/src/utils"
)

func heartbeat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	utils.WriteString(w, "{}")
}

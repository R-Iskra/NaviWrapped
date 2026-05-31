package server

import (
	"net/http"

	"github.com/R-Iskra/NaviWrapped/handlers"
)

func Start() {
	http.HandleFunc("/1/submit-listens", handlers.SubmitListens)
	http.ListenAndServe(":8080", nil)
}

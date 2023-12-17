package controller

import (
	"net/http"

	"github.com/daunsid/upload-api/pkg/util"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJson(w, 200, struct{}{})
}

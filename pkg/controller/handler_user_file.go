package controller

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/daunsid/upload-api/internal/db"
	"github.com/daunsid/upload-api/pkg/util"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) ListEntriesHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the request, for example, from the URL path or headers
	userStringID := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userStringID)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid user ID: %v", err))
		return
	}

	// Assuming you have a limit and offset as query parameters
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid limit parameter: %v", err))
		return
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid offset parameter: %v", err))
		return
	}

	// Call ListEntries function to fetch the list of files
	files, err := apiCfg.DB.ListEntries(r.Context(), db.ListEntriesParams{
		UserID: userID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching files: %v", err))
		return
	}

	// Respond with the list of files in JSON format
	util.RespondWithJson(w, http.StatusOK, files)
}

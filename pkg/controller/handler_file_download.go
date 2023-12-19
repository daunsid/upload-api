package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	//"github.com/daunsid/upload-api/internal/db"
	"github.com/daunsid/upload-api/pkg/util"
	"github.com/go-chi/chi"
	//"github.com/google/uuid"
)

func (apiCfg *ApiConfig) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	//
	// userStringID := chi.URLParam(r, "userID")
	// userID, err := uuid.Parse(userStringID)
	// if err != nil{
	// 	util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid user ID: %v", err))
	// 	return
	// }

	fileID := chi.URLParam(r, "fileID")
	if fileID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "File ID is required")
		return
	}
	// Fetch file info from the database or any other storage mechanism

	fileInfo, err := apiCfg.DB.GetFile(r.Context(), fileID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching file info: %v", err))
		return
	}
	// Fetch file content from Google Drive using the file ID
	fileContent, err := util.DownloadFromDrive(fileID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error downloading file content: %v", err))
		return
	}
	//
	// Set the appropriate headers for triggering download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.FileName))
	w.Header().Set("Content-Type", http.DetectContentType(fileContent))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(fileContent)))

	// Use ServeContent to serve the file
	http.ServeContent(w, r, fileInfo.FileName, time.Now(), bytes.NewReader(fileContent))

	//util.RespondWithJson(w, 200, )
}

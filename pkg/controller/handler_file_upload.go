package controller

import (
	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/daunsid/upload-api/internal/db"
	"github.com/go-chi/chi"
	"github.com/google/uuid"

	//"github.com/daunsid/upload-api/pkg/db"
	"github.com/daunsid/upload-api/pkg/util"
)

func (apiCfg *ApiConfig) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.

	userStringID := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userStringID)
	if err != nil {
		util.RespondWithError(w, 400, fmt.Sprintf("Could'nt create user: %v", err))
		return
	}
	//userID := r.Context().Value("user_id").(uuid.UUID)
	fileInfo := util.GoogleDrive(w, *r)
	file, err := apiCfg.DB.CreateFile(r.Context(), db.CreateFileParams{
		UserID:   userID,
		FileID:   fileInfo.FileId,
		FileName: fileInfo.FileName,
	})
	if err != nil {
		util.RespondWithError(w, 400, fmt.Sprintf("Could'nt create user: %v", err))
		return
	}

	util.RespondWithJson(w, 200, file)

}

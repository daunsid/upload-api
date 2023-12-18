package controller

import (
	"net/http"

	"github.com/daunsid/upload-api/pkg/util"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.

	util.GoogleDrive(w, r)

}

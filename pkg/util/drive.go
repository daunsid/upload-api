package util

import (
	"context"
	"fmt"
	"log"
	"net/http"

	//"os"
	//"github.com/daunsid/upload-api/pkg/util"
	"google.golang.org/api/option"

	drive "google.golang.org/api/drive/v3"
)

const (
	ServiceAccount = "service.json"
	SCOPE          = drive.DriveScope
)

func GoogleDrive(w http.ResponseWriter, r *http.Request) string {

	r.ParseMultipartForm(10 << 20)
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(ServiceAccount), option.WithScopes(SCOPE))
	if err != nil {
		log.Fatalf("Warning: Unable to create drive Client %v", err)
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error Retrieving the File: %v", err))
		return ""
	}
	defer file.Close()
	info := handler.Filename
	// Create File metadata
	f := &drive.File{Name: info}
	// Create and upload the file
	res, err := srv.Files.Create(f).
		Media(file).
		ProgressUpdater(func(now, size int64) {
			fmt.Printf("%d, %d\r", now, size)
		}).
		Do()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("New file id: %s\n", res.Id)
	return res.Id
}

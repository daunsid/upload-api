package util

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"os"
	//"github.com/daunsid/upload-api/pkg/util"
	"google.golang.org/api/option"

	drive "google.golang.org/api/drive/v3"
)

const (
	ServiceAccount = "serviceAccount.json"
	SCOPE          = drive.DriveScope
)

type FileInfo struct {
	FileId   string
	FileName string
}

func GoogleDrive(w http.ResponseWriter, r http.Request) FileInfo {

	r.ParseMultipartForm(10 << 20)

	ctx := context.Background()

	srv, err := drive.NewService(ctx, option.WithCredentialsFile(ServiceAccount), option.WithScopes(SCOPE))
	if err != nil {
		log.Fatalf("Warning: Unable to create drive Client %v", err)
	}

	file, handler, err := r.FormFile("myFile")

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error Retrieving the File: %v", err))
		return FileInfo{}
	}
	fmt.Println("debugging", handler, err)

	//defer file.Close()

	info := handler.Filename

	// // Create File metadata
	f := &drive.File{Name: info}
	// // Create and upload the file

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

	fileInfo := FileInfo{
		FileId:   res.Id,
		FileName: info,
	}
	return fileInfo
}

func DownloadFromDrive(fileID string) ([]byte, error) {
	ctx := context.Background()

	srv, err := drive.NewService(ctx, option.WithCredentialsFile(ServiceAccount))
	if err != nil {
		return nil, fmt.Errorf("unable to create Google Drive client: %v", err)

	}
	resp, err := srv.Files.Get(fileID).Download()
	if err != nil {
		return nil, fmt.Errorf("error downloading file from Google Drive: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	fileContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading file content: %v", err)
	}
	return fileContent, nil

}

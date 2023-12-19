#  Upload API

The Upload API is a service that allows users to manage file uploads, downloads, and user-related actions. The API utilizes Google Drive for storing uploaded files and provides several endpoints to interact with the system.

## Endpoints
1. UploadFile
- Endpoint: `POST /upload/{userID}`
- Description: This endpoint allows users to upload a file. The `userID` parameter identifies the user to associate the file with.
- Usage:
`curl -X POST -H "Content-Type: multipart/form-data" -F "myFile=@/path/to/file.txt" http://localhost:8080/upload/{userID}`

2. Download File
- Endpoint: `GET /download/{fileID}`
- Description: Allows users to download a file by providing the `fileID`.
- Handler Function: DownloadHandler
- Usage:
`curl -o downloaded_file.txt http://localhost:8080/download/{fileID}`

3. List User's Files

- Endpoint: `GET /files/{userID}`
- Description: Retrieves a list of files uploaded by a specific user identified by `userID`.
Handler Function: `ListEntriesHandler`
Usage:
`curl http://localhost:8080/files/{userID}`


## Google Drive API Integration

- Description: The API utilizes the Google Drive API services for storing uploaded files. The integration is implemented using the `GoogleDrive` function.

## Environment Variables
Ensure the following environment variables are set for proper API functioning:

- DATABASE_URL: The URL for the PostgreSQL database.
- GOOGLE_DRIVE_CREDENTIALS: The path to the Google Drive API credentials file.



## Running the API

- Build Docker Image:
`docker build -t upload-api .`


- Run Docker Container:
`docker run -p 8080:8080 -e DATABASE_URL="your_database_url" -e GOOGLE_DRIVE_CREDENTIALS="path/to/credentials.json" upload-api`


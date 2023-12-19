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

- Docker Compose:
`docker-compose up`


- Access the application:
`http://localhost:8090`



## Todos

A list of tasks that need to be completed or improved in this project.

### General

- [ ] Add detailed documentation for the project.
- [ ] Implement unit tests for critical components.
- [ ] Update README with comprehensive usage instructions.
- [ ] Check and update dependencies.

### Documentation

- [ ] Provide detailed API documentation.
- [ ] Document the project structure.
- [ ] Include examples and use cases in the documentation.

### Testing

- [ ] Develop unit tests for core functionality.
- [ ] Implement integration tests if applicable.
- [ ] Ensure test coverage is sufficient.




## Contributing

If you'd like to contribute to this project, please follow the [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the [License Name] - see the [LICENSE.md](LICENSE.md) file for details.
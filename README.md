# Go_Upload

This is a simple Go API server for handling file upload and retrieval. It uses SQLite for storing file metadata and Gorilla Mux for routing.

## Features

- Upload files via a POST request.
- Retrieve uploaded files using a unique identifier (UUID).

## Prerequisites

- Go installed on your machine.
- Dependencies managed using Go modules.

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/StaphoneWizzoh/Go_Upload.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Go_Upload
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Build and run the server:

    ```bash
    go run main.go
    ```

## Usage

### Uploading a File

Send a POST request to `/upload` with a form file upload:

```bash
curl -X POST -F "file=@/path/to/your/file.txt" http://localhost:8080/upload
```

### Retrieving a File

Retrieve an uploaded file using its UUID:

```bash
curl -o downloaded_file.txt http://localhost:8080/files/{uuid}
```

Replace `{uuid}` with the UUID of the file you want to retrieve.

## Contributing

Contributions are welcome! If you have any suggestions, enhancements, or bug fixes, please submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

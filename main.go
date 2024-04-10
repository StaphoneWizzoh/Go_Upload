package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    // Parse the multipart form in the request
    err := r.ParseMultipartForm(10 << 20) // 10 MB limit
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Get the file from the form
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Create a new file in the uploads directory
    filePath := filepath.Join("./uploads", handler.Filename)
    outFile, err := os.Create(filePath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer outFile.Close()

    // Copy the uploaded file content to the new file
    _, err = io.Copy(outFile, file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
    // Extract filename from URL path
    filename := filepath.Base(r.URL.Path)
    filePath := filepath.Join("./uploads", filename)

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }
    defer file.Close()

    // Stream the file to the response
    _, err = io.Copy(w, file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    // Create the uploads directory if it doesn't exist
    err := os.MkdirAll("./uploads", os.ModePerm)
    if err != nil {
        fmt.Println("Error creating uploads directory:", err)
        return
    }

    // Handle routes
    http.HandleFunc("/upload", uploadFile)
    http.HandleFunc("/files/", serveFile)

    // Start the server
    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}

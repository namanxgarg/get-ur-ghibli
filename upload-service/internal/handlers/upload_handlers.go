package handlers

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
    // In real usage, parse multipart/form-data
    fileID := fmt.Sprintf("%d", time.Now().UnixNano())
    filename := fileID + ".png"
    os.MkdirAll("./uploads", os.ModePerm)

    fullPath := filepath.Join("./uploads", filename)

    f, err := os.Create(fullPath)
    if err != nil {
        log.Println("Error creating file:", err)
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }
    defer f.Close()

    _, err = io.Copy(f, r.Body)
    if err != nil {
        log.Println("Error writing file:", err)
        http.Error(w, "Error writing file", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(fmt.Sprintf(`{"imageID":"%s"}`, fileID)))
}

package handlers

import (
    "io"
    "log"
    "net/http"
    "strings"
)

func proxyRequest(w http.ResponseWriter, r *http.Request, targetURL, method string) {
    req, err := http.NewRequest(method, targetURL, r.Body)
    if err != nil {
        log.Println("Error creating proxy request:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    req.Header = r.Header.Clone()

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error forwarding request:", err)
        http.Error(w, "Bad Gateway", http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()

    copyResponse(w, resp)
}

func proxyRequestCustomURL(w http.ResponseWriter, r *http.Request, targetURL, method string) {
    body := r.Body
    if strings.ToUpper(method) == "GET" {
        body = nil // GET requests typically don't have a body
    }

    req, err := http.NewRequest(method, targetURL, body)
    if err != nil {
        log.Println("Error creating proxy request:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    req.Header = r.Header.Clone()

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error forwarding request:", err)
        http.Error(w, "Bad Gateway", http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()

    copyResponse(w, resp)
}

func copyResponse(w http.ResponseWriter, resp *http.Response) {
    for k, v := range resp.Header {
        w.Header()[k] = v
    }
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}

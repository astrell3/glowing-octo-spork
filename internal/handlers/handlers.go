package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"index.html",
		"./index.html",
		"../index.html",
		"../../index.html",
		"../../../index.html",
	}

	var htmlContent []byte
	var err error

	for _, path := range paths {
		htmlContent, err = os.ReadFile(path)
		if err == nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(htmlContent)
			return
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(htmlContent))
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileContent := string(fileBytes)
	converted, err := service.AutoDetectAndConvert(fileContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("02-01-2006_15-04-05")
	originalExt := filepath.Ext(header.Filename)
	outputFilename := "converted_" + timestamp + originalExt

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(converted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	response := fmt.Sprintf(
		"Original: %s\nOutput: %s\n\nOriginal text: %s\nConverted: %s",
		header.Filename,
		outputFilename,
		fileContent,
		converted,
	)
	w.Write([]byte(response))
}

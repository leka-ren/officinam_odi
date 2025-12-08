package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func StartPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "method GET only available", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(res, req, "index.html")
}

func UploadForm(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "method POST only available", http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseMultipartForm(5 << 20)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	file, header, err := req.FormFile("myFile")

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	defer file.Close()

	dataFromFile, err := io.ReadAll(file)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	convertedData := service.Сonversion(string(dataFromFile))

	fileName := time.Now().UTC().String() + filepath.Ext(header.Filename)

	_, err = os.Create(fileName)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	_, err = res.Write([]byte(convertedData))

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

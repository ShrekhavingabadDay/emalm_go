package controller

import (
	"emalm/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	uuid, _ := NewUUID()

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	dst, err := os.Create(fmt.Sprintf("./uploads/images/%s%s", uuid, filepath.Ext(fileHeader.Filename)))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.ParseForm()

	img := model.Image{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Uuid:        uuid,
		// TODO: get authentication done and get user by session
		User: 1,
	}

	err = model.UploadImage(img)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GetAllImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	images, err := model.GetAllImages()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(images)
	}
}

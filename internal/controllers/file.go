package controllers

import (
	"net/http"

	"github.com/adityapp/sheds/internal/services"
	"github.com/sirupsen/logrus"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		logrus.Error(err)
		return
	}

	dir := r.FormValue("directory")

	err = services.Store(file, fileHeader, dir)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")

	err := services.Delete(path)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Move(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	targetDir := r.FormValue("target_dir")

	err := services.Move(path, targetDir)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Copy(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	targetDir := r.FormValue("target_dir")

	err := services.Copy(path, targetDir)
	if err != nil {
		logrus.Error(err)
		return
	}
}

package services

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/adityapp/sheds/configs"
	"github.com/sirupsen/logrus"
)

func Store(file multipart.File, fileHeader *multipart.FileHeader, dir string) error {
	trimmedDir := strings.Trim(dir, "/")
	path := filepath.Join(configs.Get().WorkDir, trimmedDir, fileHeader.Filename)

	destination, err := os.Create(path)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, file)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func Delete(filePath string) error {
	err := os.RemoveAll(filePath)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func Move(filePath, targetDir string) error {
	splitPath := strings.Split(filePath, "/")
	targetPath := filepath.Join(targetDir, splitPath[len(splitPath)-1])

	err := os.Rename(filePath, targetPath)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func Copy(filePath, targetDir string) error {
	splitPath := strings.Split(filePath, "/")
	targetPath := filepath.Join(targetDir, splitPath[len(splitPath)-1])

	destination, err := os.Create(targetPath)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer destination.Close()

	sourceFile, err := os.Open(filePath)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer sourceFile.Close()

	_, err = io.Copy(destination, sourceFile)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

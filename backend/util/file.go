package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func FileExists(filePath string) bool {
	_, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Unzip(file, dst string) error {
	zipFile, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	if !FileExists(dst) {
		err := os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for _, f := range zipFile.File {
		dstPath := filepath.Join(dst, f.Name)
		if f.FileInfo().IsDir() {
			err := os.MkdirAll(dstPath, f.Mode())
			if err != nil {
				return err
			}
			fmt.Println("unzipped", dstPath)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()

		fmt.Println("unzipped", dstPath)
	}
	return nil
}

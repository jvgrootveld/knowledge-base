package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"path"
)

type File struct {
	Name string `json:"name"`
}

type Directory struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Directories []Directory `json:"directories"`
	Files       []File      `json:"files"`
}

func InitDirectoryApi(r *mux.Router) {
	r.HandleFunc("/directory", listDirectoryHandler).Methods("GET")
}

func listDirectory(currentPath string) Directory {
	infoList, _ := ioutil.ReadDir(currentPath)

	directories := []Directory{}
	files := []File{}

	for _, info := range infoList {
		if info.IsDir() {
			fullPath := path.Join(currentPath, info.Name())
			directories = append(directories, listDirectory(fullPath))
		} else {
			files = append(files, File{Name: info.Name()})
		}
	}

	return Directory{Name: path.Base(currentPath), Path: currentPath, Directories: directories, Files: files}
}

func listDirectoryHandler(w http.ResponseWriter, _ *http.Request) {
	directoryPath := "testdir"

	directory := listDirectory(directoryPath)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(directory)
}

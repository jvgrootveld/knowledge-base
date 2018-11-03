package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jvgrootveld/knowledge-base/internal/config"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"net/http"
)

func InitFileApi(r *mux.Router) {
	r.HandleFunc("/file/{path:.+}", readFileHandler).Methods("GET")
	r.HandleFunc("/file/{path:.+}", writeFileHandler).Methods("POST")
	r.HandleFunc("/html/{path:.+}", htmlHandler).Methods("GET")
}

// Retrieves and returns file at requested {path}
// Path: `GET "../file/{path}"
func readFileHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	filePath := config.FileDirectory + "/" + path

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error while parsing markdown %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(fileContent)
}

// Writes body bytes into {path}
// Path: POST "../file/{path}"
//
// Example:
/*
curl -X POST \
-H "Content-Type:text/plain" \
--data-binary "$(< $GOPATH/src/github.com/jvgrootveld/knowledge-base/testdir/testfile.md)" \
http://localhost:8000/kbapi/file/testwrite.md
*/
func writeFileHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	filePath := config.FileDirectory + "/" + path
	fileData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error while getting request body data %v\n", err)
		return
	}

	err = ioutil.WriteFile(filePath, fileData, 0644)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error while writing file (%v) %v\n", path, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Retrieves markdown at requested path and returns it as parsed HTML
// Path GET "../html/{path}"
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	filePath := config.FileDirectory + "/" + path

	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error while parsing markdown %v\n", err)
		return
	}

	output := blackfriday.Run(fileContent, blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		CSS:   "https://notes.peter-baumgartner.net/archive/testfiles/2017-09-23-file-attachments.files/themeXXX.css",
		Flags: blackfriday.CompletePage | blackfriday.TOC,
	})))

	w.WriteHeader(http.StatusOK)

	w.Write(output)
}

package main

import (
	"fmt"
	"github.com/jvgrootveld/knowledge-base/internal"
	"github.com/jvgrootveld/knowledge-base/internal/config"
	"log"
	"net/http"
)

const defaultPort = "8000"

func main() {
	config.FileDirectory = "testdir"

	fmt.Printf("Started Knowledge Base on port (%v) with file directory (%v) \n", defaultPort, config.FileDirectory)

	log.Fatal(http.ListenAndServe(":"+defaultPort, internal.NewRouter()))
}

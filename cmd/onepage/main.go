package main

import (
	"github.com/maveonair/onepage/internal/server"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	pageFilePath := "./page.md"
	if _, err := os.Stat(pageFilePath); os.IsNotExist(err) {
		file, err := os.Create(pageFilePath)
		if err != nil {
			log.Fatalf("Could not create file %s", pageFilePath)
		}
		defer file.Close()
	}

	listeningAddr := "0.0.0.0:8080"
	log.Infof("Listening on %s", listeningAddr)

	srv, err := server.NewServer(pageFilePath)
	if err != nil {
		log.Fatalf("Could not create server %s", err)
	}

	log.Fatal(http.ListenAndServe(listeningAddr, srv))
}

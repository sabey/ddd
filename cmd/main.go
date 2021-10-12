package main

import (
	"log"
	"time"

	net_http "net/http"

	"github.com/sabey/ddd/http"
	"github.com/sabey/ddd/repo"
)

func main() {
	r, err := repo.NewRepository(
		repo.RepositoryOpts{
			Addr:     "192.168.2.214:5432",
			User:     "postgres",
			Password: "postgres",
			Database: "postgres",
			Drop:     true,
		},
	)
	if err != nil {
		log.Fatalf("failed to build postgres repo: %s\n", err)
	}
	defer r.Close()

	s := &net_http.Server{
		Addr:           ":8080",
		Handler:        http.NewHTTPService(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("listening on 0.0.0.0:8080")

	log.Fatal(s.ListenAndServe())
}

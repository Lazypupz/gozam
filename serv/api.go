package api

import (
	"log"
	"net/http"
)

type APISever struct {
	listenAddr string
}

func NewServer(listenAddr string) *APISever {
	return &APISever{
		listenAddr: listenAddr,
	}
}

func (s *APISever) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte(userID))
	})

	server := http.Server{
		Addr:    s.listenAddr,
		Handler: router,
	}

	log.Printf("Server running on addr %s", s.listenAddr)

	return server.ListenAndServe()
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(`{"message": "hello world"}`))

	if err != nil {
		fmt.Println("error al devolver la respuesta a la petición")
		log.Fatal(err)
	}
}
func main() {
	s := &Server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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
	greetingJson := GetHelloJSON()
	_, err := w.Write([]byte(greetingJson))

	if err != nil {
		fmt.Println("error al devolver la respuesta a la petición")
		log.Fatal(err)
	}
}
func main() {
	s := &Server{}
	http.Handle("/", s)
	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

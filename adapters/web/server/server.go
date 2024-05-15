package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alzo91/go-hexagonal/adapters/web/handler"
	"github.com/alzo91/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	service application.ProductServiceInterface
}

func NewWebServer(service application.ProductServiceInterface) *WebServer {
	return &WebServer{
		service: service,
	}
}

func (ws *WebServer) Serve() {
	
	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	// handler.MakeProductHandlers(r, n, ws.service)
	handler.MakeProductHandlers(r, n, ws.service)
	http.Handle("/", r)
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 		10 * time.Second,
		Addr: ":9000",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

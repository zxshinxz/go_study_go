package service

import (
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	sgm "github.com/the-gigi/delinkcious/pkg/post_manager"
	"log"
	"net/http"
)

var (
	// return when an expected path variable is missing.
	BadRoutingError = errors.New("inconsistent mapping between route and handler")
)

func Run() {
	store, err := sgm.NewEtcdPostStore("localhost", 5432, "postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	svc, err := sgm.NewPostManager(store)
	if err != nil {
		log.Fatal(err)
	}

	postHandler := httptransport.NewServer(
		makePostEndpoint(svc),
		decodePostRequest,
		encodeResponse,
	)

	getPostHandler := httptransport.NewServer(
		makeGetPostsEndpoint(svc),
		decodeGetPostsRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/post").Handler(postHandler)
	r.Methods("GET").Path("/posts").Handler(getPostHandler)

	log.Println("Listening on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", r))
}

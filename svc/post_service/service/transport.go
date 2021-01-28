package service

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	om "github.com/the-gigi/delinkcious/pkg/object_model"
	"net/http"
)

type postRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type postResponse struct {
	Err string `json:"err"`
}

type getPostsRequest struct {
}

func decodePostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request postRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetPostsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getPostsRequest
	//err := json.NewDecoder(r.Body).Decode(&request)
	//if err != nil {
	//	return nil, err
	//}
	return request, nil
}

type getPostsResponse struct {
	Posts string `json:"posts"`
	Err   string `json:"err"`
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makePostEndpoint(svc om.PostManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(postRequest)
		err := svc.Post(req.Title, req.Content)
		res := postResponse{}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeGetPostsEndpoint(svc om.PostManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		postsMap, err := svc.GetPosts()
		res := getPostsResponse{Posts: postsMap}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

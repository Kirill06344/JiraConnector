package utils

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Envelope map[string]interface{}

type PageParams struct {
	Limit  int
	Page   int
	Search string
}

func ReadIdParam(r *http.Request) (uint, error) {
	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return uint(id), nil
}

func GetKeyQuery(r *http.Request) (string, error) {
	query := r.URL.Query()
	key := query.Get("key")
	if key == "" {
		return "", errors.New("please, write key parameter to download project")
	}
	return key, nil
}

func GetQueryParams(r *http.Request) (*PageParams, error) {
	query := r.URL.Query()
	params := &PageParams{}

	limitParam := query.Get("limit")
	if limitParam != "" {
		limit, err := strconv.Atoi(limitParam)
		if err != nil || limit < 1 {
			return nil, errors.New("incorrect value for limit parameter")
		}
		params.Limit = limit
	} else {
		params.Limit = 20
	}

	pageParam := query.Get("page")
	if pageParam != "" {
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			return nil, errors.New("incorrect value for page parameter")
		}
		params.Page = page
	} else {
		params.Page = 1
	}

	params.Search = query.Get("search")
	return params, nil
}

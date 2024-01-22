package utils

import (
	"errors"
	"net/http"
	"strconv"
)

type PageParams struct {
	Limit  int
	Page   int
	Search string
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

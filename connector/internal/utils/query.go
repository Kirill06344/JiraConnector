package utils

import (
	"github.com/stewie/internal/pb"
)

type PageParams struct {
	Limit  int
	Page   int
	Search string
}

func GetQueryParams(r *pb.AllProjectsRequest) (*PageParams, error) {
	params := &PageParams{}

	if r.Limit > 0 {
		params.Limit = int(r.Limit)
	} else {
		params.Limit = 20
	}

	if r.Page > 0 {
		params.Page = int(r.Page)
	} else {
		params.Page = 1
	}

	params.Search = r.Search
	return params, nil
}

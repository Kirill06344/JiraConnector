package utils

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Envelope map[string]interface{}

func ReadIdParam(r *http.Request) (int64, error) {
	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return int64(id), nil
}

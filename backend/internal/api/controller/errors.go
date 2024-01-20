package controller

import (
	utils "backend/internal/utils"
	"fmt"
	"net/http"
)

func logError(r *http.Request, err error) {
	utils.Logger.Warnln(err)
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := utils.Envelope{"errors": message}

	err := utils.WriteJSON(w, status, env, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)

	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

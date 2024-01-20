package controller

import (
	utils "backend/internal/utils"
	"fmt"
	"net/http"
)

func logError(r *http.Request, err error) {
	utils.Logger.Warnln(err)
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}, name string) {
	env := utils.Envelope{
		"message": message,
		"name":    name,
		"status":  false,
	}

	err := utils.WriteJSON(w, status, env, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error, name string) {
	logError(r, err)

	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message, name)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, name string) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message, name)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message, "")
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error, name string) {
	errorResponse(w, r, http.StatusBadRequest, err.Error(), name)
}

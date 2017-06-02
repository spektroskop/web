package web

import (
	"fmt"
	"net/http"
)

func makeError(w http.ResponseWriter, code int, msg string, arg ...interface{}) {
	http.Error(w, fmt.Sprintf(msg, arg...), code)
}

func BadRequest(w http.ResponseWriter, msg string, arg ...interface{}) {
	makeError(w, http.StatusBadRequest, msg, arg...)
}

func NotFound(w http.ResponseWriter, msg string, arg ...interface{}) {
	makeError(w, http.StatusNotFound, msg, arg...)
}

func InternalServerError(w http.ResponseWriter, msg string, arg ...interface{}) {
	makeError(w, http.StatusInternalServerError, msg, arg...)
}

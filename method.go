package web

import (
	"fmt"
	"net/http"
)

type Methods []string

func HandleMethods(methods ...string) Methods {
	return methods
}

func Post(handler http.Handler) http.Handler {
	return HandleMethods("GET").With(handler)
}

func Get(handler http.Handler) http.Handler {
	return HandleMethods("GET").With(handler)
}

func Put(handler http.Handler) http.Handler {
	return HandleMethods("GET").With(handler)
}

func Delete(handler http.Handler) http.Handler {
	return HandleMethods("GET").With(handler)
}

func (methods Methods) With(handlers ...http.Handler) http.Handler {
	if len(handlers) != len(methods) {
		panic(fmt.Sprintf("incorrect number of handlers (%d) to methods (%d)",
			len(handlers), len(methods),
		))
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for i, method := range methods {
			if method == r.Method {
				handlers[i].ServeHTTP(w, r)
				return
			}
		}

		http.Error(w,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed,
		)
	})
}

package web

import (
	"net/http"
)

type Params []string

func Query(params ...string) Params {
	return params
}

func (params Params) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, param := range params {
			if v := r.URL.Query().Get(param); v == "" {
				http.Error(w,
					http.StatusText(http.StatusBadRequest),
					http.StatusBadRequest,
				)

				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

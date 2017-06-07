package web

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(handlers ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		for i := len(handlers) - 1; i >= 0; i-- {
			handler = handlers[i](handler)
		}

		return handler
	}
}

func WithChain(handler http.Handler, chain ...Middleware) http.Handler {
	return Chain(chain...)(handler)
}

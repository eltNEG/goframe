package utils

import (
	"net/http"
)

type Http struct {
	mux *http.ServeMux
}

func (h *Http) Get(path string, handlers http.HandlerFunc) *Http {
	h.mux.HandleFunc("GET "+path, handlers)
	return h
}

func (h *Http) Post(path string, handlers http.HandlerFunc) *Http {
	h.mux.HandleFunc("POST "+path, handlers)
	return h
}

func (h *Http) Put(path string, handlers http.HandlerFunc) *Http {
	h.mux.HandleFunc("PUT "+path, handlers)
	return h
}

func (h *Http) Delete(path string, handlers http.HandlerFunc) *Http {
	h.mux.HandleFunc("DELETE "+path, handlers)
	return h
}

func (h *Http) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, h.mux)
}

func NewHttp() *Http {
	return &Http{mux: http.NewServeMux()}
}

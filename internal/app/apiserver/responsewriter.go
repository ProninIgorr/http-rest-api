package apiserver

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statucCode int) {
	w.code = statucCode
	w.ResponseWriter.WriteHeader(statucCode)
}

package common

import (
	"log"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	w.length = len(b)
	return w.ResponseWriter.Write(b)
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		start := time.Now()
		writer := statusWriter{w, 0, 0}
		inner.ServeHTTP(&writer, request)
		end := time.Now()
		latency := end.Sub(start)
		statusCode := writer.status
		length := writer.length
		if request.URL.RawQuery != "" {
			log.Printf("%v %s %s \"%s %s%s%s %s\" %d %d \"%s\" %v", end.Format("2006/01/02 15:04:05"), request.Host, request.RemoteAddr, request.Method, request.URL.Path, "?", request.URL.RawQuery, request.Proto, statusCode, length, request.Header.Get("User-Agent"), latency)
		} else {
			log.Printf("%v %s %s \"%s %s %s\" %d %d \"%s\" %v", end.Format("2006/01/02 15:04:05"), request.Host, request.RemoteAddr, request.Method, request.URL.Path, request.Proto, statusCode, length, request.Header.Get("User-Agent"), latency)
		}
	})
}

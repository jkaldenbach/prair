package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger : log incoming requests
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s\t%s",
			start.Format("2006-01-02 3:04:05PM"),
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

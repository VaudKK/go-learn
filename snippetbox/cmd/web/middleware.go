package main

import (
	"fmt"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w,r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s",r.RemoteAddr,r.Proto,r.Method,r.URL)

		next.ServeHTTP(w,r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create a deferred func that will always be called in case of a panic
		defer func ()  {
			if err := recover(); err != nil {
				// set a connection close header on the response
				w.Header().Set("Connection", "close")
				app.serverError(w,fmt.Errorf("%s",err))
			}
		}()

		next.ServeHTTP(w,r)
	})
}
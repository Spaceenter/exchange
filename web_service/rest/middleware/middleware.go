package middleware

import (
	"log"
	"net/http"
)

//LoggingMiddleware logging the request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

//Authenticate is an middleware to check user auth status
func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO this is just an example.
		//Get the cookie or something and check it
		cookie, err := r.Cookie("session")
		if err != nil || cookie == nil {
			// TODO if an err, then redirect
			// http.Redirect(w, r, "/", 401)
		}
		//If the auth check passes, then handle continue down the chain
		h.ServeHTTP(w, r)
	})
}

//SetContentTypeText this only exists to demonstrate how we can chain middlewares
func SetContentTypeText(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		h.ServeHTTP(w, r)
	})
}

//PostProcess //TODO this should be called after controller returnes
// func PostProcess(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/plain")
// 		h.ServeHTTP(w, r)
// 	})
// }

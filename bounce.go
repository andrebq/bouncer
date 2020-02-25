package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

const (
	BounceCheckPath = "/bouncer/check"
)

// Bounce the request if the required cookie is not set
// or let it through if the cookie is set
func Bounce(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("bouncer")
	if err == http.ErrNoCookie {
		// missing cookie, redirect the user to check access
		http.Redirect(w, req, BounceCheckPath, http.StatusSeeOther)
		return
	}
	if err != nil {
		log.Error().Err(err).Msg("Unable to extract cookie from request")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	log.Info().Interface("cookie", c)

	proxy.ServeHTTP(w, req)
}

// CheckAccess allows bouncer to selectively allow others to access the protected resource
func CheckAccess(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		validateAccess(w, req)
		return
	case "GET":
		renderAuthentication(w, req)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	// log.Error().Msg("Will check access. but for now, here is a 403 for you")
	// http.Error(w, "Heute nicht", http.StatusForbidden)
}

func validateAccess(w http.ResponseWriter, req *http.Request) {
}

func renderAuthentication(w http.ResponseWriter, req *http.Request) {
}

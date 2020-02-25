package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	proxy  http.Handler
	server *http.Server
)

func init() {
	var proxyURL string
	var readTimeout time.Duration
	var writeTimeout time.Duration
	var addr string

	flag.StringVar(&proxyURL, "target", "http://localhost:8081", "URL to proxy requests to")
	flag.StringVar(&addr, "addr", "127.0.0.1:8080", "Address to bind for incoming requests")
	flag.DurationVar(&readTimeout, "read-timeout", time.Minute, "How long to wait for the first read")
	flag.DurationVar(&writeTimeout, "write-timeout", time.Minute, "How long to wait for the first write")

	flag.Parse()

	targetURL, err := url.Parse(proxyURL)
	if err != nil {
		panic(err.Error())
	}

	proxy = httputil.NewSingleHostReverseProxy(targetURL)

	server = &http.Server{
		Addr:           addr,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 * 1024 * 1024 * 1024,
	}

}

func main() {
	http.HandleFunc(BounceCheckPath, CheckAccess)
	http.HandleFunc("/", Bounce)

	log.Info().Str("Addr", server.Addr).Msg("Starting server")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Unexpected server exit")
	} else {
		log.Info().Msg("Clean server exit")
	}
}

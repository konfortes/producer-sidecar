package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/konfortes/go-server-utils/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func createHTTPHandlers() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/produceAsync", produceAsync)

	addr := fmt.Sprintf("%s:%s", host, *port)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		log.Printf("Listeneing on " + addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	server.GracefulShutdown(srv)
}

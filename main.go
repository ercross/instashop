package main

import (
	"context"
	"errors"
	"instashop/api"
	"instashop/db"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	run(ctx)
}

func run(ctx context.Context) {
	repo, err := db.NewDB(os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}
	srv := api.NewServer(repo)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort("127.0.0.1", "15001"),
		Handler: srv,
	}
	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
}

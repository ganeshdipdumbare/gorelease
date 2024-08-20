package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/internal/app"
)

const (
	nilArgErr   = "nil %v not allowed"
	emptyArgErr = "empty %v not allowed"
)

// RestApi defines methods to handle rest server
type RestApi interface {
	StartServer()
	GracefulStopServer()
}

type apiDetails struct {
	app    app.CodePlayer
	server *http.Server
}

// NewApi creates new api instance, otherwise returns error
func NewApi(a app.CodePlayer, port string) (RestApi, error) {
	if a == nil {
		return nil, fmt.Errorf(nilArgErr, "app")
	}

	if port == "" {
		return nil, fmt.Errorf(emptyArgErr, "port")
	}

	api := &apiDetails{
		app: a,
	}

	router := api.setupRouter()
	api.server = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: router,
	}

	return api, nil
}

// StartServer starts rest server and wait for kill signal to stop it gracefully
// otherwise returns error
func (a *apiDetails) StartServer() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	a.GracefulStopServer()
}

// GracefulStopServer gracefully stops the rest server
func (a *apiDetails) GracefulStopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

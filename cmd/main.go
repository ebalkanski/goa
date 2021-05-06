package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ebalkanski/goa/gen/play"
	"github.com/ebalkanski/goa/internal/service"
)

func main() {
	// Load configuration from environment variables
	var cfg config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("error loading configuration")
	}

	// Init Mongo client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.URI))
	if err != nil {
		log.Fatal(err)
	}

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[playapi] ", log.Ltime)
	}

	// Initialize the services.
	var (
		playSvc play.Service
	)
	{
		playSvc = service.NewPlay(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		playEndpoints *play.Endpoints
	)
	{
		playEndpoints = play.NewEndpoints(playSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel = context.WithCancel(context.Background())

	// Start server
	handleHTTPServer(ctx, "localhost:8080", playEndpoints, &wg, errc, logger, false)

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

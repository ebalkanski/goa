package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/kelseyhightower/envconfig"

	goauser "github.com/ebalkanski/goa/gen/user"
	storage "github.com/ebalkanski/goa/internal/clients/storage/mongo"
	"github.com/ebalkanski/goa/internal/clients/storage/mongo_repo"
	"github.com/ebalkanski/goa/internal/service/user"
)

func main() {
	// Init context
	ctx, cancel := context.WithCancel(context.Background())

	// Setup logger
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[goa] ", log.Ltime)
	}

	// Load configuration from environment variables
	var cfg config
	if err := envconfig.Process("", &cfg); err != nil {
		logger.Fatal("error loading configuration")
	}

	// Init Mongo client
	mongoDB := storage.NewMongo(logger, ctx, cfg.Mongo.URI)
	defer mongoDB.Disconnect(ctx)

	// Create repositories
	userRepo := mongo_repo.NewUserRepo(logger, mongoDB, cfg.Mongo.DB)

	// Initialize the services.
	var (
		userSvc goauser.Service
	)
	{
		userSvc = user.NewUser(logger, userRepo)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		userEndpoints *goauser.Endpoints
	)
	{
		userEndpoints = goauser.NewEndpoints(userSvc)
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

	// Start server
	handleHTTPServer(ctx, "localhost:8080", userEndpoints, &wg, errc, logger, false)

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

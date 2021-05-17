package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
	goa "goa.design/goa/v3/pkg"

	openapisvr "github.com/ebalkanski/goa/gen/http/openapi/server"
	usersvr "github.com/ebalkanski/goa/gen/http/user/server"
	goauser "github.com/ebalkanski/goa/gen/user"
	"github.com/ebalkanski/goa/internal/clients/storage"
	"github.com/ebalkanski/goa/internal/config"
	"github.com/ebalkanski/goa/internal/service/goa_errors"
	"github.com/ebalkanski/goa/internal/service/user"
)

func main() {
	// Init context
	ctx := context.Background()

	// Setup logger
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[goa] ", log.Ltime)
	}

	// Load configuration from environment variables
	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		logger.Fatal("error loading configuration")
	}

	// Init Mongo client
	mongoDB := storage.NewMongo(logger, ctx, cfg.Mongo.URI)
	defer mongoDB.Disconnect(ctx)

	// Create storage clients
	collection := mongoDB.Database(cfg.Mongo.DB).Collection("users")
	userStorage := storage.NewUser(logger, collection)

	// Initialize the services.
	var (
		userSvc goauser.Service
	)
	{
		userSvc = user.NewUser(userStorage)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		userEndpoints *goauser.Endpoints
	)
	{
		userEndpoints = goauser.NewEndpoints(userSvc)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	mux := goahttp.NewMuxer()

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		userServer *usersvr.Server
	)
	{
		eh := errorHandler(logger)
		userServer = usersvr.New(userEndpoints, mux, dec, enc, eh, customErrorResponse)
	}

	// Configure the mux.
	usersvr.Mount(mux, userServer)
	openapisvr.Mount(mux)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	for _, m := range userServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	logger.Println("starting server at localhost:8080")

	if err := start(srv, 10*time.Second, nil); err != nil {
		logger.Fatal(err)
	}

	logger.Println("bye bye")
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}

// customErrorResponse converts err into a global goa error
func customErrorResponse(err error) goahttp.Statuser {
	if serr, ok := err.(*goa.ServiceError); ok {
		return goa_errors.NewBadRequestError(serr)
	}

	if gerr, ok := err.(*goa_errors.Error); ok {
		return gerr
	}

	return goa_errors.NewInternalServerError(err)
}

type shutdownFunc func(ctx context.Context) error

// start starts a given HTTP server and gracefully stops the server
// upon receiving a stop signal. The server will wait for the active
// connections to be closed for {timeout} period of time.
//
// Optional functions can be passed as arguments which shall be called
// after receiving the stop signal. These funcs can be used to stop
// other internal components or update some internal program state.
func start(srv *http.Server, timeout time.Duration, shutdowns ...shutdownFunc) error {
	done := make(chan error, 1)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		ctx := context.Background()
		var cancel context.CancelFunc
		if timeout > 0 {
			ctx, cancel = context.WithTimeout(ctx, timeout)
			defer cancel()
		}

		for _, shutdown := range shutdowns {
			_ = shutdown(ctx)
		}

		done <- srv.Shutdown(ctx)
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return <-done
}

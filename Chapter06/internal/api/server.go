package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	port   string
	server http.Server
	router *mux.Router
	wg     sync.WaitGroup
}

func NewServer(port int) *Server {
	router := mux.NewRouter().StrictSlash(true)
	return &Server{
		router: router,
		port:   fmt.Sprintf(":%d", port),
	}
}

func (c *Server) AddRoute(path string, handler http.HandlerFunc, method string, mwf ...mux.MiddlewareFunc) {
	subRouter := c.router.PathPrefix(path).Subrouter()
	subRouter.Use(mwf...)
	subRouter.HandleFunc("", handler).Methods(method)
	log.Printf("Added route: [%v] [%v]", path, method)
}

// MustStart will start the server and if it cannot bind to the port
// it will exit with a fatal log message
func (c *Server) MustStart() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create the HTML Server
	c.server = http.Server{
		Addr:           fmt.Sprintf("0.0.0.0%s", c.port),
		Handler:        c.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   0 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	// Add to the WaitGroup for the listener goroutine
	c.wg.Add(1)

	// Start the listener
	go func() {
		log.Printf("API server started at %v on http://%s", time.Now().Format(time.Stamp), c.server.Addr)
		if err := c.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("API server failed to start with error: %v\n", err)
		}
		log.Println("API server stopped")
		c.wg.Done()
	}()
}

// Stop stops the API Server
func (c *Server) Stop() error {
	// Create a context to attempt a graceful 5-second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Println("API server: stopping")

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := c.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := c.server.Close(); err != nil {
			log.Printf("API server: stopped with error %v", err)
			return err
		}
		return err
	}

	c.wg.Wait()
	return nil
}

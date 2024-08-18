package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/BelyaevEI/matchmaker/internal/config"

	"github.com/BelyaevEI/platform_common/pkg/closer"
	"github.com/go-chi/chi"
)

// Structure application
type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

// Create new app with config and depends
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDependencies(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {

	// Gracefull shutdown
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Run http srever for given new players
	go func() {
		defer wg.Done()

	}()

	// Create a new playing group
	go func() {
		defer wg.Done()

	}()

	gracefulShutdown(ctx, cancel, wg)
	return nil
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-waitSignal():
		log.Println("terminating: via signal")
	}

	cancel()
	if wg != nil {
		wg.Wait()
	}
}

func waitSignal() chan os.Signal {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	return sigterm
}

// Initialization depends
func (a *App) initDependencies(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Initialization config
func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

// Initialization service provider
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// inititalization http server
func (a *App) initHTTPServer(ctx context.Context) error {

	// New router
	route := chi.NewRouter()

	// Cascade initialization layers
	user := a.serviceProvider.UsersImpl(ctx)

	// Handlers
	route.Post("/users", user.SearchMatch) // Given players for search match

	return nil
}

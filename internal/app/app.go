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

// App structure of application
type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

// NewApp create new app with config and depends
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDependencies(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run application
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

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}

	}()

	// Create a new playing group
	go func() {
		defer wg.Done()

		for {
			err := a.searchingMatch()
			if err != nil {
				log.Fatalf("failed to search match: %v", err)
			}
		}

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
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
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
	err := config.Load("config.env")
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
	route.Post("/users", user.AddUserToPool) // Given players for search match

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: route,
	}

	return nil
}

// Listen and serve
func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// Searching new match
func (a *App) searchingMatch() error {
	// log.Printf("start making new matchs for %v users", a.serviceProvider.ENVConfig().GroupSize())

	if err := a.serviceProvider.usersImpl.CreateMatch(context.Background()); err != nil {
		return err
	}
	return nil
}

package app

import (
	"context"
	"log"

	"github.com/BelyaevEI/matchmaker/internal/api/users"
	"github.com/BelyaevEI/matchmaker/internal/config"
	"github.com/BelyaevEI/matchmaker/internal/repository"
	"github.com/BelyaevEI/matchmaker/internal/service"

	"github.com/BelyaevEI/platform_common/pkg/closer"
	"github.com/BelyaevEI/platform_common/pkg/db"
	"github.com/BelyaevEI/platform_common/pkg/db/pg"
	"github.com/BelyaevEI/platform_common/pkg/db/transaction"
)

// Structure with all entity
type serviceProvider struct {
	httpConfig config.HTTPConfig
	pgConfig   config.PGConfig
	enfConfig  config.EnvConfig

	dbClient  db.Client
	txManager db.TxManager

	usersImpl      *users.Implementation
	userService    service.UserServicer
	userRepository repository.UserRepositorer
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// Read http config
func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

// Read postgres config
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// Read http config
func (s *serviceProvider) ENVConfig() config.EnvConfig {
	if s.enfConfig == nil {
		cfg, err := config.NewEnvConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.enfConfig = cfg
	}

	return s.enfConfig
}

// New connect to postgres db
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {

	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// New transation manager
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// Implemetation api layer
func (s *serviceProvider) UsersImpl(ctx context.Context) *users.Implementation {
	if s.usersImpl == nil {
		s.usersImpl = users.NewImplementation(s.UserService(ctx))
	}

	return s.usersImpl
}

// Implementation service layer
func (s *serviceProvider) UserService(ctx context.Context) service.UserServicer {
	if s.userService == nil {
		s.userService = service.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepositorer {

	if s.userRepository == nil {
		s.userRepository = repository.NewRepository(s.DBClient(ctx),
			s.ENVConfig().StorageFlag(),
			s.ENVConfig().GroupSize(),
		)
	}

	return s.userRepository
}

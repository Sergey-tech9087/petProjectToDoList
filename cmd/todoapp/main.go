package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/Sergey-tech9087/petProjectToDoList/internal/core/logger"
	core_pgx_pool "github.com/Sergey-tech9087/petProjectToDoList/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/Sergey-tech9087/petProjectToDoList/internal/core/transport/http/middleware"
	core_http_server "github.com/Sergey-tech9087/petProjectToDoList/internal/core/transport/http/server"
	users_postgres_repository "github.com/Sergey-tech9087/petProjectToDoList/internal/features/users/repository/postgres"
	users_service "github.com/Sergey-tech9087/petProjectToDoList/internal/features/users/service"
	users_transport_http "github.com/Sergey-tech9087/petProjectToDoList/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("Failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Initialized postgres connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)

	if err != nil {
		logger.Fatal("Failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("Initializing feature", zap.String("feature", "users"))
	usersRepositury := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepositury)

	userTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("Initializing HTTP server")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouteV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouteV1.RegisterRoutes(userTransportHTTP.Routes()...)

	// apiVersionRouteV2 := core_http_server.NewAPIVersionRouter(
	// 	core_http_server.ApiVersion2,
	// 	core_http_middleware.Dummy("api v2 middleware"),
	// )
	// apiVersionRouteV2.RegisterRoutes(userTransportHTTP.Routes()...)

	httpServer.RegisterAPIRoutes(
		apiVersionRouteV1,
		// apiVersionRouteV2,
	)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error ", zap.Error(err))
	}
}

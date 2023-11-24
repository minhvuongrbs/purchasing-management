package cmd

import (
	"purchasing-management/config"
	"purchasing-management/internal/adapters/repository/product"
	"purchasing-management/internal/app"
	"purchasing-management/internal/common/database"
	"purchasing-management/internal/common/logging"
	"purchasing-management/internal/ports"
	"purchasing-management/server"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func ServerRun(_ *cli.Context) error {
	conf := config.LoadConfig()

	logging.Init()
	defer func() {
		// flush logger before exit
		_ = zap.L().Sync()
	}()

	mysqlDatabaseConn := database.NewMySQLDatabaseConn(conf.PrimaryDataSource)
	productRepo := product.NewRepository(mysqlDatabaseConn)
	orderService := app.NewOrderService(productRepo)
	httpServer := ports.NewHttpServer(orderService)

	server.RunHTTPServer(conf.HTTP, httpServer)
	return nil
}

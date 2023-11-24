package server

import (
	"fmt"
	"purchasing-management/internal/common/http"
	"purchasing-management/internal/ports"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

const (
	serverHeader = "purchasing management"
)

func RunHTTPServer(conf http.Config, httpServer ports.HttpServer) {
	fiberApp := initFiberApp(conf)
	apiV1 := fiberApp.Group("/api/v1")

	ports.RegisterHandlers(apiV1, &httpServer)

	listenAddress(fiberApp, conf.Port)
}

func initFiberApp(conf http.Config) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		ServerHeader:     serverHeader,
		Concurrency:      conf.MaxConcurrentConn,
		ReadTimeout:      conf.ReadTimeout,
		WriteTimeout:     conf.WriteTimeout,
		DisableKeepalive: true,
		CaseSensitive:    true,
	})

	fiberApp.Get("/health", healthCheck)

	return fiberApp
}

func listenAddress(fiberApp *fiber.App, port int) {
	logger := zap.S()
	address := fmt.Sprintf(":%d", port)
	logger.Infof("HTTP server serve on %v", address)
	err := fiberApp.Listen(address)
	if err != nil {
		logger.Fatalf("Failed to start HTTP server, %v", err)
	}
}

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("OK")
}

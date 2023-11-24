package ports

import (
	"purchasing-management/internal/app"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type HttpServer struct {
	orderService app.OrderService
}

func NewHttpServer(o app.OrderService) HttpServer {
	return HttpServer{
		orderService: o,
	}
}

func (s HttpServer) CreateProduct(ctx *fiber.Ctx) error {
	err := s.orderService.CreateProduct(ctx.Context(), app.CreateProductRequest{Name: "demo"})
	if err != nil {
		return errors.Wrap(err, "create product")
	}
	return respondSuccess(ctx)
}

func (s HttpServer) GetProducts(ctx *fiber.Ctx) error {
	products, err := s.orderService.GetProducts(ctx.Context())
	if err != nil {
		return errors.Wrap(err, "get products")
	}
	return respondWithData(ctx, products)
}

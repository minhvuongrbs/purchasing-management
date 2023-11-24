package ports

import "github.com/gofiber/fiber/v2"

func respondWithData(fCtx *fiber.Ctx, data any) error {
	return fCtx.Status(fiber.StatusOK).JSON(dataResponse{
		Code:    "success",
		Message: "Success",
		Data:    data,
	})
}

func respondSuccess(fCtx *fiber.Ctx) error {
	return fCtx.Status(fiber.StatusOK).JSON(dataResponse{
		Code:    "success",
		Message: "Success",
	})
}

type dataResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

package handler

import (
	"context"

	"github.com/afiflampard/boilerplate-domain/product"
	"github.com/afiflampard/boilerplate/cmd/service"
	"github.com/cloudwego/hertz/pkg/app"
)

type Handler struct {
	svc service.ProductService
}

func NewHandler(svc service.ProductService) Handler {
	return Handler{svc: svc}
}

func (h *Handler) CreateProductSendEvent(ctx context.Context, c *app.RequestContext) {
	var productInput product.ProductInput
	if err := c.BindJSON(&productInput); err != nil {
		c.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if err := h.svc.CreateProductSendEvent(ctx, &productInput); err != nil {
		c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"message": "Product created successfully",
	})
}

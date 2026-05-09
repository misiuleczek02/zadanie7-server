package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/misiuleczek02/zadanie7-server/internal/store"
)

type ProductHandler struct {
	store *store.ProductStore
}

func NewProductHandler(s *store.ProductStore) *ProductHandler {
	return &ProductHandler{store: s}
}

func (h *ProductHandler) List(c echo.Context) error {
	return c.JSON(http.StatusOK, h.store.All())
}

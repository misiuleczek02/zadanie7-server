package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/misiuleczek02/zadanie7-server/internal/handler"
	"github.com/misiuleczek02/zadanie7-server/internal/store"
)

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	productHandler := handler.NewProductHandler(store.NewProductStore())
	paymentHandler := handler.NewPaymentHandler()

	api := e.Group("/api")
	api.GET("/products", productHandler.List)
	api.POST("/payments", paymentHandler.Create)
	api.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	return e
}

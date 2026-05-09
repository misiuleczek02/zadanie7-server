package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/misiuleczek02/zadanie7-server/internal/model"
)

type PaymentHandler struct{}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (h *PaymentHandler) Create(c echo.Context) error {
	var payment model.Payment
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if payment.Amount <= 0 || payment.Method == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "amount must be positive and method is required"})
	}
	c.Logger().Infof("payment received amount=%.2f method=%s", payment.Amount, payment.Method)
	return c.JSON(http.StatusOK, model.PaymentResponse{
		Status:  "success",
		Message: "Płatność zaakceptowana",
	})
}

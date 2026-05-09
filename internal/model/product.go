package model

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Payment struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
	Method string  `json:"method" validate:"required"`
}

type PaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

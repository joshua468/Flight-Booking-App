package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type PaymentHandler struct {
	stripeKey string
}

func NewPaymentHandler(stripeKey string) *PaymentHandler {
	stripe.Key = stripeKey
	return &PaymentHandler{stripeKey: stripeKey}
}

func (h *PaymentHandler) CreatePaymentIntent(c *gin.Context) {
	var req struct {
		Amount int64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create PaymentIntent with amount
	intent, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_secret": intent.ClientSecret})
}

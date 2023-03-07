package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfItGetsAnErrorIfPriceIsGreaterThanOrEqualZero(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "price required")
}

func TestIfItGetsAnErrorIfTaxIsGreaterThanOrEqualZero(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "tax required")
}

func TestWithAllValidParams(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}

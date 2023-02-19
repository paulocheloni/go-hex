package application_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/paulocheloni/gohex/application"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := &application.Product{
		ID:     "1",
		Name:   "Product 1",
		Price:  10,
		Status: application.DISABLED,
	}

	err := product.Enable()
	assert.Nil(t, err)
	assert.Equal(t, application.ENABLED, product.GetStatus())
}

func TestProduct_Disable(t *testing.T) {
	product := &application.Product{}
	product.Name = "Product 1"
	product.Price = 10
	product.Status = application.DISABLED

	err := product.Disable()
	require.Equal(t, "Product is already disabled", err.Error())

	err = product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.GetStatus())

	err = product.Disable()
	require.Equal(t, "price must be zero to disable the product", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestProduct_IsValid(t *testing.T) {
	product := &application.Product{}
	product.ID = uuid.New().String()
	product.Name = "Product 1"
	product.Price = 10
	product.Status = application.DISABLED

	valid, err := product.IsValid()
	require.Nil(t, err)
	require.True(t, valid)

	product.Status = "invalid"
	valid, err = product.IsValid()
	require.Equal(t, "status must be 'enabled' or 'disabled'", err.Error())
	require.False(t, valid)

	product.Status = application.ENABLED
	product.Price = 0
	valid, err = product.IsValid()
	require.Equal(t, "price must be greater than zero", err.Error())
	require.False(t, valid)

	product.Price = 10
	product.Name = ""
	valid, err = product.IsValid()
	require.Equal(t, "name is required", err.Error())
	require.False(t, valid)

	product.Name = "Product 1"
	product.ID = ""
	valid, err = product.IsValid()
	require.Equal(t, "ID is required", err.Error())
	require.False(t, valid)

	product.ID = uuid.New().String()
	product.Status = "invalid"
	valid, err = product.IsValid()
	require.Equal(t, "status must be 'enabled' or 'disabled'", err.Error())
	require.False(t, valid)

	product.Status = application.ENABLED
	product.Price = -1
	valid, err = product.IsValid()
	require.Equal(t, "price must be greater than zero", err.Error())
	require.False(t, valid)

	product.Price = 10
	product.Name = "ok"
	product.ID = uuid.New().String()
	product.Status = application.ENABLED
	valid, err = product.IsValid()
	require.Nil(t, err)
	require.True(t, valid)

}

package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/paulocheloni/gohex/application"
	mock_application "github.com/paulocheloni/gohex/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get("123").Return(product, nil)

	service := application.ProductService{Persistence: persistence}

	result, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

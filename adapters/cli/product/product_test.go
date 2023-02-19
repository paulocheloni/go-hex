package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	cli "github.com/paulocheloni/gohex/adapters/cli/product"
	mock_application "github.com/paulocheloni/gohex/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 10.0
	productStatus := "enabled"
	productId := "123"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetId().Return(productId).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Save(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product %s enabled, name: %s, price: %f, status: %s", productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product with name %s has been created with the price %f and status of %s", productName, productPrice, productStatus)

	result, err = cli.Run(service, "disable", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("product with name %s has been created with the price %f and status of %s", productName, productPrice, productStatus)

	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}

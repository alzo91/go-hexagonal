package cli_test

import (
	"fmt"
	"testing"

	"github.com/alzo91/go-hexagonal/adapters/cli"
	"github.com/alzo91/go-hexagonal/application"
	mock_application "github.com/alzo91/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := application.NewProduct()
	product.Name = "Product 1"
	product.Price = 25.99
	product.Enable()

	productMocked := mock_application.NewMockProductInterface(ctrl);
	productMocked.EXPECT().GetID().Return(product.GetID()).AnyTimes()
	productMocked.EXPECT().GetName().Return(product.GetName()).AnyTimes()
	productMocked.EXPECT().GetPrice().Return(product.GetPrice()).AnyTimes()
	productMocked.EXPECT().GetStatus().Return(product.GetStatus()).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl);
	service.EXPECT().Create(product.GetName(), product.GetPrice()).Return(productMocked, nil).AnyTimes()
	service.EXPECT().Get(product.GetID()).Return(productMocked, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMocked, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMocked, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %.2f", product.GetID(), product.GetName(), product.GetPrice())

	result, err := cli.Run(service, "create", "", product.GetName(), product.GetPrice())
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product with ID %s has been enabled", product.GetID())
	result, err = cli.Run(service, "enable", product.GetID(), "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product with ID %s has been disabled", product.GetID())
	result, err = cli.Run(service, "disable", product.GetID(), "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	result, err = cli.Run(service, "get", product.GetID(), "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
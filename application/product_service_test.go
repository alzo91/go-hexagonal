package application_test

import (
	"testing"

	"github.com/alzo91/go-hexagonal/application"
	mock_application "github.com/alzo91/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T){
	// Helping us to create a mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// defer ctrl.Finish() // Ele espera tudo que esta acontecendo dentro desta função e depois ele executa

	// Create a new product
	product := mock_application.NewMockProductInterface(ctrl)

	// Create a new service mocked. It is a product fake
	mockRepo := mock_application.NewMockProductPersistenceInterface(ctrl)
	// Expect to call the Get method and return the product
	mockRepo.EXPECT().Get(gomock.Any()).Return(product, nil).Times(1)

	// Create a new service using the database or repository mocked
	service := application.ProductService{
		Persistence: mockRepo,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
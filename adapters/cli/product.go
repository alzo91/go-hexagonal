package cli

import (
	"fmt"

	"github.com/alzo91/go-hexagonal/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productID string,
	productName string,
	productPrice float64,
) (string, error){
	
	switch action {
		case "create":
			product, err := service.Create(productName, productPrice);
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Product ID %s with the name %s has been created with the price %.2f", product.GetID(), product.GetName(), product.GetPrice()), nil
		case "enable":
			product, err := service.Get(productID)
			if err != nil {
				return "", err
			}
			_, err = service.Enable(product)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Product with ID %s has been enabled", productID), nil
		case "disable":
			product, err := service.Get(productID)
			if err != nil {
				return "", err
			}
			_,err = service.Disable(product)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Product with ID %s has been disabled", productID), nil
		
		default:
			res, err := service.Get(productID)
			if err != nil {
				return "", err
			}
			fmt.Println(res)
			return fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %.2f\nStatus: %s", res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus()), err
	}
}

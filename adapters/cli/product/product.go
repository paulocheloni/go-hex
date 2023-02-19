package cli

import (
	"fmt"

	"github.com/paulocheloni/gohex/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Save(productName, productPrice)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s enabled, name: %s, price: %f, status: %s", product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product.GetId())
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product with name %s has been created with the price %f and status of %s", res.GetName(), res.GetPrice(), res.GetStatus())
	case "disable":
		product, _ := service.Get(productId)
		product, _ = service.Disable(product.GetId())
		result = fmt.Sprintf("Product %s enabled with name %f has been created with the price  and status of %s", product.GetName(), product.GetPrice(), product.GetStatus())
	case "get":
		product, _ := service.Get(productId)
		result = fmt.Sprintf("product with name %s has been created with the price %f and status of %s", product.GetName(), product.GetPrice(), product.GetStatus())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product with name %s has been created with the price %f and status of %s", res.GetName(), res.GetPrice(), res.GetStatus())

	}
	return result, nil

}

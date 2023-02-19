package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"optional"`
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Save(name string, price float64) (ProductInterface, error)
	Enable(id string) (ProductInterface, error)
	Disable(id string) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

func NewProduct() *Product {
	product := Product{
		ID: uuid.New().String(),

		Status: DISABLED,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.ID == "" {
		return false, errors.New("ID is required")
	}

	if p.Name == "" {
		return false, errors.New("name is required")
	}

	if p.Price <= 0 {
		return false, errors.New("price must be greater than zero")
	}

	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("status must be 'enabled' or 'disabled'")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Status == DISABLED {
		return errors.New("Product is already disabled")
	}

	if p.Price == 0 {
		p.Status = DISABLED
		return nil

	}

	return errors.New("price must be zero to disable the product")

}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

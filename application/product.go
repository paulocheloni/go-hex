package application

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GEtName() string
	GetStatus() string
	GetPrice() float64
}

package factory

type ICar interface {
	getTypeCar() string
	getBrand() string
	setTypeCar(typeCar string)
	setBrand(brand string)
}

package factory

type MercedesCar struct {
	Car
}

func NewMercedesCar() Car {
	return &MercedesCar{
		Car: Car{
			typeCar: "SUV Car",
			brand:   "Mercedes GLC 300",
		},
	}
}

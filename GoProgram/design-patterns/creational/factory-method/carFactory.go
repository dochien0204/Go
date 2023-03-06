package factory

import "errors"

func GetCar(brand string) (ICar, error) {
	if brand == "Mercedes GLC 300" {
		return NewMercedesCar(), nil
	}

	if brand == "BMW" {
		return NewBMWCar(), nill
	}

	return nil, errors.New("Brand Incorrect")
}

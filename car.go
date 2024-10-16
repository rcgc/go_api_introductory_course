package main

type Car struct {
	Id       string  `json:id`
	Make     string  `json:make`
	Model    string  `json:model`
	Package  string  `json:package`
	Color    string  `json:color`
	Year     int     `json:year`
	Category string  `json:category`
	Mileage  float64 `json:mileage`
	Price    float64 `json:price`
}

var db Db

var m carMiddleware

func (c *Car) getAllCars() ([]Car, error) {
	cars, err := db.getAll()
	if err != nil {
		return []Car{}, err
	}
	return cars, err
}

func (c *Car) createCar() (Car, error) {
	err := m.validate_create(c)
	if err != nil {
		return Car{}, err
	}

	car, err := db.add(c)

	if err != nil {
		return Car{}, err
	}

	return car, nil
}
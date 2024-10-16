package main

import "fmt"

type Db struct {
	cars []Car
}

func (db *Db) getAll() ([]Car, error) {
	return db.cars, nil
}

func (db *Db) add(c *Car) (Car, error) {
	car := Car{Id: c.Id, Make: c.Make, Model: c.Model, Package: c.Package, Color: c.Color, Year: c.Year, Category: c.Category, Mileage: c.Mileage, Price: c.Price}

	for _, v := range db.cars {
		if car.Id == v.Id {
			return car, fmt.Errorf("id already exists")
		}
	}

	db.cars = append(db.cars, car)
	return car, nil
}
package repository

import "github.com/miloalej-dev/vehicle-dealership/pkg/models"

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	// FindByID is a method that returns a vehicle by its ID
	FindByID(id int) (v models.Vehicle, err error)
	// FindByBrand is a method that returns a map of vehicles by brand
	FindByBrand(brand string) (v map[int]models.Vehicle, err error)
	// FindByTransmission is a method that returns a map of vehicles by transmission
	FindByTransmission(transmission string) (v map[int]models.Vehicle, err error)
	// FindByColorAndYear is a method that returns a map of vehicles by color and year
	FindByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error)
	// FindByBrandAndBetweenYears is a method that returns a map of vehicles by brand and between years
	FindByBrandAndBetweenYears(brand string, startYear, endYear int) (v map[int]models.Vehicle, err error)
	// FindByFuelType is a method that returns a map of vehicles by fuel type
	FindByFuelType(fuel string) (v map[int]models.Vehicle, err error)
	// FindByWeight is a method that returns a map of vehicles by weight
	FindByWeight(minWeight, maxWeight float64) (v map[int]models.Vehicle, err error)
	// Save is a method that saves a vehicle
	Save(vehicle models.Vehicle) (v models.Vehicle, err error)
	// SaveAll is a method that saves all vehicles
	SaveAll(vehicles []models.Vehicle) (v map[int]models.Vehicle, err error)
	// UpdateMaxSpeedByID is a method that updates the max speed of a vehicle by its ID
	UpdateMaxSpeedByID(id int, maxSpeed float64) (v models.Vehicle, err error)
	// UpdateFuelTypeByID is a method that updates the fuel type of vehicle by its ID
	UpdateFuelTypeByID(id int, fuelType string) (v models.Vehicle, err error)
	// RemoveByID is a method that removes a vehicle by its ID
	RemoveByID(id int) (err error)
}

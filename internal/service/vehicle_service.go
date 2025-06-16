package service

import "github.com/miloalej-dev/vehicle-dealership/pkg/models"

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// GetVehicles is a method that returns a map of all vehicles
	GetVehicles() (v map[int]models.Vehicle, err error)
	// GetVehiclesByColorAndYear is a method that returns a map of vehicles by brand
	GetVehiclesByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error)
	// GetVehiclesByBrandAndBetweenYears is a method that returns a map of vehicles by brand and between years
	GetVehiclesByBrandAndBetweenYears(brand string, startYear, endYear int) (v map[int]models.Vehicle, err error)
	// GetVehiclesByFuelType is a method that returns a map of vehicles by fuel type
	GetVehiclesByFuelType(fuelType string) (v map[int]models.Vehicle, err error)
	// GetVehiclesByTransmission is a method that returns a map of vehicles by transmission
	GetVehiclesByTransmission(transmission string) (v map[int]models.Vehicle, err error)
	// GetVehiclesByDimensions is a method that returns a map of vehicles by dimensions
	GetVehiclesByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]models.Vehicle, err error)
	// GetVehiclesByWeight is a method that returns a map of vehicles by weight
	GetVehiclesByWeight(minWeight, maxWeight float64) (v map[int]models.Vehicle, err error)
	// ChangeMaxSpeed is a method that changes the max speed of a vehicle by its ID
	ChangeMaxSpeed(id int, maxSpeed float64) (v models.Vehicle, err error)
	// ChangeFuelType is a method that changes the fuel type of a vehicle by its ID
	ChangeFuelType(id int, fuelType string) (v models.Vehicle, err error)
	// AverageMaxSpeedByBrand is a method that returns the average max speed of vehicles by brand
	AverageMaxSpeedByBrand(brand string) (average float64, err error)
	// AverageCapacityByBrand is a method that returns the average capacity of vehicles by brand
	AverageCapacityByBrand(brand string) (average float64, err error)
	// AddVehicle is a method that adds a vehicle to the inventory
	AddVehicle(vehicle models.Vehicle) (v models.Vehicle, err error)
	// AddVehicles is a method that adds vehicles to the inventory
	AddVehicles(vehicles []models.Vehicle) (v map[int]models.Vehicle, err error)
	// DeleteVehicle is a method that deletes a vehicle from the inventory
	DeleteVehicle(id int) (err error)
}

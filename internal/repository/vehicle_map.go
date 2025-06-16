package repository

import (
	"errors"
	"github.com/miloalej-dev/vehicle-dealership/pkg/models"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]models.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]models.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]models.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// FindByID is a method that returns a vehicle by its ID
func (r *VehicleMap) FindByID(id int) (v models.Vehicle, err error) {
	v, ok := r.db[id]
	if !ok {
		err = errors.New("vehicle not found")
		return
	}
	return
}

// FindByBrand is a method that returns a map of vehicles by brand
func (r *VehicleMap) FindByBrand(brand string) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, vehicle := range r.db {
		if vehicle.Brand == brand {
			v[key] = vehicle
		}
	}
	return
}

// FindByTransmission is a method that returns a map of vehicles by transmission
func (r *VehicleMap) FindByTransmission(transmission string) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, vehicle := range r.db {
		if vehicle.Transmission == transmission {
			v[key] = vehicle
		}
	}
	return
}

// FindByColorAndYear is a method that returns a map of vehicles by color and year
func (r *VehicleMap) FindByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, vehicle := range r.db {
		if vehicle.Color == color && vehicle.FabricationYear == year {
			v[key] = vehicle
		}
	}
	return
}

// FindByBrandAndBetweenYears is a method that returns a map of vehicles by brand and between years
func (r *VehicleMap) FindByBrandAndBetweenYears(brand string, startYear, endYear int) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, vehicle := range r.db {
		if vehicle.Brand == brand && vehicle.FabricationYear >= startYear && vehicle.FabricationYear <= endYear {
			v[key] = vehicle
		}
	}
	return
}

// FindByFuelType is a method that returns a map of vehicles by fuel type
func (r *VehicleMap) FindByFuelType(fuel string) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, vehicle := range r.db {
		if vehicle.FuelType == fuel {
			v[key] = vehicle
		}
	}
	return
}

// FindByDimensions is a method that returns a map of vehicles by dimensions
func (r *VehicleMap) FindByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, vehicle := range r.db {
		if vehicle.Length >= minLength && vehicle.Length <= maxLength && vehicle.Width >= minWidth && vehicle.Width <= maxWidth {
			v[key] = vehicle
		}
	}
	return
}

// FindByWeight is a method that returns a map of vehicles by weight
func (r *VehicleMap) FindByWeight(minWeight, maxWeight float64) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, vehicle := range r.db {
		if vehicle.Weight >= minWeight && vehicle.Weight <= maxWeight {
			v[key] = vehicle
		}
	}
	return
}

// Save is a method that saves a vehicle
func (r *VehicleMap) Save(vehicle models.Vehicle) (v models.Vehicle, err error) {

	// validate if the vehicle already exists
	_, ok := r.db[vehicle.Id]

	if ok {
		err = errors.New("vehicle already exists")
		return
	}

	return
}

// SaveAll is a method that saves all vehicles
func (r *VehicleMap) SaveAll(vehicles []models.Vehicle) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for _, vehicle := range vehicles {
		// validate if the vehicle already exists
		_, ok := r.db[vehicle.Id]

		if ok {
			err = errors.New("vehicle already exists")
			return
		}
	}

	v = make(map[int]models.Vehicle)

	for _, vehicle := range vehicles {
		r.db[vehicle.Id] = vehicle

		v[vehicle.Id] = vehicle
	}

	return
}

// UpdateMaxSpeedByID is a method that updates the max speed of a vehicle by its ID
func (r *VehicleMap) UpdateMaxSpeedByID(id int, maxSpeed float64) (v models.Vehicle, err error) {
	v, ok := r.db[id]
	if !ok {
		err = errors.New("vehicle not found")
		return
	}
	v.MaxSpeed = maxSpeed
	r.db[id] = v
	return
}

// UpdateFuelTypeByID is a method that updates the fuel type of vehicle by its ID
func (r *VehicleMap) UpdateFuelTypeByID(id int, fuelType string) (v models.Vehicle, err error) {
	v, ok := r.db[id]
	if !ok {
		err = errors.New("vehicle not found")
		return
	}
	v.FuelType = fuelType
	r.db[id] = v
	return
}

// RemoveByID is a method that removes a vehicle by its ID
func (r *VehicleMap) RemoveByID(id int) (err error) {
	_, ok := r.db[id]
	if !ok {
		err = errors.New("vehicle not found")
		return
	}
	delete(r.db, id)
	return
}

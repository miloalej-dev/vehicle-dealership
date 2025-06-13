package loader

import "github.com/miloalej-dev/vehicle-dealership/pkg/models"

// VehicleLoader is an interface that represents the loader for vehicles
type VehicleLoader interface {
	// Load is a method that loads the vehicles
	Load() (v map[int]models.Vehicle, err error)
}

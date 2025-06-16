package service

import (
	"github.com/miloalej-dev/vehicle-dealership/internal/repository"
	"github.com/miloalej-dev/vehicle-dealership/pkg/models"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp repository.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp repository.VehicleRepository
}

func (s *VehicleDefault) GetVehicles() (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) GetVehiclesByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYear(color, year)
	return
}

func (s *VehicleDefault) GetVehiclesByBrandAndBetweenYears(brand string, startYear, endYear int) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindByBrandAndBetweenYears(brand, startYear, endYear)
	return
}

func (s *VehicleDefault) GetVehiclesByFuelType(fuelType string) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindByFuelType(fuelType)
	return
}

func (s *VehicleDefault) GetVehiclesByTransmission(transmission string) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindByTransmission(transmission)
	return
}

func (s *VehicleDefault) GetVehiclesByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindByDimensions(minLength, maxLength, minWidth, maxWidth)
	return
}

func (s *VehicleDefault) GetVehiclesByWeight(minWeight, maxWeight float64) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindByWeight(minWeight, maxWeight)
	return
}

func (s *VehicleDefault) ChangeMaxSpeed(id int, maxSpeed float64) (v models.Vehicle, err error) {
	v, err = s.rp.UpdateMaxSpeedByID(id, maxSpeed)
	return
}

func (s *VehicleDefault) ChangeFuelType(id int, fuelType string) (v models.Vehicle, err error) {
	v, err = s.rp.UpdateFuelTypeByID(id, fuelType)
	return
}

func (s *VehicleDefault) AverageMaxSpeedByBrand(brand string) (average float64, err error) {
	v, err := s.rp.FindByBrand(brand)
	if err != nil {
		return
	}
	for _, vehicle := range v {
		average += vehicle.MaxSpeed
	}
	average /= float64(len(v))
	return
}

func (s *VehicleDefault) AverageCapacityByBrand(brand string) (average float64, err error) {
	v, err := s.rp.FindByBrand(brand)
	if err != nil {
		return
	}
	for _, vehicle := range v {
		average += float64(vehicle.Capacity)
	}
	average /= float64(len(v))
	return
}

func (s *VehicleDefault) AddVehicle(vehicle models.Vehicle) (v models.Vehicle, err error) {
	v, err = s.rp.Save(vehicle)
	return
}

func (s *VehicleDefault) AddVehicles(vehicles []models.Vehicle) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.SaveAll(vehicles)
	return
}

func (s *VehicleDefault) DeleteVehicle(id int) (err error) {
	err = s.rp.RemoveByID(id)
	return
}

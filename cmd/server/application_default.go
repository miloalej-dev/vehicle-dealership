package server

import (
	"github.com/miloalej-dev/vehicle-dealership/internal/handler"
	"github.com/miloalej-dev/vehicle-dealership/internal/loader"
	"github.com/miloalej-dev/vehicle-dealership/internal/repository"
	"github.com/miloalej-dev/vehicle-dealership/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ChiServer {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ChiServer{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ChiServer is a struct that implements the Application interface
type ChiServer struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the server
func (a *ChiServer) Run() (err error) {
	// dependencies
	// - loader
	ld := loader.NewVehicleJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}
	// - repository
	rp := repository.NewVehicleMap(db)
	// - service
	sv := service.NewVehicleDefault(rp)
	// - handler
	hd := handler.NewVehicleDefault(sv)
	// router
	rt := chi.NewRouter()
	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	// - endpoints
	rt.Route("/vehicles", func(rt chi.Router) {
		// - GET /vehicles
		rt.Get("/", hd.GetAll())
		rt.Get("/color/{color}/year/{year}", hd.GetByColorAndYear)
		rt.Get("/brand/{brand}/year/{startYear}/{endYear}", hd.GetByBrandAndBetweenYears)
		rt.Get("/average_speed/brand/{brand}", hd.GetSpeedAverageByBrand)
		rt.Get("/fuel_type/{type}", hd.GetByFuelType)
		rt.Get("/transmission/{transmission}", hd.GetByTransmissionType)
		rt.Get("/average_capacity/brand/{brand}", hd.GetCapacityAverageByBrand)
		rt.Get("/dimensions", hd.GetByDimensions)
		rt.Get("/weight", hd.GetByWeight)

		// - POST /vehicles
		rt.Post("/", hd.CreateVehicle)
		rt.Post("/batch", hd.CreateVehicles)

		// - PUT /vehicles/{id}
		rt.Put("/{id}/update_fuel", hd.UpdateFuelType)

		// - PATCH /vehicles/{id}
		rt.Patch("/{id}/update_speed", hd.UpdateMaxSpeed)

		// - DELETE /vehicles/{id}
		rt.Delete("/{id}", hd.DeleteVehicle)
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}

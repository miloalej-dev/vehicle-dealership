package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/vehicle-dealership/internal/service"
	"github.com/miloalej-dev/vehicle-dealership/pkg/models"
	"golang.org/x/text/language"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
	"golang.org/x/text/cases"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv service.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv service.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.GetVehicles()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]models.VehicleDoc)
		for key, value := range v {
			data[key] = models.VehicleDoc{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByColorAndYear Obtains vehicles by color and year
func (h *VehicleDefault) GetByColorAndYear(w http.ResponseWriter, r *http.Request) {
	color := chi.URLParam(r, "color")
	year, _ := strconv.Atoi(chi.URLParam(r, "year"))

	if color == "" || year == 0 {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "color and year are required",
			"data":    nil,
		})
		return
	}

	v, err := h.sv.GetVehiclesByColorAndYear(cases.Title(language.English, cases.NoLower).String(color), year)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := make(map[int]models.VehicleDoc)
	for key, value := range v {
		data[key] = models.VehicleDoc{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Height,
			Length:          value.Length,
			Width:           value.Width,
		}
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *VehicleDefault) GetByBrandAndBetweenYears(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	startYear, _ := strconv.Atoi(chi.URLParam(r, "start_year"))
	endYear, _ := strconv.Atoi(chi.URLParam(r, "end_year"))

	if brand == "" || startYear == 0 || endYear == 0 {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "brand, start_year and end_year are required",
			"data":    nil,
		})
		return
	}

	v, err := h.sv.GetVehiclesByBrandAndBetweenYears(cases.Title(language.English, cases.NoLower).String(brand), startYear, endYear)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := make(map[int]models.VehicleDoc)
	for key, value := range v {
		data[key] = models.VehicleDoc{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Height,
			Length:          value.Length,
			Width:           value.Width,
		}
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *VehicleDefault) GetSpeedAverageByBrand(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")

	if brand == "" {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "brand is required",
			"data":    nil,
		})
		return
	}

	v, err := h.sv.AverageMaxSpeedByBrand(cases.Title(language.English, cases.NoLower).String(brand))

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    v,
	})
}

func (h *VehicleDefault) GetByFuelType(w http.ResponseWriter, r *http.Request) {
	fuelType := chi.URLParam(r, "type")

	if fuelType == "" {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "type is required",
			"data":    nil,
		})
		return
	}

	v, err := h.sv.GetVehiclesByFuelType(fuelType)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := make(map[int]models.VehicleDoc)

	for key, value := range v {
		data[key] = models.VehicleDoc{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Height,
			Length:          value.Length,
			Width:           value.Width,
		}
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *VehicleDefault) GetByTransmissionType(w http.ResponseWriter, r *http.Request) {
	transmission := chi.URLParam(r, "transmission")

	if transmission == "" {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "type is required",
			"data":    nil,
		})
		return
	}

	v, err := h.sv.GetVehiclesByTransmission(transmission)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := make(map[int]models.VehicleDoc)

	for key, value := range v {
		data[key] = models.VehicleDoc{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Height,
			Length:          value.Length,
			Width:           value.Width,
		}
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *VehicleDefault) GetCapacityAverageByBrand(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")

	if brand == "" {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "brand is required",
			"data":    nil,
		})
		return
	}

	v, err := h.sv.AverageCapacityByBrand(cases.Title(language.English, cases.NoLower).String(brand))

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    v,
	})

}

func (h *VehicleDefault) GetByDimensions(w http.ResponseWriter, r *http.Request) {
	length := r.URL.Query().Get("length")
	width := r.URL.Query().Get("width")

	if length == "" || width == "" {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "length and width are required",
			"data":    nil,
		})
		return
	}

	lengths := strings.Split(length, "-")
	widths := strings.Split(width, "-")

	if len(lengths) != 2 || len(widths) != 2 {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "length and width must be in the format min-max",
			"data":    nil,
		})
		return
	}

	minLength, _ := strconv.ParseFloat(lengths[0], 64)
	maxLength, _ := strconv.ParseFloat(lengths[1], 64)
	minWidth, _ := strconv.ParseFloat(widths[0], 64)
	maxWidth, _ := strconv.ParseFloat(widths[1], 64)

	v, err := h.sv.GetVehiclesByDimensions(minLength, maxLength, minWidth, maxWidth)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := make(map[int]models.VehicleDoc)

	for key, value := range v {
		data[key] = models.VehicleDoc{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Height,
			Length:          value.Length,
			Width:           value.Width,
		}
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *VehicleDefault) GetByWeight(w http.ResponseWriter, r *http.Request) {
	minParam := r.URL.Query().Get("min")
	maxParam := r.URL.Query().Get("max")

	if minParam == "" || maxParam == "" {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "min and maxParam are required",
			"data":    nil,
		})
		return
	}

	minWeight, _ := strconv.ParseFloat(minParam, 64)
	maxWeight, _ := strconv.ParseFloat(maxParam, 64)

	v, err := h.sv.GetVehiclesByWeight(minWeight, maxWeight)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := make(map[int]models.VehicleDoc)

	for key, value := range v {
		data[key] = models.VehicleDoc{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Height,
			Length:          value.Length,
			Width:           value.Width,
		}
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}
func (h *VehicleDefault) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle models.VehicleDoc
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}

	v, err := h.sv.AddVehicle(
		models.Vehicle{
			Id: vehicle.ID,
			VehicleAttributes: models.VehicleAttributes{
				Brand:           vehicle.Brand,
				Model:           vehicle.Model,
				Registration:    vehicle.Registration,
				Color:           vehicle.Color,
				FabricationYear: vehicle.FabricationYear,
				Capacity:        vehicle.Capacity,
				MaxSpeed:        vehicle.MaxSpeed,
				FuelType:        vehicle.FuelType,
				Transmission:    vehicle.Transmission,
				Dimensions: models.Dimensions{
					Height: vehicle.Height,
					Length: vehicle.Length,
					Width:  vehicle.Width,
				},
			},
		},
	)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := models.VehicleDoc{
		ID:              v.Id,
		Brand:           v.Brand,
		Model:           v.Model,
		Registration:    v.Registration,
		Color:           v.Color,
		FabricationYear: v.FabricationYear,
		Capacity:        v.Capacity,
		MaxSpeed:        v.MaxSpeed,
		FuelType:        v.FuelType,
		Transmission:    v.Transmission,
		Weight:          v.Weight,
		Height:          v.Height,
		Length:          v.Length,
		Width:           v.Width,
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *VehicleDefault) CreateVehicles(w http.ResponseWriter, r *http.Request) {
	var vehicles []models.VehicleDoc
	if err := json.NewDecoder(r.Body).Decode(&vehicles); err != nil {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}
	newVehicles := make([]models.Vehicle, 0)
	for _, v := range vehicles {
		newVehicles = append(newVehicles, models.Vehicle{
			Id: v.ID,
			VehicleAttributes: models.VehicleAttributes{
				Brand:           v.Brand,
				Model:           v.Model,
				Registration:    v.Registration,
				Color:           v.Color,
				FabricationYear: v.FabricationYear,
				Capacity:        v.Capacity,
				MaxSpeed:        v.MaxSpeed,
				FuelType:        v.FuelType,
				Transmission:    v.Transmission,
				Dimensions: models.Dimensions{
					Height: v.Height,
					Length: v.Length,
					Width:  v.Width,
				},
			},
		})
	}

	v, err := h.sv.AddVehicles(newVehicles)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := make(map[int]models.VehicleDoc)

	for key, value := range v {
		data[key] = models.VehicleDoc{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Height,
			Length:          value.Length,
			Width:           value.Width,
		}
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *VehicleDefault) UpdateMaxSpeed(w http.ResponseWriter, r *http.Request) {
	urlParam := chi.URLParam(r, "id")

	if urlParam == "" {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "id is required",
			"data":    nil,
		})
		return
	}

	id, _ := strconv.Atoi(urlParam)

	attributes := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&attributes); err != nil {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}

	v, err := h.sv.ChangeMaxSpeed(id, attributes["max_speed"].(float64))

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := models.VehicleDoc{
		ID:              v.Id,
		Brand:           v.Brand,
		Model:           v.Model,
		Registration:    v.Registration,
		Color:           v.Color,
		FabricationYear: v.FabricationYear,
		Capacity:        v.Capacity,
		MaxSpeed:        v.MaxSpeed,
		FuelType:        v.FuelType,
		Transmission:    v.Transmission,
		Weight:          v.Weight,
		Height:          v.Height,
		Length:          v.Length,
		Width:           v.Width,
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *VehicleDefault) UpdateFuelType(w http.ResponseWriter, r *http.Request) {
	urlParam := chi.URLParam(r, "id")

	if urlParam == "" {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "id is required",
			"data":    nil,
		})
		return
	}

	id, _ := strconv.Atoi(urlParam)

	var vehicle models.VehicleDoc

	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}

	v, err := h.sv.ChangeFuelType(id, vehicle.FuelType)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	data := models.VehicleDoc{
		ID:              v.Id,
		Brand:           v.Brand,
		Model:           v.Model,
		Registration:    v.Registration,
		Color:           v.Color,
		FabricationYear: v.FabricationYear,
		Capacity:        v.Capacity,
		MaxSpeed:        v.MaxSpeed,
		FuelType:        v.FuelType,
		Transmission:    v.Transmission,
		Weight:          v.Weight,
		Height:          v.Height,
		Length:          v.Length,
		Width:           v.Width,
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *VehicleDefault) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	urlParam := chi.URLParam(r, "id")

	if urlParam == "" {
		response.JSON(w, http.StatusBadRequest, map[string]any{
			"message": "id is required",
			"data":    nil,
		})
		return
	}

	id, _ := strconv.Atoi(urlParam)

	err := h.sv.DeleteVehicle(id)

	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"data":    nil,
		})
		return
	}

	response.JSON(w, http.StatusNoContent, map[string]any{
		"message": "success",
		"data":    nil,
	})
}

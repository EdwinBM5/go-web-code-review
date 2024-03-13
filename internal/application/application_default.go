package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
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
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
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

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the application
func (a *ServerChi) Run() (err error) {
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

		// - POST /vehicles
		rt.Post("/", hd.Create())

		// - GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
		rt.Get("/brand/{brand}/between/{start_year}/{end_year}", hd.GetByBrandAndRangeYear())

		// - GET /vehicles/color/{color}/year/{year}
		rt.Get("/color/{color}/year/{year}", hd.GetByColorAndYear())

		// - GET /vehicles/average-speed/brand/{brand}
		rt.Get("/average-speed/brand/{brand}", hd.GetAverageSpeedByBrand())

		// - PATCH /vehicles/{id}/update-speed
		rt.Patch("/{id}/update-speed", hd.UpdateMaxSpeed())

		// - GET /vehicles/fuel-type/{type}
		rt.Get("/fuel-type/{type}", hd.GetByFuelType())

		// - DELETE /vehicles/{id}
		rt.Delete("/{id}", hd.Delete())

		// - GET /vehicles/transmission/{type}
		rt.Get("/transmission/{type}", hd.GetByTransmissionType())

		// - PATCH /vehicles/{id}/update-fuel
		rt.Patch("/{id}/update-fuel", hd.UpdateFuelType())

		// - GET /vehicles/average-capacity/brand/{brand}
		rt.Get("/average-capacity/brand/{brand}", hd.GetAverageCapacityByBrand())

		// - GET /vehicles/dimensions?length={min_length}-{max_length}&width={min_width}-{max_width}
		rt.Get("/dimensions", hd.GetByDimensions())

		// - GET /vehicles/weight?min={min_weight}&max={max_weight}
		rt.Get("/weight", hd.GetByWeightRange())
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}

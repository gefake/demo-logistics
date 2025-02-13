package main

import (
	"logistic_api/pkg/database"
	"logistic_api/pkg/handler"
	"logistic_api/pkg/logger"
	"logistic_api/pkg/service"
)

//geo "github.com/kellydunn/golang-geo"

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @host localhost:8080

func main() {
	// Create coordinates for City A and City B
	// cityALat := 40.7128
	// cityALng := -74.0060
	// cityBLat := 34.0522
	// cityBLng := -118.2437

	// // Create new points for City A and City B
	// cityAPoint := geo.NewPoint(cityALat, cityALng)
	// cityBPoint := geo.NewPoint(cityBLat, cityBLng)

	// // Calculate distance between City A and City B
	// distance := cityAPoint.GreatCircleDistance(cityBPoint)

	// fmt.Printf("Distance between City A and City B: %.2f km\n", distance)

	dbServices := database.DBService{
		CargoRepository:            &database.Cargo{},
		UserRepository:             &database.User{},
		PositionRepository:         &database.Position{},
		RoleRepository:             &database.Role{},
		OrderRepository:            &database.Order{},
		SupplierRepository:         &database.Supplier{},
		ProductRepository:          &database.Product{},
		DeliveryRouteRepository:    &database.DeliveryRoute{},
		WarehouseRepository:        &database.Warehouse{},
		DeliveryScheduleRepository: &database.DeliverySchedule{},
	}
	services := service.NewService(&dbServices)

	handlers := handler.Handler{
		Services: services,
	}

	server := Server{}
	if err := server.Start("8080", handlers.InitRoutes()); err != nil {
		logger.Log.Error(err)
	}
}

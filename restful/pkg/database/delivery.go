package database

import (
	"time"
)

type DeliveryRoute struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	CargoID       uint      `json:"cargo_id" binding:"required"`
	Cargo         Cargo     `gorm:"foreignKey:CargoID" binding:"-"`
	DriverID      uint      `json:"driver_id" binding:"required"`
	Driver        User      `gorm:"foreignKey:DriverID" binding:"-"`
	DepartureDate time.Time `json:"departure_date" binding:"required"`
	StartPoint    Point     `json:"start_point" gorm:"type:point" binding:"required"`
	EndPoint      Point     `json:"end_point" gorm:"type:point" binding:"required"`
	ArrivalDate   time.Time `json:"arrival_date" binding:"required"`
	Status        string    `json:"status" binding:"required"`
}

func (d *DeliveryRoute) CreateDeliveryRoute(route DeliveryRoute) (uint, error) {
	result := DataSource.Context.Create(&route)
	if result.Error != nil {
		return 0, result.Error
	}
	return route.ID, nil
}

// GetDeliveryRouteByID retrieves a delivery route by its ID from the database.
func (d *DeliveryRoute) GetDeliveryRouteByID(routeID uint) (DeliveryRoute, error) {
	var route DeliveryRoute
	result := DataSource.Context.Preload("Cargo.Order.OrderItems.Product").Preload("Cargo.Client").Preload("Driver").First(&route, routeID)
	if result.Error != nil {
		return DeliveryRoute{}, result.Error
	}
	return route, nil
}

// UpdateDeliveryRoute updates a delivery route by its ID in the database.
func (d *DeliveryRoute) UpdateDeliveryRoute(routeID uint, route DeliveryRoute) error {
	result := DataSource.Context.Model(&DeliveryRoute{}).Where("id = ?", routeID).Updates(route)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *DeliveryRoute) GetAllRoutes() ([]DeliveryRoute, error) {
	var routes []DeliveryRoute
	result := DataSource.Context.Preload("Cargo.Order").Preload("Cargo.Client").Preload("Driver").Find(&routes)
	if result.Error != nil {
		return nil, result.Error
	}
	return routes, nil
}

// DeleteDeliveryRoute deletes a delivery route by its ID from the database.
func (d *DeliveryRoute) DeleteDeliveryRoute(routeID uint) error {
	result := DataSource.Context.Delete(&DeliveryRoute{}, routeID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

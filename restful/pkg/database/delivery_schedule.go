package database

import (
	"time"
)

type DeliveryScheduleProduct struct {
	ID                 uint             `json:"id" gorm:"primaryKey"`
	DeliveryScheduleID uint             `json:"delivery_schedule_id" binding:"-"`
	DeliverySchedule   DeliverySchedule `json:"delivery_schedule"  gorm:"foreignKey:DeliveryScheduleID" binding:"-"`
	ProductID          uint             `json:"product_id" binding:"required"`
	Product            Product          `json:"product"  gorm:"foreignKey:ProductID" binding:"-"`
	Quantity           int              `json:"quantity" binding:"required"`
}

type DeliverySchedule struct {
	ID          uint                      `json:"id" gorm:"primaryKey"`
	Date        time.Time                 `json:"date" binding:"required"`
	WarehouseID uint                      `json:"warehouse_id" binding:"required"`
	Warehouse   Warehouse                 `json:"warehouse"  gorm:"foreignKey:WarehouseID" binding:"-"`
	Products    []DeliveryScheduleProduct `json:"products" gorm:"foreignKey:DeliveryScheduleID" binding:"required"`
}

func (ds *DeliverySchedule) CreateDeliverySchedule(schedule DeliverySchedule) (uint, error) {
	products := schedule.Products
	schedule.Products = nil
	result := DataSource.Context.Create(&schedule)
	if result.Error != nil {
		return 0, result.Error
	}

	if len(products) > 0 {
		for _, product := range products {
			dsProduct := DeliveryScheduleProduct{
				DeliveryScheduleID: schedule.ID,
				ProductID:          product.ProductID,
				Quantity:           product.Quantity,
			}
			result := DataSource.Context.Create(&dsProduct)
			if result.Error != nil {
				return 0, result.Error
			}
		}
	}

	return schedule.ID, nil
}

func (ds *DeliverySchedule) UpdateDeliverySchedule(scheduleID uint, newSchedule DeliverySchedule) error {
	result := DataSource.Context.Model(ds).Where("id = ?", scheduleID).Updates(newSchedule)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ds *DeliverySchedule) DeleteDeliverySchedule(scheduleID uint) error {
	result := DataSource.Context.Delete(&DeliverySchedule{}, scheduleID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ds *DeliverySchedule) GetDeliveryScheduleByID(scheduleID uint) (DeliverySchedule, error) {
	var schedule DeliverySchedule
	result := DataSource.Context.Preload("Warehouse").Preload("Products.Product").First(&schedule, scheduleID)
	if result.Error != nil {
		return DeliverySchedule{}, result.Error
	}
	return schedule, nil
}

func (ds *DeliverySchedule) GetAllDeliverySchedules() ([]DeliverySchedule, error) {
	var schedules []DeliverySchedule
	result := DataSource.Context.Preload("Warehouse").Preload("Products.Product").Find(&schedules)
	if result.Error != nil {
		return nil, result.Error
	}
	return schedules, nil
}

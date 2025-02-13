package database

import (
	"fmt"
)

type CargoProduct struct {
	CargoID   uint
	ProductID uint
	Quantity  int
}

type Cargo struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name" binding:"required"`
	Description   string         `json:"description" binding:"required"`
	Weight        float64        `json:"weight" binding:"required"`
	Status        string         `json:"status" binding:"required"`
	OrderID       uint           `json:"order_id" binding:"required"`
	Order         Order          `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" binding:"-"`
	ClientID      uint           `json:"client_id" binding:"required"`
	Client        User           `gorm:"foreignKey:ClientID" binding:"-"`
	CargoProducts []CargoProduct `gorm:"foreignKey:CargoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" binding:"-"`
}

func (c *Cargo) CreateCargo(cargo Cargo) (uint, error) {
	for _, cargoProduct := range cargo.CargoProducts {
		var product Product
		result := DataSource.Context.First(&product, cargoProduct.ProductID)
		if result.Error != nil {
			return 0, result.Error
		}

		if product.Quantity < cargoProduct.Quantity {
			return 0, fmt.Errorf("Недопустимое количество продукта %s на складе. Доступно: %d, Запрашиваемое: %d", product.Name, product.Quantity, cargoProduct.Quantity)
		}
	}

	db := DataSource.Context

	result := db.Create(&cargo)

	if result.Error != nil {
		return 0, result.Error
	}

	return cargo.ID, nil
}

func (c *Cargo) GetCargoByID(cargoID uint) (Cargo, error) {
	db := DataSource.Context
	var cargo Cargo
	result := db.Preload("CargoProducts").Preload("Client").First(&cargo, cargoID)
	if result.Error != nil {
		return Cargo{}, result.Error
	}
	return cargo, nil
}

func (c *Cargo) UpdateCargo(cargoID uint, cargo Cargo) error {
	db := DataSource.Context
	result := db.Model(&cargo).Where("id = ?", cargoID).Updates(cargo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Cargo) DeleteCargo(cargoID uint) error {
	db := DataSource.Context
	result := db.Delete(&Cargo{}, cargoID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Cargo) GetAllCargos(page, perPage int) ([]Cargo, int, error) {
	db := DataSource.Context

	// Calculate the offset based on the page and perPage values
	offset := (page - 1) * perPage

	// Define the query to retrieve all cargos with pagination
	var cargos []Cargo
	result := db.Offset(offset).Limit(perPage).Find(&cargos)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	// Get the total count of cargos
	var totalCount int64
	db.Model(&Cargo{}).Count(&totalCount)

	return cargos, int(totalCount), nil
}

func (c *Cargo) GetCargoOrdersByUserID(userID uint) ([]Cargo, error) {
	db := DataSource.Context
	var cargos []Cargo
	result := db.Preload("CargoProducts").Preload("Client").Where("client_id = ?", userID).Find(&cargos)
	if result.Error != nil {
		return nil, result.Error
	}
	return cargos, nil
}

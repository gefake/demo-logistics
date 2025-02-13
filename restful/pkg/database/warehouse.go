package database

type Warehouse struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Address  string    `json:"address" binding:"required"`
	Position Point     `json:"position" gorm:"type:point" binding:"required"`
	Products []Product `json:"products" gorm:"foreignKey:WarehouseID;references:ID" binding:"-"`
}

func (w *Warehouse) CreateWarehouse(warehouse Warehouse) (uint, error) {
	result := DataSource.Context.Create(&warehouse)
	if result.Error != nil {
		return 0, result.Error
	}
	return warehouse.ID, nil
}

func (w *Warehouse) GetWarehouseByID(warehouseID uint) (Warehouse, error) {
	var warehouse Warehouse
	result := DataSource.Context.Preload("Products").First(&warehouse, warehouseID)
	if result.Error != nil {
		return warehouse, result.Error
	}
	return warehouse, nil
}

func (w *Warehouse) GetAllWarehouses() ([]Warehouse, error) {
	var warehouses []Warehouse

	result := DataSource.Context.Preload("Products").Find(&warehouses)

	if result.Error != nil {
		return nil, result.Error
	}

	return warehouses, nil
}

func (w *Warehouse) UpdateWarehouse(warehouseID uint, newWarehouse Warehouse) error {
	result := DataSource.Context.Model(&w).Where("id = ?", warehouseID).Updates(&newWarehouse)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w *Warehouse) DeleteWarehouse(warehouseID uint) error {
	result := DataSource.Context.Delete(&w, warehouseID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

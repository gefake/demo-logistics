package database

type Supplier struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name" binding:"required"`
	Contact  string    `json:"contact" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Products []Product `gorm:"foreignKey:SupplierID" json:"products"`
}

func (s *Supplier) CreateSupplier(supplier Supplier) (uint, error) {
	result := DataSource.Context.Create(&supplier)

	if result.Error != nil {
		return 0, result.Error
	}

	return s.ID, nil
}

func (s *Supplier) GetSupplierByID(supplierID uint) (Supplier, error) {
	var supplier Supplier
	result := DataSource.Context.Preload("Products").First(&supplier, supplierID)
	if result.Error != nil {
		return Supplier{}, result.Error
	}
	return supplier, nil
}

func (s *Supplier) UpdateSupplier(supplierID uint, newSupplier Supplier) error {
	result := DataSource.Context.Model(s).Where("id = ?", supplierID).Updates(newSupplier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Supplier) DeleteSupplier(supplierID uint) error {
	result := DataSource.Context.Delete(s, supplierID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Supplier) GetAllSuppliers() ([]Supplier, error) {
	var suppliers []Supplier
	result := DataSource.Context.Preload("Products").Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return suppliers, nil
}

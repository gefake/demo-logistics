package database

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	SupplierID  uint      `json:"supplier_id" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Supplier    Supplier  `gorm:"foreignKey:SupplierID" binding:"-"`
	Unit        string    `json:"unit" binding:"required"`
	WarehouseID uint      `json:"warehouse_id" binding:"required"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" binding:"-"`
	Description string    `json:"description" binding:"required"`
	Quantity    int       `json:"quantity" binding:"required"`
}

type CatsWithProds struct {
	Category string
	Products []Product
}

func (p *Product) GetCategories() ([]CatsWithProds, error) {
	var cats []CatsWithProds

	err := DataSource.Context.Model(&Product{}).
		Select("category, GROUP_CONCAT(name) AS product_names, GROUP_CONCAT(id) AS product_ids").
		Distinct("category").
		Where("category IS NOT NULL").
		Group("category").
		Find(&cats).Error

	if err != nil {
		return nil, err
	}

	for i := range cats {
		var products []Product
		err := DataSource.Context.Model(&Product{}).
			Where("category = ?", cats[i].Category).
			Preload("Supplier").
			Find(&products).Error

		if err != nil {
			return nil, err
		}

		cats[i].Products = products
	}

	return cats, nil
}

func (p *Product) CreateProduct(product Product) (uint, error) {
	result := DataSource.Context.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return p.ID, nil
}

func (p *Product) GetProductByID(productID uint) (Product, error) {
	var product Product
	result := DataSource.Context.Preload("Supplier").Preload("Warehouse").First(&product, productID)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

func (p *Product) HasQuantity(productID uint, needQuantity int) (bool, error) {
	var product Product
	result := DataSource.Context.First(&product, productID)

	if result.Error != nil {
		return false, result.Error
	}

	return product.Quantity >= needQuantity, nil
}

func (p *Product) UpdateProduct(productID uint, newProduct Product) error {
	result := DataSource.Context.Model(p).Where("id = ?", productID).Updates(newProduct)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Product) DeleteProduct(productID uint) error {
	result := DataSource.Context.Delete(p, productID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Product) GetAllProducts() ([]Product, error) {
	var products []Product
	result := DataSource.Context.Preload("Supplier").Preload("Warehouse").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

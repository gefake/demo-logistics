package database

import "gorm.io/gorm"

type Connect struct {
	Context *gorm.DB
}

const (
	UserRoleName  = "Пользователь"
	AdminRoleName = "Администратор"
)

var Positions = []string{"Клиент", "Водитель", "Менеджер", "Сборщик заказов", "Директор"}

func (s *Connect) createDefaultRolesIfNotExist() error {
	var userRole Role
	db := s.Context

	db.Where("name = ?", UserRoleName).First(&userRole)
	if userRole.ID == 0 {
		userRole = Role{Name: UserRoleName}
		if err := db.Create(&userRole).Error; err != nil {
			return err
		}
	}

	var adminRole Role
	db.Where("name = ?", AdminRoleName).First(&adminRole)
	if adminRole.ID == 0 {
		adminRole = Role{Name: AdminRoleName, IsAdmin: true}
		if err := db.Create(&adminRole).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *Connect) createDefaultPositionsIfNotExist() error {
	db := s.Context

	for _, positionName := range Positions {
		var position Position
		db.Where("name = ?", positionName).First(&position)
		if position.ID == 0 {
			position = Position{Name: positionName}
			if err := db.Create(&position).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

var DataSource *Connect

type CargoRepository interface {
	CreateCargo(cargo Cargo) (uint, error)
	GetCargoByID(cargoID uint) (Cargo, error)
	UpdateCargo(cargoID uint, cargo Cargo) error
	DeleteCargo(cargoID uint) error
	GetAllCargos(page, perPage int) ([]Cargo, int, error)
	GetCargoOrdersByUserID(userID uint) ([]Cargo, error)
}

type UserRepository interface {
	CreateUser(user User) (uint, error)
	GetUserByID(userID uint) (User, error)
	GetByLogin(login string) (User, error)
	UpdateUser(userID uint, user User) error
	DeleteUser(userID uint) error
}

type PositionRepository interface {
	CreatePosition(position Position) (uint, error)
	GetPositionByID(positionID uint) (Position, error)
	UpdatePosition(positionID uint, position Position) error
	DeletePosition(positionID uint) error
	GetPositionByName(position string) (Position, error)
}

type RoleRepository interface {
	CreateRole(role Role) (uint, error)
	GetRoleByID(roleID uint) (Role, error)
	GetRoleByName(role string) (Role, error)
	UpdateRole(roleID uint, role Role) error
	DeleteRole(roleID uint) error
}

type OrderRepository interface {
	CreateOrder(order Order) (uint, error)
	GetOrderByID(orderID uint) (Order, error)
	UpdateOrder(orderID uint, order Order) error
	DeleteOrder(orderID uint) error
	GetAllOrders() ([]Order, error)
	GetDeliveryByOrder(orderID uint) (DeliveryRoute, error)
}

type SupplierRepository interface {
	CreateSupplier(supplier Supplier) (uint, error)
	GetSupplierByID(supplierID uint) (Supplier, error)
	UpdateSupplier(supplierID uint, supplier Supplier) error
	DeleteSupplier(supplierID uint) error
	GetAllSuppliers() ([]Supplier, error)
}

type ProductRepository interface {
	CreateProduct(product Product) (uint, error)
	GetProductByID(productID uint) (Product, error)
	UpdateProduct(productID uint, product Product) error
	DeleteProduct(productID uint) error
	HasQuantity(productID uint, needQuantity int) (bool, error)
	GetCategories() ([]CatsWithProds, error)
	GetAllProducts() ([]Product, error)
}

type DeliveryRouteRepository interface {
	CreateDeliveryRoute(route DeliveryRoute) (uint, error)
	GetDeliveryRouteByID(routeID uint) (DeliveryRoute, error)
	UpdateDeliveryRoute(routeID uint, route DeliveryRoute) error
	DeleteDeliveryRoute(routeID uint) error
	GetAllRoutes() ([]DeliveryRoute, error)
}

type WarehouseRepository interface {
	CreateWarehouse(warehouse Warehouse) (uint, error)
	GetWarehouseByID(warehouseID uint) (Warehouse, error)
	UpdateWarehouse(warehouseID uint, warehouse Warehouse) error
	DeleteWarehouse(warehouseID uint) error
	GetAllWarehouses() ([]Warehouse, error)
}

type DeliveryScheduleRepository interface {
	CreateDeliverySchedule(deliverySchedule DeliverySchedule) (uint, error)
	GetDeliveryScheduleByID(deliveryScheduleID uint) (DeliverySchedule, error)
	UpdateDeliverySchedule(deliveryScheduleID uint, deliverySchedule DeliverySchedule) error
	DeleteDeliverySchedule(deliveryScheduleID uint) error
	GetAllDeliverySchedules() ([]DeliverySchedule, error)
}

type DBService struct {
	CargoRepository
	UserRepository
	PositionRepository
	RoleRepository
	OrderRepository
	SupplierRepository
	ProductRepository
	DeliveryRouteRepository
	WarehouseRepository
	DeliveryScheduleRepository
}

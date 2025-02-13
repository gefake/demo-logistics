package service

import (
	"logistic_api/pkg/database"
)

type CargoRepository interface {
	CreateCargo(cargo database.Cargo) (uint, error)
	GetCargoByID(cargoID uint) (database.Cargo, error)
	UpdateCargo(cargoID uint, cargo database.Cargo) error
	DeleteCargo(cargoID uint) error
	GetAllCargos(page, perPage int) ([]database.Cargo, int, error)
	GetCargoOrdersByUserID(userID uint) ([]database.Cargo, error)
}

type UserRepository interface {
	CreateUser(user database.User) (uint, error)
	GetUserByID(userID uint) (database.User, error)
	GetByLogin(login string) (database.User, error)
	UpdateUser(userID uint, user database.User) error
	DeleteUser(userID uint) error
}

type PositionRepository interface {
	CreatePosition(position database.Position) (uint, error)
	GetPositionByID(positionID uint) (database.Position, error)
	GetPositionByName(position string) (database.Position, error)
	UpdatePosition(positionID uint, position database.Position) error
	DeletePosition(positionID uint) error
}

type RoleRepository interface {
	CreateRole(role database.Role) (uint, error)
	GetRoleByID(roleID uint) (database.Role, error)
	GetRoleByName(role string) (database.Role, error)
	UpdateRole(roleID uint, role database.Role) error
	DeleteRole(roleID uint) error
}

type OrderRepository interface {
	CreateOrder(order database.Order) (uint, error)
	GetOrderByID(orderID uint) (database.Order, error)
	UpdateOrder(orderID uint, order database.Order) error
	DeleteOrder(orderID uint) error
	GetAllOrders() ([]database.Order, error)
	GetDeliveryByOrder(orderID uint) (database.DeliveryRoute, error)
}

type SupplierRepository interface {
	CreateSupplier(supplier database.Supplier) (uint, error)
	GetSupplierByID(supplierID uint) (database.Supplier, error)
	UpdateSupplier(supplierID uint, supplier database.Supplier) error
	DeleteSupplier(supplierID uint) error
	GetAllSuppliers() ([]database.Supplier, error)
}

type ProductRepository interface {
	CreateProduct(product database.Product) (uint, error)
	GetProductByID(productID uint) (database.Product, error)
	UpdateProduct(productID uint, product database.Product) error
	DeleteProduct(productID uint) error
	HasQuantity(productID uint, needQuantity int) (bool, error)
	GetCategories() ([]database.CatsWithProds, error)
	GetAllProducts() ([]database.Product, error)
}

type DeliveryRouteRepository interface {
	CreateDeliveryRoute(route database.DeliveryRoute) (uint, error)
	GetDeliveryRouteByID(routeID uint) (database.DeliveryRoute, error)
	UpdateDeliveryRoute(routeID uint, route database.DeliveryRoute) error
	DeleteDeliveryRoute(routeID uint) error
	GetAllRoutes() ([]database.DeliveryRoute, error)
}

type WarehouseRepository interface {
	CreateWarehouse(warehouse database.Warehouse) (uint, error)
	GetWarehouseByID(warehouseID uint) (database.Warehouse, error)
	UpdateWarehouse(warehouseID uint, warehouse database.Warehouse) error
	DeleteWarehouse(warehouseID uint) error
	GetAllWarehouses() ([]database.Warehouse, error)
}

type DeliveryScheduleRepository interface {
	CreateDeliverySchedule(deliverySchedule database.DeliverySchedule) (uint, error)
	GetDeliveryScheduleByID(deliveryScheduleID uint) (database.DeliverySchedule, error)
	UpdateDeliverySchedule(deliveryScheduleID uint, deliverySchedule database.DeliverySchedule) error
	DeleteDeliverySchedule(deliveryScheduleID uint) error
	GetAllDeliverySchedules() ([]database.DeliverySchedule, error)
}

type Service struct {
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

func NewService(dbRepo *database.DBService) *Service {
	return &Service{
		CargoRepository:            dbRepo.CargoRepository,
		UserRepository:             dbRepo.UserRepository,
		PositionRepository:         dbRepo.PositionRepository,
		RoleRepository:             dbRepo.RoleRepository,
		OrderRepository:            dbRepo.OrderRepository,
		SupplierRepository:         dbRepo.SupplierRepository,
		ProductRepository:          dbRepo.ProductRepository,
		DeliveryRouteRepository:    dbRepo.DeliveryRouteRepository,
		WarehouseRepository:        dbRepo.WarehouseRepository,
		DeliveryScheduleRepository: dbRepo.DeliveryScheduleRepository,
	}
}

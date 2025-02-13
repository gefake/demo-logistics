package database

import (
	"fmt"
	"time"
)

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id" binding:"required"`
	Order     Order   `gorm:"foreignKey:OrderID" binding:"-"`
	ProductID uint    `json:"product_id" binding:"required"`
	Product   Product `json:"product_name" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" binding:"-"`
	Quantity  int     `json:"quantity" binding:"required"`
}

func (s *OrderItem) CreateOrderItem(orderItem OrderItem) (uint, error) {
	prod := &Product{}
	result := DataSource.Context.First(&prod, orderItem.ProductID)

	if result.Error != nil {
		return 0, result.Error
	}

	hasQuantity, err := orderItem.Product.HasQuantity(orderItem.ProductID, orderItem.Quantity)
	if err != nil {
		return 0, err
	}

	if !hasQuantity {
		return 0, fmt.Errorf("недостаточное количество продукта %s. Доступно: %d, Запрашиваемое: %d", prod.Name, prod.Quantity, orderItem.Quantity)
	}

	err = DataSource.Context.Create(&orderItem).Error

	if err != nil {
		return 0, err
	}

	prod.Quantity -= orderItem.Quantity
	err = prod.UpdateProduct(prod.ID, *prod)
	if err != nil {
		return 0, err
	}

	return s.ID, nil
}

func (s *OrderItem) GetOrderItemByID(orderItemID uint) (OrderItem, error) {
	orderItem := &OrderItem{}
	err := DataSource.Context.Where("id = ?", orderItemID).First(orderItem).Error
	if err != nil {
		return *orderItem, err
	}
	return *orderItem, nil
}

func (s *OrderItem) UpdateOrderItem(orderItemID uint, orderItem OrderItem) error {
	err := DataSource.Context.Model(orderItem).Where("id = ?", orderItemID).Updates(s).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderItem) DeleteOrderItem(orderItemID uint) error {
	err := DataSource.Context.Where("id = ?", orderItemID).Delete(s).Error
	if err != nil {
		return err
	}
	return nil
}

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	ClientID   uint        `json:"client_id" binding:"required"`
	Client     User        `gorm:"foreignKey:ClientID" binding:"-"`
	Address    string      `json:"address" binding:"required"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" binding:"-"`
	OrderDate  time.Time   `json:"order_date" binding:"required"`
	Status     string      `json:"status" binding:"required"`
}

func (o *Order) CreateOrder(order Order) (uint, error) {
	err := DataSource.Context.Create(&order).Error
	if err != nil {
		return 0, err
	}
	return o.ID, nil
}

func (o *Order) GetOrderByID(orderID uint) (Order, error) {
	order := &Order{}
	err := DataSource.Context.Where("id = ?", orderID).
		Preload("Client.Role").
		Preload("Client.Position").
		Preload("OrderItems.Order").
		Preload("OrderItems.Product").
		First(order).Error
	if err != nil {
		return *order, err
	}
	return *order, nil
}

func (o *Order) GetDeliveryByOrder(orderID uint) (DeliveryRoute, error) {
	var route DeliveryRoute
	res := DataSource.Context.Model(&DeliveryRoute{}).Where("order_id = ?", orderID).
		Preload("Client.Role").
		Preload("Client.Position").
		Preload("OrderItems.Order").
		Preload("OrderItems.Product").First(&route)
	if res.Error != nil {
		return route, res.Error
	}
	return route, nil
}

func (o *Order) GetAllOrders() ([]Order, error) {
	var orders []Order

	err := DataSource.Context.Preload("Client.Role").
		Preload("Client.Role").
		Preload("Client.Position").
		Preload("OrderItems.Order").
		Preload("OrderItems.Product").
		Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (o *Order) UpdateOrder(orderID uint, order Order) error {
	err := DataSource.Context.Model(order).Where("id = ?", orderID).Updates(o).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) DeleteOrder(orderID uint) error {
	// Delete related order items
	var orderItems []OrderItem
	result := DataSource.Context.Where("order_id = ?", orderID).Find(&orderItems)
	if result.Error != nil {
		return result.Error
	}

	// Delete each order item
	for _, orderItem := range orderItems {
		if err := orderItem.DeleteOrderItem(orderItem.ID); err != nil {
			return err
		}
	}

	// Delete the order
	result = DataSource.Context.Where("id = ?", orderID).Delete(o)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

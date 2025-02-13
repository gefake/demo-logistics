package handler

import (
	"github.com/gin-gonic/gin"
	"logistic_api/pkg/service"
)

type Handler struct {
	Services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
	}

	api := router.Group("/api")
	{
		geoByAddress := api.Group("/geo")
		{
			geoByAddress.GET("/address", h.GetGeoByAddress)
		}
		// Группировка маршрутов для CargoRepository
		cargoRoutes := api.Group("/cargo")
		{
			cargoRoutes.POST("", h.CreateCargo)
			cargoRoutes.GET("/:id", h.GetCargoByID)
			cargoRoutes.PUT("/:id", h.UpdateCargo)
			cargoRoutes.DELETE("/:id", h.DeleteCargo)

			cargoUser := cargoRoutes.Group("/user")
			{
				cargoUser.GET("/:id", h.GetCargosByUserID)
			}
		}

		api.GET("/cargos", h.GetAllCargos)

		// Группировка маршрутов для UserRepository
		userRoutes := api.Group("/user")
		{
			userRoutes.POST("", h.CreateUser)
			userRoutes.GET("/:id", h.GetUserByID)
			userRoutes.PUT("/:id", h.UpdateUser)
			userRoutes.DELETE("/:id", h.DeleteUser)
		}

		// Группировка маршрутов для PositionRepository
		positionRoutes := api.Group("/position")
		{
			positionRoutes.POST("", h.CreatePosition)
			positionRoutes.GET("/:id", h.GetPositionByID)
			positionRoutes.PUT("/:id", h.UpdatePosition)
			positionRoutes.DELETE("/:id", h.DeletePosition)
		}

		// Группировка маршрутов для RoleRepository
		roleRoutes := api.Group("/role")
		{
			roleRoutes.POST("", h.CreateRole)
			roleRoutes.GET("/:id", h.GetRoleByID)
			roleRoutes.PUT("/:id", h.UpdateRole)
			roleRoutes.DELETE("/:id", h.DeleteRole)
		}

		// Группировка маршрутов для OrderRepository
		orderRoutes := api.Group("/order")
		{
			orderRoutes.POST("", h.CreateOrder)
			orderRoutes.GET("/:id", h.GetOrderByID)
			orderRoutes.PUT("/:id", h.UpdateOrder)
			orderRoutes.GET("/", h.GetAllOrders)
			orderRoutes.DELETE("/:id", h.DeleteOrder)

			orderProducts := orderRoutes.Group("/products")
			{
				orderProducts.POST("", h.CreateOrderProduct)
				orderProducts.DELETE("/:id", h.DeleteOrderProduct)
			}
		}

		// Группировка маршрутов для SupplierRepository
		supplierRoutes := api.Group("/supplier")
		{
			supplierRoutes.POST("", h.CreateSupplier)
			supplierRoutes.GET("/:id", h.GetSupplierByID)
			supplierRoutes.GET("/", h.GetAllSuppliers)
			supplierRoutes.PUT("/:id", h.UpdateSupplier)
			supplierRoutes.DELETE("/:id", h.DeleteSupplier)
		}

		reports := api.Group("/reports")
		{
			sales := reports.Group("sales")
			{
				sales.GET("", h.salesReport)
				sales.GET("excel", h.salesReportExcel)
			}

			delivery := reports.Group("delivery")
			{
				delivery.GET("", h.deliveryReport)
				delivery.GET("excel", h.deliveryReportExcel)
			}
		}

		// Группировка маршрутов для ProductRepository
		productRoutes := api.Group("/product")
		{
			productRoutes.POST("", h.CreateProduct)
			productRoutes.GET("/:id", h.GetProductByID)
			productRoutes.PUT("/:id", h.UpdateProduct)
			productRoutes.DELETE("/:id", h.DeleteProduct)
			productRoutes.GET("/cats", h.getCats)
			productRoutes.GET("/", h.getAllProducts)
		}

		// Группировка маршрутов для DeliveryRouteRepository
		deliveryRoutes := api.Group("/delivery")
		{
			deliveryRoutes.POST("", h.CreateDeliveryRoute)
			deliveryRoutes.GET("/:id", h.GetDeliveryRouteByID)
			deliveryRoutes.GET("", h.GetAllRoutes)
			deliveryRoutes.PUT("/:id", h.UpdateDeliveryRoute)
			deliveryRoutes.DELETE("/:id", h.DeleteDeliveryRoute)
		}

		warehouseRoutes := api.Group("/warehouse")
		{
			warehouseRoutes.POST("", h.createWarehouse)
			warehouseRoutes.GET("/:id", h.getWarehouseByID)
			warehouseRoutes.GET("", h.getAllWarehouses)
			warehouseRoutes.PUT("/:id", h.updateWarehouse)
			warehouseRoutes.DELETE("/:id", h.deleteWarehouse)
		}

		deliverySchedule := api.Group("/delivery-schedule")
		{
			deliverySchedule.POST("", h.createDeliverySchedule)
			deliverySchedule.GET("/:id", h.getDeliveryScheduleByID)
			deliverySchedule.GET("", h.getAllDeliverySchedules)
			deliverySchedule.PUT("/:id", h.updateDeliverySchedule)
			deliverySchedule.DELETE("/:id", h.deleteDeliverySchedule)
		}
	}

	return router
}

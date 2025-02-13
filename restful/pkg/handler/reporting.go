package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"logistic_api/pkg/database"
	"logistic_api/pkg/service"
	"math"
	"net/http"
	"os"
	"strconv"
)

func loadReportData(model any, date string, report service.Report, arr any, preloadFields ...string) error {
	db := database.DataSource.Context.Model(&model).Where(date+" BETWEEN ? AND ?", report.GetDateStart(), report.GetDateEnd())

	for _, field := range preloadFields {
		db = db.Preload(field)
	}

	err := db.Find(arr).Error

	if err != nil {
		return err
	}

	return nil
}

// salesReport godoc
// @Summary Get sales report
// @Description Get sales report
// @Tags Reporting
// @Accept json
// @Produce json
// @Param diagramReport body service.DiagramReport true "DiagramReport object"
// @Success 200 {object} service.DiagramReport
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/reports/sales [post]
func (h *Handler) salesReport(c *gin.Context) {
	var diagramReport service.DiagramReport
	var orders []database.Order
	var labelData map[string]float64 = make(map[string]float64)

	if err := c.ShouldBindJSON(&diagramReport); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	loadReportData(database.Order{}, "order_date", &diagramReport, &orders, "OrderItems.Product")

	for _, order := range orders {
		for _, product := range order.OrderItems {
			referItem := product.Product

			labelData[referItem.Name] += float64(product.Quantity) * referItem.Price
		}
	}

	for label, value := range labelData {
		diagramReport.DiagramData.Labels = append(diagramReport.DiagramData.Labels, label)
		diagramReport.DiagramData.Data = append(diagramReport.DiagramData.Data, math.Floor(value))
	}

	c.JSON(http.StatusOK, diagramReport)
}

type tableReport struct {
	sum   float64
	count int
}

// salesReportExcel godoc
// @Summary Get sales report
// @Description Get sales report
// @Tags Reporting
// @Accept json
// @Produce json
// @Param diagramReport body service.TableReport true "TableReport object"
// @Success 200 {object} service.TableReport
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/reports/sales/excel [post]
func (h *Handler) salesReportExcel(c *gin.Context) {
	var rows [][]string
	var report service.TableReport
	var orders []database.Order
	var labelData map[string]tableReport = make(map[string]tableReport)

	if err := c.ShouldBindJSON(&report); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	report.UniqueID = "sales"

	loadReportData(database.Order{}, "order_date", &report, &orders, "OrderItems.Product")

	for _, order := range orders {
		for _, product := range order.OrderItems {
			referItem := product.Product

			// Update the sum value for the corresponding label in labelData
			currentReport := labelData[referItem.Name]
			currentReport.sum += float64(product.Quantity) * referItem.Price
			currentReport.count += product.Quantity
			labelData[referItem.Name] = currentReport
		}
	}

	rows = append(rows, []string{"Наименование", "Сумма продаж", "Количество проданных товаров"})

	for itemName, value := range labelData {
		sum := strconv.FormatFloat(math.Floor(value.sum), 'f', -1, 64)
		count := strconv.Itoa(value.count)
		rows = append(rows, []string{itemName, sum, count})
	}

	report.TableData.Rows = rows

	path, err := report.ToExcel()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := os.Open(path)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	// Установка заголовков ответа для скачивания файла
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename"+path)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// Передача файла клиенту
	c.Writer.WriteHeader(http.StatusOK)
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func contains(s []string, e any) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// deliveryReport godoc
// @Summary Get delivery report
// @Description Get delivery report
// @Tags Reporting
// @Accept json
// @Produce json
// @Param diagramReport body service.DiagramReport true "DiagramReport object"
// @Success 200 {object} service.DiagramReport
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/reports/delivery [post]
func (h *Handler) deliveryReport(c *gin.Context) {
	var deliveries []database.DeliveryRoute
	var data = make(map[string]int)
	var diagramReport service.DiagramReport

	if err := c.ShouldBindJSON(&diagramReport); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := loadReportData(database.DeliveryRoute{}, "arrival_date", &diagramReport, &deliveries, "Cargo.Order.OrderItems.Product")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for _, delivery := range deliveries {
		data[delivery.Status]++
	}

	for k, v := range data {
		diagramReport.DiagramData.Labels = append(diagramReport.DiagramData.Labels, k)
		diagramReport.DiagramData.Data = append(diagramReport.DiagramData.Data, float64(v))
	}

	c.JSON(http.StatusOK, diagramReport)
}

type deliveryReport struct {
	id            int
	status        string
	receiver      string
	driver        string
	address       string
	weight        float64
	arrivalDate   string
	departureDate string
}

// deliveryReportExcel godoc
// @Summary Get delivery report
// @Description Get delivery report
// @Tags Reporting
// @Accept json
// @Produce json
// @Param diagramReport body service.TableReport true "TableReport object"
// @Success 200 {object} service.TableReport
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/reports/delivery/excel [post]
func (h *Handler) deliveryReportExcel(c *gin.Context) {
	var rows [][]string
	var report service.TableReport
	var deliveries []database.DeliveryRoute
	var data = make(map[uint]deliveryReport)

	if err := c.ShouldBindJSON(&report); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	report.UniqueID = "delivery"

	err := loadReportData(database.DeliveryRoute{}, "arrival_date", &report, &deliveries, "Cargo.Order.Client", "Cargo.Order.OrderItems.Product", "Driver")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for _, delivery := range deliveries {
		if delivery.Status != "Доставлен" && delivery.Status != "Доставлено" && delivery.Status != "Завершено" {
			continue
		}

		curDelivery := data[delivery.ID]
		curDelivery.id = int(delivery.ID)
		curDelivery.address = delivery.Cargo.Order.Address
		curDelivery.arrivalDate = delivery.ArrivalDate.Format("02-01-2006")
		curDelivery.departureDate = delivery.DepartureDate.Format("02-01-2006")
		curDelivery.driver = fmt.Sprintf("%s %s", delivery.Driver.Firstname, delivery.Driver.Lastname)
		curDelivery.receiver = fmt.Sprintf("%s %s", delivery.Cargo.Order.Client.Firstname, delivery.Cargo.Order.Client.Lastname)
		curDelivery.weight = delivery.Cargo.Weight
		curDelivery.status = delivery.Status
		data[delivery.ID] = curDelivery
	}

	rows = append(rows, []string{"Наименование", "Статус", "Получатель", "Водитель", "Адрес", "Вес груза", "Дата прибытия", "Дата отправки"})

	for _, v := range data {
		rows = append(rows, []string{"Доставка #" + strconv.Itoa(v.id), v.status, v.receiver, v.driver, v.address, strconv.FormatFloat(v.weight, 'f', 2, 64), v.arrivalDate, v.departureDate})
	}

	report.TableData.Rows = rows

	path, err := report.ToExcel()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := os.Open(path)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename"+path)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	c.Writer.WriteHeader(http.StatusOK)
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

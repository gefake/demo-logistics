package handler

import (
	"context"
	"github.com/ekomobile/dadata/v2"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Metro struct {
	Name     string  `json:"name"`
	Line     string  `json:"line"`
	Distance float64 `json:"distance"`
}

// Address base struct for dadata.Address
type Address struct {
	Source               string `json:"source"`                  // Исходный адрес одной строкой
	Result               string `json:"result"`                  // Стандартизованный адрес одной строкой
	PostalCode           string `json:"postal_code"`             // Индекс
	Country              string `json:"country"`                 // Страна
	CountryIsoCode       string `json:"country_iso_code"`        // ISO-код страны
	FederalDistrict      string `json:"federal_district"`        // Федеральный округ
	RegionFiasID         string `json:"region_fias_id"`          // Код ФИАС региона
	RegionKladrID        string `json:"region_kladr_id"`         // Код КЛАДР региона
	RegionIsoCode        string `json:"region_iso_code"`         // ISO-код региона
	RegionWithType       string `json:"region_with_type"`        // Регион с типом
	RegionType           string `json:"region_type"`             // Тип региона (сокращенный)
	RegionTypeFull       string `json:"region_type_full"`        // Тип региона
	Region               string `json:"region"`                  // Регион
	AreaFiasID           string `json:"area_fias_id"`            // Код ФИАС района в регионе
	AreaKladrID          string `json:"area_kladr_id"`           // Код КЛАДР района в регионе
	AreaWithType         string `json:"area_with_type"`          // Район в регионе с типом
	AreaType             string `json:"area_type"`               // Тип района в регионе (сокращенный)
	AreaTypeFull         string `json:"area_type_full"`          // Тип района в регионе
	Area                 string `json:"area"`                    // Район в регионе
	CityFiasID           string `json:"city_fias_id"`            // Код ФИАС города
	CityKladrID          string `json:"city_kladr_id"`           // Код КЛАДР города
	CityWithType         string `json:"city_with_type"`          // Город с типом
	CityType             string `json:"city_type"`               // Тип города (сокращенный)
	CityTypeFull         string `json:"city_type_full"`          // Тип города
	City                 string `json:"city"`                    // Город
	CityArea             string `json:"city_area"`               // Административный округ (только для Москвы)
	CityDistrictFiasID   string `json:"city_district_fias_id"`   // Код ФИАС района города (заполняется, только если район есть в ФИАС)
	CityDistrictKladrID  string `json:"city_district_kladr_id"`  // Код КЛАДР района города (не заполняется)
	CityDistrictWithType string `json:"city_district_with_type"` // Район города с типом
	CityDistrictType     string `json:"city_district_type"`      // Тип района города (сокращенный)
	CityDistrictTypeFull string `json:"city_district_type_full"` // Тип района города
	CityDistrict         string `json:"city_district"`           // Район города
	SettlementFiasID     string `json:"settlement_fias_id"`      // Код ФИАС нас. пункта
	SettlementKladrID    string `json:"settlement_kladr_id"`     // Код КЛАДР нас. пункта
	SettlementWithType   string `json:"settlement_with_type"`    // Населенный пункт с типом
	SettlementType       string `json:"settlement_type"`         // Тип населенного пункта (сокращенный)
	SettlementTypeFull   string `json:"settlement_type_full"`    // Тип населенного пункта
	Settlement           string `json:"settlement"`              // Населенный пункт
	StreetFiasID         string `json:"street_fias_id"`          // Код ФИАС улицы
	StreetKladrID        string `json:"street_kladr_id"`         // Код КЛАДР улицы
	StreetWithType       string `json:"street_with_type"`        // Улица с типом
	StreetType           string `json:"street_type"`             // Тип улицы (сокращенный)
	StreetTypeFull       string `json:"street_type_full"`        // Тип улицы
	Street               string `json:"street"`                  // Улица
	HouseFiasID          string `json:"house_fias_id"`           // Код ФИАС дома
	HouseKladrID         string `json:"house_kladr_id"`          // Код КЛАДР дома
	HouseType            string `json:"house_type"`              // Тип дома (сокращенный)
	HouseTypeFull        string `json:"house_type_full"`         // Тип дома
	House                string `json:"house"`                   // Дом
	HouseCadNum          string `json:"house_cadnum"`            // Кадастровый номер дома (22.4+). Заполняется в зависимости от тарифа «Дадаты».
	BlockType            string `json:"block_type"`              // Тип корпуса/строения (сокращенный)
	BlockTypeFull        string `json:"block_type_full"`         // Тип корпуса/строения
	Block                string `json:"block"`                   // Корпус/строение
	Entrance             string `json:"entrance"`                // Подъезд
	Floor                string `json:"floor"`                   // Этаж
	FlatFiasId           string `json:"flat_fias_id"`            // ФИАС-код квартиры
	FlatType             string `json:"flat_type"`               // Тип квартиры (сокращенный)
	FlatTypeFull         string `json:"flat_type_full"`          // Тип квартиры
	Flat                 string `json:"flat"`                    // Квартира
	FlatArea             string `json:"flat_area"`               // Площадь квартиры
	FlatCadNum           string `json:"flat_cadnum"`             // Кадастровый номер квартиры (22.4+). Заполняется в зависимости от тарифа «Дадаты».
	SquareMeterPrice     string `json:"square_meter_price"`      // Рыночная стоимость м²
	FlatPrice            string `json:"flat_price"`              // Рыночная стоимость квартиры
	PostalBox            string `json:"postal_box"`              // Абонентский ящик
	FiasID               string `json:"fias_id"`                 // Код ФИАС
	FiasCode             string `json:"fias_code"`               // Иерархический код адреса в ФИАС (СС+РРР+ГГГ+ППП+СССС+УУУУ+ДДДД)
	FiasLevel            string `json:"fias_level"`              // Уровень детализации, до которого адрес найден в ФИАС
	FiasActualityState   string `json:"fias_actuality_state"`    // Признак актуальности адреса в ФИАС
	KladrID              string `json:"kladr_id"`                // Код КЛАДР
	CapitalMarker        string `json:"capital_marker"`          // Статус центра
	Okato                string `json:"okato"`                   // Код ОКАТО
	Oktmo                string `json:"oktmo"`                   // Код ОКТМО
	TaxOffice            string `json:"tax_office"`              // Код ИФНС для физических лиц
	TaxOfficeLegal       string `json:"tax_office_legal"`        // Код ИФНС для организаций
	Timezone             string `json:"timezone"`                // Часовой пояс
	GeoLat               string `json:"geo_lat"`                 // Координаты: широта
	GeoLon               string `json:"geo_lon"`                 // Координаты: долгота
	BeltwayHit           string `json:"beltway_hit"`             // Внутри кольцевой?
	BeltwayDistance      string `json:"beltway_distance"`        // Расстояние от кольцевой в км.
	// QualityCodeGeoRaw для clean вызовов он int для suggest в адресе банков он string поэтому в поле поставил interface{} чтобы работало и там и там)\
	QualityCodeGeoRaw      interface{} `json:"qc_geo"`         // Код точности координат
	QualityCodeCompleteRaw interface{} `json:"qc_complete"`    // Код полноты
	QualityCodeHouseRaw    interface{} `json:"qc_house"`       // Код проверки дома
	QualityCodeRaw         interface{} `json:"qc"`             // Код качества
	UnparsedParts          string      `json:"unparsed_parts"` // Нераспознанная часть адреса. Для адреса
	Metro                  []*Metro    `json:"metro"`
}

// GetGeoByAddress retrieves geolocation data for a given address. godoc
//
// @Summary Get geolocation data for an address
// @Description Retrieves geolocation data for a given address using the DaData API.
// @Tags Geolocation
// @Accept json
// @Produce json
// @Param address body string true "Address"
// @Success 200 array Address "Geolocation data retrieved"
// @Failure 400 {object} errorResponse "Invalid request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/geo/address [post]
func (h *Handler) GetGeoByAddress(c *gin.Context) {
	api := dadata.NewCleanApi()

	body, err := c.Request.GetBody()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	address, err := io.ReadAll(body)
	query := string(address)

	ctx := context.WithValue(context.Background(), "Content-Type", "application/json")
	ctx = context.WithValue(ctx, "Accept", "application/json")
	ctx = context.WithValue(ctx, "Authorization:", "Token a2b4418c7d1219b92285a2ea60f0767080bc6a9b")
	ctx = context.WithValue(ctx, "X-Secret:", "a8d3b7e116b99bfd4c3bde8d9368c9fc438715db")

	result, err := api.Address(c, query)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

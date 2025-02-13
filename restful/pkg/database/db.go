package database

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"logistic_api/pkg/logger"
	"os"
)

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (p Point) Value() (driver.Value, error) {
	return fmt.Sprintf("(%f,%f)", p.Lat, p.Lon), nil
}

func (p *Point) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("Scan source is not a string")
	}

	_, err := fmt.Sscanf(str, "(%f,%f)", &p.Lat, &p.Lon)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	logger.Log.Info("Connecting to database")

	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal(".env файл не существует. Дальнешее подключение к БД невозможно осуществить")
	}

	dbPort, _ := os.LookupEnv("DB_PORT")
	dbName, _ := os.LookupEnv("DB_NAME")
	dbUser, _ := os.LookupEnv("DB_USER")
	dbPass, _ := os.LookupEnv("DB_PASSWORD")

	dataSetName := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbUser, dbPass, dbName, dbPort)

	//logger.Log.Info("Connect with using dataset" + dataSetName)

	if db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dataSetName,
		PreferSimpleProtocol: true,
	}), &gorm.Config{}); err != nil {
		logger.Log.Fatal("Failed to connect to database " + err.Error())
		return
	} else {

		// db.Exec("CREATE DATABASE logistics_api")

		// fmt.Println("База данных 'logistics_api' успешно создана")

		logger.Log.Info("Connected to database")

		_ = db.AutoMigrate(&Warehouse{}, &Position{}, &User{}, &Order{}, &Supplier{}, &Product{}, &Cargo{}, &DeliveryRoute{}, &CargoProduct{}, &OrderItem{}, &DeliverySchedule{}, &DeliveryScheduleProduct{})
		//_ = db.AutoMigrate(&Warehouse{}, &DeliverySchedule{})

		DataSource = &Connect{
			Context: db,
		}

		err := DataSource.createDefaultRolesIfNotExist()
		if err != nil {
			logger.Log.Warning("Failed to create default roles " + err.Error())
		}

		err = DataSource.createDefaultPositionsIfNotExist()
		if err != nil {
			logger.Log.Warning("Failed to create default position " + err.Error())
		}
	}
}

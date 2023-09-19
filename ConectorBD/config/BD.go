package config

import (
	ModelosVehiculos "example/fleetwise/modelos/vehiculos"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3307)/gestor_de_flotillas")

	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	database.AutoMigrate(&ModelosVehiculos.Vehiculo{})

	DB = database
}

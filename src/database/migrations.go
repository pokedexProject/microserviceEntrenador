package database

import (
	"log"

	model "github.com/pokedexProject/microserviceEntrenador/dominio"
	"gorm.io/gorm"
)

// EjecutarMigraciones realiza todas las migraciones necesarias en la base de datos.
func EjecutarMigraciones(db *gorm.DB) {

	db.AutoMigrate(&model.EntrenadorGORM{})

	log.Println("Migraciones completadas")
}

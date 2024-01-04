package dominio

import (
	"strconv"
)

// UsuarioGORM es el modelo de usuario para GORM de Usuario
type EntrenadorGORM struct {
	ID          uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Nombre      string `gorm:"type:varchar(255);not null"`
	Correo      string `gorm:"type:varchar(255);not null;unique"`
	Contrasena  string `gorm:"type:varchar(255);not null"`
	Nivel       string `gorm:"type:varchar(255);not null"`
	IDCompanero string `gorm:"type:varchar(255);not null"`
}

// TableName especifica el nombre de la tabla para UsuarioGORM
func (EntrenadorGORM) TableName() string {
	return "entrenadores"
}

func (entrenadorGORM *EntrenadorGORM) ToGQL() (*Entrenador, error) {

	return &Entrenador{
		ID:          strconv.Itoa(int(entrenadorGORM.ID)),
		Nombre:      entrenadorGORM.Nombre,
		Correo:      entrenadorGORM.Correo,
		Contrasena:  entrenadorGORM.Contrasena,
		Nivel:       entrenadorGORM.Nivel,
		IDCompanero: entrenadorGORM.IDCompanero,
	}, nil
}

package ports

import (
	model "github.com/pokedexProject/microserviceEntrenador/dominio"
)

// puerto de salida
type MasterRepository interface {
	CrearEntrenador(input model.CrearEntrenadorInput) (*model.Entrenador, error)
	Entrenador(id string) (*model.Entrenador, error)
	ActualizarEntrenador(id string, input *model.ActualizarEntrenadorInput) (*model.Entrenador, error)
	EliminarEntrenador(id string) (*model.EliminacionEntrenador, error)
	Entrenadores() ([]*model.Entrenador, error)
	ExistePorCorreo(correo string) (bool, error)
	Retrieve(correo string, contrasena string) (*model.Entrenador, error)
	Login(input model.LoginInput) (*model.AuthPayload, error)
	Logout(id string) (model.EliminacionEntrenador, error)
}

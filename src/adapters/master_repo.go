package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pokedexProject/microserviceEntrenador/database"
	model "github.com/pokedexProject/microserviceEntrenador/dominio"
	"github.com/pokedexProject/microserviceEntrenador/ports"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/**
* Es un adaptador de salida
usuario
*/

type masterRepository struct {
	db             *database.DB
	activeSessions map[string]string
}

func NewMasterRepository(db *database.DB) ports.MasterRepository {
	return &masterRepository{
		db:             db,
		activeSessions: make(map[string]string),
	}
}

func ToJSON(obj interface{}) (string, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

// ExistePorCorreo verifica si existe un usuario con el correo proporcionado.
func (ur *masterRepository) ExistePorCorreo(correo string) (bool, error) {
	var entrenadorGORM model.EntrenadorGORM
	result := ur.db.GetConn().Where("correo = ?", correo).First(&entrenadorGORM)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		log.Printf("Error al buscar el usuario con correo %s: %v", correo, result.Error)
		return false, result.Error
	}

	return true, result.Error
}

// Retrieve obtiene un usuario por su correo y contraseña.
// Retorna nil si no se encuentra el usuario.
func (ur *masterRepository) Retrieve(correo string, contrasena string) (*model.Entrenador, error) {
	var entrenadorGORM model.EntrenadorGORM
	fmt.Printf("correo: %s\n", correo)

	if err := ur.db.GetConn().Where("correo = ?", correo).First(&entrenadorGORM).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Usuario con correo %s no encontrado", correo)
		}
		return nil, fmt.Errorf("Error al buscar entrenador: %v", err)
	}

	// Verificar la contraseña con bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(entrenadorGORM.Contrasena), []byte(contrasena)); err != nil {
		// Contraseña incorrecta
		return nil, fmt.Errorf("Credenciales incorrectas")
	}
	return entrenadorGORM.ToGQL()
}

// ObtenerTrabajo obtiene un trabajo por su ID.
func (ur *masterRepository) Entrenador(id string) (*model.Entrenador, error) {
	if id == "" {
		return nil, errors.New("El ID de entrenador es requerido")
	}

	var entrenadorGORM model.EntrenadorGORM
	//result := ur.db.GetConn().First(&usuarioGORM, id)
	result := ur.db.GetConn().First(&entrenadorGORM, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Printf("Error al obtener el entrenador con ID %s: %v", id, result.Error)
		return nil, result.Error
	}

	return entrenadorGORM.ToGQL()
}

// Entrenadores obtiene todos los usuarios de la base de datos.
func (ur *masterRepository) Entrenadores() ([]*model.Entrenador, error) {
	var entrenadoresGORM []model.EntrenadorGORM
	result := ur.db.GetConn().Find(&entrenadoresGORM)

	if result.Error != nil {
		log.Printf("Error al obtener los entrenadores: %v", result.Error)
		return nil, result.Error
	}

	var entrenadores []*model.Entrenador
	for _, entrenadorGORM := range entrenadoresGORM {
		entrenador, _ := entrenadorGORM.ToGQL()
		entrenadores = append(entrenadores, entrenador)
	}

	// usuariosJSON, err := json.Marshal(usuarios)
	// if err != nil {
	// 	log.Printf("Error al convertir usuarios a JSON: %v", err)
	// 	return "[]", err
	// }
	// return ToJSON(usuarios)
	return entrenadores, nil
}
func (ur *masterRepository) CrearEntrenador(input model.CrearEntrenadorInput) (*model.Entrenador, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Contrasena), bcrypt.DefaultCost)
	log.Printf("Hashed password: %s", string(hashedPassword))

	if err != nil {
		log.Printf("Error al crear el hash de la contraseña: %v", err)
		return nil, err
	}

	entrenadorGORM :=
		&model.EntrenadorGORM{
			Nombre:      input.Nombre,
			Correo:      input.Correo,
			Contrasena:  string(hashedPassword),
			Nivel:       input.Nivel,
			IDCompanero: input.IDCompanero,
		}
	result := ur.db.GetConn().Create(&entrenadorGORM)
	if result.Error != nil {
		log.Printf("Error al crear el entrenador: %v", result.Error)
		return nil, result.Error
	}

	response, err := entrenadorGORM.ToGQL()
	return response, err
}

func (ur *masterRepository) ActualizarEntrenador(id string, input *model.ActualizarEntrenadorInput) (*model.Entrenador, error) {
	var entrenadorGORM model.EntrenadorGORM
	if id == "" {
		return nil, errors.New("El ID de entrenador es requerido")
	}

	result := ur.db.GetConn().First(&entrenadorGORM, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Entrenador con ID %s no encontrado", id)
		}
		return nil, result.Error
	}

	// Solo actualiza los campos proporcionados
	if input.Nombre != nil {
		entrenadorGORM.Nombre = *input.Nombre
	}
	if input.Correo != nil {
		entrenadorGORM.Correo = *input.Correo
	}
	if input.Nivel != nil {
		entrenadorGORM.Nivel = *input.Nivel
	}
	if input.IDCompanero != nil {
		entrenadorGORM.IDCompanero = *input.IDCompanero
	}

	result = ur.db.GetConn().Save(&entrenadorGORM)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Printf("Entrenador actualizado: %v", entrenadorGORM)
	return entrenadorGORM.ToGQL()
}

// EliminarEntrenador elimina un entrenador de la base de datos por su ID.
func (ur *masterRepository) EliminarEntrenador(id string) (*model.EliminacionEntrenador, error) {
	// Intenta buscar el entrenador por su ID
	var entrenadorGORM model.EntrenadorGORM
	result := ur.db.GetConn().First(&entrenadorGORM, id)

	if result.Error != nil {
		// Manejo de errores
		if result.Error == gorm.ErrRecordNotFound {
			// El entrenador no se encontró en la base de datos
			response := &model.EliminacionEntrenador{
				Mensaje: "El entrenador no existe",
			}
			return response, result.Error

		}
		log.Printf("Error al buscar el entrenador con ID %s: %v", id, result.Error)
		response := &model.EliminacionEntrenador{
			Mensaje: "Error al buscar el entrenador",
		}
		return response, result.Error
	}

	// Elimina el entrenador de la base de datos
	result = ur.db.GetConn().Delete(&entrenadorGORM, id)

	if result.Error != nil {
		log.Printf("Error al eliminar el entrenador con ID %s: %v", id, result.Error)
		response := &model.EliminacionEntrenador{
			Mensaje: "Error al eliminar el entrenador",
		}
		return response, result.Error
	}

	// Éxito al eliminar el entrenador
	response := &model.EliminacionEntrenador{
		Mensaje: "Entrenador eliminado con éxito",
	}
	return response, result.Error

}

func (ur *masterRepository) Login(input model.LoginInput) (*model.AuthPayload, error) {
	// Verificar las credenciales del entrenador (correo y contraseña)
	if input.Correo == "" || input.Contrasena == "" {
		return nil, errors.New("Correo y contraseña son requeridos")
	}
	// if len(input.Contrasena) < 6 || len(input.Contrasena) > 50 {
	// 	return nil, errors.New("La contraseña debe tener al menos 6 caracteres")
	// }
	// if len(input.Correo) < 3 || len(input.Correo) > 50 {
	// 	return nil, errors.New("El correo debe tener al menos 3 caracteres")
	// }

	entrenador, err := ur.Retrieve(input.Correo, input.Contrasena)
	if err != nil {
		fmt.Printf("Error al verificar las credenciales: %v", err)
		return nil, errors.New("Credenciales inválidas")
	}

	// Comprueba si el Entrenador ya tiene una sesión activa (esto podría ser a través de una base de datos)
	if ur.isSessionActive(entrenador.ID) {
		return nil, errors.New("Ya existe una sesión activa")
	}
	// Generar un token de autenticación para el entrenador
	token, err := CreateToken(entrenador)
	if err != nil {
		fmt.Printf("Error al generar el token de autenticación: %v", err)
		return nil, fmt.Errorf("Error al generar el token: %v", err)
	}

	ur.registerSession(entrenador.ID, token)

	// Crear el objeto AuthPayload con el token y los datos del entrenador
	authPayload := &model.AuthPayload{
		Token:      token,
		Entrenador: entrenador,
	}
	log.Printf("Entrenador autenticado: %v", entrenador.ID)
	return authPayload, nil

}

func (ur *masterRepository) Logout(masterID string) (model.EliminacionEntrenador, error) {
	var respuesta model.EliminacionEntrenador
	if masterID == "" {
		return model.EliminacionEntrenador{
			Mensaje: "El ID de entrenador es requerido",
		}, errors.New("El ID de entrenador es requerido")
	}
	if !ur.isSessionActive(masterID) {
		return model.EliminacionEntrenador{
			Mensaje: "No hay una sesión activa para este entrenador",
		}, errors.New("No hay una sesión activa para este entrenador")
	}
	delete(activeSessions, masterID)
	log.Printf("Sesión cerrada para el entrenador: %v", masterID)
	respuesta = model.EliminacionEntrenador{
		Mensaje: "Sesión cerrada exitosamente",
	}
	return respuesta, nil
}

// Clave secreta que no se expone! es una clvve
// del servidor
var jwtKey = []byte("clave_secreta")

// Estructura del token
type Claims struct {
	EntrenadorID string `json:"entrenador_id"`
	//Role   string `json:"role"`
	jwt.StandardClaims
}

func CreateToken(entrenador *model.Entrenador) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		EntrenadorID: entrenador.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ExtraerInfoToken es una función que decodifica un token JWT y extrae los claims (afirmaciones) del mismo.
func ExtraerInfoToken(tokenStr string) (*Claims, error) {
	// jwt.ParseWithClaims intenta analizar el token JWT.
	// Se le pasa el token como string, una instancia de Claims para mapear los datos del token,
	// y una función de callback para validar el algoritmo de firma del token.
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Comprueba que el algoritmo de codificación del token sea el esperado.
		// En este caso, se espera que el algoritmo sea HMAC (jwt.SigningMethodHMAC).
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Si el algoritmo es el esperado, se devuelve la clave secreta utilizada para firmar el token.
		return jwtKey, nil
	})

	// Si no hay errores y el token es válido, extrae los claims y los devuelve.
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		// Si hay un error o el token no es válido, devuelve el error.
		return nil, err
	}
}

var activeSessions = make(map[string]string) // Mapa de ID de usuario a token

func (ur *masterRepository) isSessionActive(userID string) bool {
	_, active := activeSessions[userID]
	return active
}
func (ur *masterRepository) registerSession(userID, token string) {
	activeSessions[userID] = token
}

func (ur *masterRepository) endSession(userID string) {
	delete(activeSessions, userID)
}

package dominio

type ActualizarEntrenadorInput struct {
	Nombre      *string `json:"nombre,omitempty"`
	Correo      *string `json:"correo,omitempty"`
	Contrasena  *string `json:"contrasena,omitempty"`
	Nivel       *string `json:"nivel,omitempty"`
	IDCompanero *string `json:"idCompanero,omitempty"`
}

type AuthPayload struct {
	Token      string      `json:"token"`
	Entrenador *Entrenador `json:"entrenador"`
}

type CrearEntrenadorInput struct {
	Nombre      string `json:"nombre"`
	Correo      string `json:"correo"`
	Contrasena  string `json:"contrasena"`
	Nivel       string `json:"nivel"`
	IDCompanero string `json:"idCompanero"`
}

type LoginInput struct {
	Correo     string `json:"correo"`
	Contrasena string `json:"contrasena"`
}

type EliminacionEntrenador struct {
	Mensaje string `json:"mensaje"`
}

type Entrenador struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Correo      string `json:"correo"`
	Contrasena  string `json:"contrasena"`
	Nivel       string `json:"nivel"`
	IDCompanero string `json:"idCompanero"`
}

func (Entrenador) IsEntity() {}

package service

import (
	"context"
	"fmt"

	model "github.com/pokedexProject/microserviceEntrenador/dominio"
	repository "github.com/pokedexProject/microserviceEntrenador/ports"
	pb "github.com/pokedexProject/microserviceEntrenador/proto"
)

// este servicio implementa la interfaz MasterServiceServer
// que se genera a partir del archivo proto
type MasterService struct {
	pb.UnimplementedMasterServiceServer
	repo repository.MasterRepository
}

func NewMasterService(repo repository.MasterRepository) *MasterService {
	return &MasterService{repo: repo}
}

func (s *MasterService) CreateMaster(ctx context.Context, req *pb.CreateMasterRequest) (*pb.CreateMasterResponse, error) {

	crearEntrenadorInput := model.CrearEntrenadorInput{
		Nombre:      req.GetNombre(),
		Correo:      req.GetCorreo(),
		Contrasena:  req.GetContrasena(),
		Nivel:       req.GetNivel(),
		IDCompanero: req.GetIdCompanero(),
	}
	u, err := s.repo.CrearEntrenador(crearEntrenadorInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Entrenador creado: %v", u)
	response := &pb.CreateMasterResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		Correo:      u.Correo,
		Nivel:       u.Nivel,
		IdCompanero: u.IDCompanero,
	}
	fmt.Printf("Entrenador creado: %v", response)
	return response, nil
}

func (s *MasterService) GetMaster(ctx context.Context, req *pb.GetMasterRequest) (*pb.GetMasterResponse, error) {
	u, err := s.repo.Entrenador(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.GetMasterResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		Correo:      u.Correo,
		Nivel:       u.Nivel,
		IdCompanero: u.IDCompanero,
	}
	return response, nil
}

func (s *MasterService) ListMasters(ctx context.Context, req *pb.ListMastersRequest) (*pb.ListMastersResponse, error) {
	masters, err := s.repo.Entrenadores()
	if err != nil {
		return nil, err
	}
	var response []*pb.Master
	for _, u := range masters {
		master := &pb.Master{
			Id:          u.ID,
			Nombre:      u.Nombre,
			Correo:      u.Correo,
			Nivel:       u.Nivel,
			IdCompanero: u.IDCompanero,
		}
		response = append(response, master)
	}

	return &pb.ListMastersResponse{Masters: response}, nil
}

func (s *MasterService) UpdateUser(ctx context.Context, req *pb.UpdateMasterRequest) (*pb.UpdateMasterResponse, error) {
	nombre := req.GetNombre()
	correo := req.GetCorreo()
	nivel := req.GetNivel()
	idCompanero := req.GetIdCompanero()
	fmt.Printf("Nombre: %v", nombre)
	actualizarEntrenadorInput := &model.ActualizarEntrenadorInput{
		Nombre:      &nombre,
		Correo:      &correo,
		Nivel:       &nivel,
		IDCompanero: &idCompanero,
	}
	fmt.Printf("Entrenador actualizado input: %v", actualizarEntrenadorInput)
	u, err := s.repo.ActualizarEntrenador(req.GetId(), actualizarEntrenadorInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Entrenador actualizado: %v", u)
	response := &pb.UpdateMasterResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		Correo:      u.Correo,
		Nivel:       u.Nivel,
		IdCompanero: u.IDCompanero,
	}
	return response, nil
}

func (s *MasterService) DeleteMaster(ctx context.Context, req *pb.DeleteMasterRequest) (*pb.DeleteMasterResponse, error) {
	respuesta, err := s.repo.EliminarEntrenador(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.DeleteMasterResponse{
		Mensaje: respuesta.Mensaje,
	}
	return response, nil
}

func (s *MasterService) LoginMaster(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	loginInput := model.LoginInput{
		Correo:     req.GetCorreo(),
		Contrasena: req.GetContrasena(),
	}
	authPayload, err := s.repo.Login(loginInput)
	if err != nil {
		return nil, err
	}
	response := &pb.LoginResponse{
		Token: authPayload.Token,
		Master: &pb.Master{
			Id:          authPayload.Entrenador.ID,
			Nombre:      authPayload.Entrenador.Nombre,
			Correo:      authPayload.Entrenador.Correo,
			Nivel:       authPayload.Entrenador.Nivel,
			IdCompanero: authPayload.Entrenador.IDCompanero,
		},
	}
	return response, nil
}

func (s *MasterService) LogoutMaster(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	respuesta, err := s.repo.Logout(req.GetMasterID())
	if err != nil {
		return nil, err
	}
	response := &pb.LogoutResponse{
		Mensaje: respuesta.Mensaje,
	}
	return response, nil
}

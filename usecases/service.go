package ordem

import (
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/vanigabriel/OrdemServico-Tasy/entity"
)

type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//InsertOS insere OS
func (s *Service) InsertOS(os *entity.OrdemServico) (string, error) {
	_, err := govalidator.ValidateStruct(os)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return s.repo.InsertOS(os)
}

func (s *Service) InsertAnexos(ordem string, filename string, file []byte) error {
	return s.repo.InsertAnexos(ordem, filename, file)
}
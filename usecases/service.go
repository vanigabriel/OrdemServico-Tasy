package ordem

import (
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/vanigabriel/OrdemServico-Tasy/entity"
)

type Service struct {
	repo Repository
}

func (s *Service) InsertOS(os *entity.OrdemServico) error {
	_, err := govalidator.ValidateStruct(os)
	if err != nil {
		log.Println(err)
		return err
	}

	return s.repo.InsertOS(os)
}

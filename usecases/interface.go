package ordem

import "github.com/vanigabriel/OrdemServico-Tasy/entity"

type Repository interface {
	InsertOS(os *entity.OrdemServico) (string, error)
	InsertAnexos(ordem string, filename string, file []byte) error
}

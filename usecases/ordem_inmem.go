package ordem

import (
	"github.com/vanigabriel/OrdemServico-Tasy/entity"
)

// RepoInMem define struct
type RepoInMem struct {
	inMem map[string]*entity.OrdemServico
}

//NewInmemRepository create new repository
func NewInmemRepository() *RepoInMem {
	var m = map[string]*entity.OrdemServico{}
	return &RepoInMem{
		inMem: m,
	}
}

func (r *RepoInMem) InsertOS(os *entity.OrdemServico) (string, error) {
	r.inMem[os.NrCPF] = os
	return os.NrCPF, nil
}

func (r *RepoInMem) InsertAnexos(ordem string, filename string, file []byte) error {
	return nil
}

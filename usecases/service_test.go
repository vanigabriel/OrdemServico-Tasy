package ordem

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanigabriel/OrdemServico-Tasy/entity"
)

func TestInsertOS(t *testing.T) {
	repo := NewInmemRepository()
	s := NewService(repo)

	// Insert Correto
	os := &entity.OrdemServico{
		NrCPF:     "12345678912",
		Descricao: "Teste",
		Contato:   "7744",
	}

	err := s.InsertOS(os)

	assert.Nil(t, err)

	// Insert incorreto, sem os campos obrigatórios
	os2 := &entity.OrdemServico{
		NrCPF: "12345678912",
	}

	err = s.InsertOS(os2)

	assert.NotNil(t, err)
}

package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/vanigabriel/OrdemServico-Tasy/entity"
	ordem "github.com/vanigabriel/OrdemServico-Tasy/usecases"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPostOS(t *testing.T) {
	// Prepara para testar em memória
	repo := ordem.NewInmemRepository()
	s := ordem.NewService(repo)

	router := SetupRouter(s)

	// Resposta quando deu certo
	body := gin.H{
		"message": "Processado sem erros.",
	}

	// Insert Correto
	os := &entity.OrdemServico{
		NrCPF:     "12345678912",
		Descricao: "Teste",
		Contato:   "7744",
	}

	b, _ := json.Marshal(os)
	send := bytes.NewBuffer(b)

	// Performa Post
	w := performRequest(router, "POST", "/ordemservico", send)

	// Valida se retornou 201
	assert.Equal(t, http.StatusCreated, w.Code)

	// Converte resposta para um MAP
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Recupera a tag message e verifica se existe
	value, exists := response["message"]

	// Validações
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["message"], value)

	// Insert Incorreto
	os2 := &entity.OrdemServico{
		NrCPF: "12345678912",
	}

	b, _ = json.Marshal(os2)
	send = bytes.NewBuffer(b)

	// Performa Post
	w = performRequest(router, "POST", "/ordemservico", send)

	// Valida se retornou erro
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Converte resposta para um MAP
	var response2 map[string]string
	err = json.Unmarshal([]byte(w.Body.String()), &response2)

	// Verifica se existe o erro
	_, exists = response2["error"]

	// Validações
	assert.Nil(t, err)
	assert.True(t, exists)
}

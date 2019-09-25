package entity

type OrdemServico struct {
	NrCPF     string `json:"nrCPF" valid:"required"`
	Descricao string `json:"damage" valid:"required"`
	Contato   string `json:"contato" valid:"required"`
}

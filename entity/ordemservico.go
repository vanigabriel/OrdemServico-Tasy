package entity

type OrdemServico struct {
	NrCPF     string `json:"nrCPF" valide:"required"`
	Descricao string `json:"damage" valide:"required"`
	Contato   string `json:"contato" valide:"required"`
}

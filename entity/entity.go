package entity

type Entity interface {
	InsertOS(os *OrdemServico) error
}

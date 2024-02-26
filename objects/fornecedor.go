package objects

type Fornecedor struct {
	Nome     string
	Servicos []*Servico
	Deletado bool
}

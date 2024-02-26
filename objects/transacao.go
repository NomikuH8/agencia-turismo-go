package objects

type Transacao struct {
	Comprador  Cliente
	Servicos   []Servico
	Fornecedor Fornecedor
}

func (t Transacao) GerarValorCobrado() float32 {
	var valorCobrado float32

	for _, s := range t.Servicos {
		valorCobrado += s.Valor
	}

	return valorCobrado
}

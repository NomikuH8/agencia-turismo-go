package objects

import "time"

type InformacoesPagamento struct {
	DataVencimento time.Time
	Transacao      Transacao
	ValorCobrado   float32
	Fornecedor     Fornecedor
}

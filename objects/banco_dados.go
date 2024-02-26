package objects

type BancoDados struct {
	Clientes              []Cliente
	Fornecedores          []Fornecedor
	InformacoesPagamentos []InformacoesPagamento
	Servicos              []Servico
	Transacoes            []Transacao
}

func (db *BancoDados) PopularBanco() {
	db.popularClientes()
	db.popularServicos()
	db.popularFornecedores()
}

func (db *BancoDados) popularClientes() {
	db.Clientes = append(db.Clientes, Cliente{
		Nome:  "Menna",
		Email: "menna@gmail.com",
	})
	db.Clientes = append(db.Clientes, Cliente{
		Nome:  "Sophia",
		Email: "sophia@gmail.com",
	})
	db.Clientes = append(db.Clientes, Cliente{
		Nome:  "Graziela",
		Email: "grazi@gmail.com",
	})
	db.Clientes = append(db.Clientes, Cliente{
		Nome:  "Richard",
		Email: "ricardo@gmail.com",
	})
	db.Clientes = append(db.Clientes, Cliente{
		Nome:  "Henrique",
		Email: "henrique@gmail.com",
	})
}

func (db *BancoDados) popularServicos() {
	db.Servicos = append(db.Servicos, Servico{
		Nome:                "Passagens aereas",
		ModalidadePagamento: "3 vezes no crédito",
		Valor:               900.0,
	})
	db.Servicos = append(db.Servicos, Servico{
		Nome:                "Hospedagem 3 dias",
		ModalidadePagamento: "A vista",
		Valor:               200.0,
	})
	db.Servicos = append(db.Servicos, Servico{
		Nome:                "Aluguel de carro",
		ModalidadePagamento: "30 dias em 3 vezes",
		Valor:               500.0,
	})
}

func (db *BancoDados) popularFornecedores() {
	// db.Servicos[0] == Passagens aereas
	// db.Servicos[1] == Hospedagem 3 dias
	// db.Servicos[2] == Aluguel de carro

	db.Fornecedores = append(db.Fornecedores, Fornecedor{
		Nome:     "Mercury",
		Servicos: []*Servico{&db.Servicos[2]},
	})

	db.Fornecedores = append(db.Fornecedores, Fornecedor{
		Nome:     "Localiza",
		Servicos: []*Servico{&db.Servicos[2]},
	})

	db.Fornecedores = append(db.Fornecedores, Fornecedor{
		Nome:     "123milhas",
		Servicos: []*Servico{&db.Servicos[0], &db.Servicos[1], &db.Servicos[2]},
	})

	db.Fornecedores = append(db.Fornecedores, Fornecedor{
		Nome:     "Decolar.com",
		Servicos: []*Servico{&db.Servicos[0], &db.Servicos[1]},
	})

	db.Fornecedores = append(db.Fornecedores, Fornecedor{
		Nome:     "Azul",
		Servicos: []*Servico{&db.Servicos[1]},
	})
}

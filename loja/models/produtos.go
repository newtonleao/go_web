package models

import (
	"loja/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	/*produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	} */

	rows, err := db.Query("SELECT * FROM produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	var produtos []Produto
	for rows.Next() {
		var produto Produto
		err := rows.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtos = append(produtos, produto)
	}

	defer rows.Close()
	defer db.Close()

	return produtos
}

func CriaNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}

func DeletaProduto(idProduto int) {
	db := db.ConectaComBancoDeDados()
	deletaOProduto, err := db.Prepare("delete from produtos where id=?")
	if err != nil {
		panic(err.Error())
	}
	deletaOProduto.Exec(idProduto)
	defer db.Close()
}

func ConsultaProduto(idProduto int) Produto {
	var produto Produto
	db := db.ConectaComBancoDeDados()
	rows, err := db.Query("select * from produtos where id=?", idProduto)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {

		err := rows.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			panic(err.Error())
		}
	}

	defer rows.Close()
	defer db.Close()
	return produto
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	atualizaProduto, err := db.Prepare("update produtos set nome = ?, descricao= ?, preco= ?, quantidade= ? where id=?")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	db.Close()
}

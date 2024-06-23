package models

import (
	"github.com/pedrolessa-dev/loja-golang/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscarTodosOsProdutos() []Produto {
	db := db.ConexaoDB()

	selectProdutos, err := db.Query("SELECT * FROM tbl_produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConexaoDB()

	insertProduto, err := db.Prepare("INSERT INTO tbl_produtos VALUES (null, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insertProduto.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletarProduto(idProduto string) {
	db := db.ConexaoDB()

	deleteProduto, err := db.Prepare("DELETE FROM tbl_produtos WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	deleteProduto.Exec(idProduto)
	defer db.Close()
}

func EditarProduto(id string) Produto {
	db := db.ConexaoDB()

	produtoExistente, err := db.Query("SELECT * FROM tbl_produtos WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	produtoAtualizado := Produto{}
	for produtoExistente.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoExistente.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoAtualizado.Id = id
		produtoAtualizado.Nome = nome
		produtoAtualizado.Descricao = descricao
		produtoAtualizado.Preco = preco
		produtoAtualizado.Quantidade = quantidade
	}
	defer db.Close()
	return produtoAtualizado
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConexaoDB()

	updateProduto, err := db.Prepare("UPDATE tbl_produtos SET nome = ?, descricao = ?, preco = ?, quantidade = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	updateProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}

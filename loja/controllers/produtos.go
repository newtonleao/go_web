package controllers

import (
	"html/template"
	"log"
	"loja/models"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

const MENSAGEM301 int = 301

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosOsProdutos()
	tmpl.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat64, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat64, quantidadeConvertidoParaInt)
		http.Redirect(w, r, "/", MENSAGEM301)

	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	idDoProdutoInt, err := strconv.Atoi(idDoProduto)
	if err != nil {
		log.Println("Erro na conversão do ID em int:", err)
	}
	models.DeletaProduto(idDoProdutoInt)
	http.Redirect(w, r, "/", MENSAGEM301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	idDoProdutoInt, err := strconv.Atoi(idDoProduto)
	if err != nil {
		log.Println("Erro na conversão do ID em int:", err)
	}

	produto := models.ConsultaProduto(idDoProdutoInt)

	tmpl.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para Inteiro", err)
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da Quantidade para Inteiro", err)
		}

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do Preço para Float", err)
		}

		models.AtualizaProduto(idInt, nome, descricao, precoFloat, quantidadeInt)
		http.Redirect(w, r, "/", MENSAGEM301)
	}
}

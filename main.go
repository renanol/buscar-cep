package main

import (
	"github.com/go-martini/martini"
	"github.com/renanol/buscaCep/controllers"
)

func main() {

	m := martini.Classic()

	appCtrl := controllers.App{}
	cepCtrl := controllers.Cep{}
	cpfCtrl := controllers.CpfController{}

	m.Get("/busca-cep/:cep", cepCtrl.BuscarCep)
	m.Get("/", appCtrl.Index)

	m.Get("/cpf", cpfCtrl.Index)
	m.Get("/busca-cpf", cpfCtrl.BuscarCpf)

	m.Run()
}

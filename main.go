package main

import (
	"github.com/go-martini/martini"
	"github.com/renanol/buscaCep/controllers"
)

func main() {

	m := martini.Classic()

	appCtrl := controllers.App{}
	cepCtrl := controllers.Cep{}

	m.Get("/busca-cep/:cep", cepCtrl.BuscarCep)
	m.Get("/", appCtrl.Index)

	m.Run()
}

package main

import (
	"./controllers"
	"github.com/go-martini/martini"
)

func main() {

	m := martini.Classic()

	appCtrl := controllers.App{}
	cepCtrl := controllers.Cep{}

	m.Get("/busca-cep/:cep", cepCtrl.BuscarCep)
	m.Get("/", appCtrl.Index)

	m.Run()
}

package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/renanol/buscar-cep/controllers"
)

func main() {

	m := martini.Classic()

	m.Use(render.Renderer())

	appCtrl := controllers.App{}
	cepCtrl := controllers.Cep{}
	cpfCtrl := controllers.CpfController{}

	m.Get("/busca-cep/:cep", cepCtrl.BuscarCep)
	m.Get("/", appCtrl.Index)

	m.Get("/cpf", cpfCtrl.Index)
	m.Get("/busca-cpf", cpfCtrl.BuscarCpf)

	m.Run()
}

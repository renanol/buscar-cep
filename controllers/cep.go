package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-martini/martini"
	"github.com/renanol/buscaCep/models"
	"log"
	"net/http"
)

type Cep struct{}

type Response struct {
	Message string
}

const (
	URL = "http://www.buscacep.correios.com.br/servicos/dnec/consultaLogradouroAction.do?Metodo=listaLogradouro&TipoConsulta=cep&StartRow=1&EndRow=10&CEP="
)

func (cep *Cep) BuscarCep(rw http.ResponseWriter, req *http.Request, params martini.Params) {

	doc, err := goquery.NewDocument(URL + params["cep"])

	if err != nil {
		log.Fatal(err)
	}

	rw.Header().Set("Content-Type", "application/json")

	if doc.Find("table").Size() > 0 {

		doc.Find("table").Eq(2).Each(func(i int, s *goquery.Selection) {
			logradouro := s.Find("td").Eq(0).Text()
			bairro := s.Find("td").Eq(1).Text()
			localidade := s.Find("td").Eq(2).Text()
			uf := s.Find("td").Eq(3).Text()
			jsonCep, _ := json.Marshal(models.Cep{logradouro, bairro, localidade, uf})
			fmt.Fprint(rw, string(jsonCep))
		})

	} else {

		rw.WriteHeader(404)
		jsonResponse, _ := json.Marshal(Response{"Cep n&#227;o encontrado!"})
		fmt.Fprint(rw, string(jsonResponse))
	}

}

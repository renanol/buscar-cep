package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type Response struct {
	Message string
}
type Cep struct {
	Logradouro string
	Bairro     string
	Localidade string
	Uf         string
}

const (
	URL = "http://www.buscacep.correios.com.br/servicos/dnec/consultaLogradouroAction.do?Metodo=listaLogradouro&TipoConsulta=cep&StartRow=1&EndRow=10&CEP="
)

func buscaHandler(rw http.ResponseWriter, req *http.Request) {
	cep := req.URL.Query().Get("cep")

	fmt.Println(cep)

	doc, err := goquery.NewDocument(URL + cep)
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
			jsonCep, _ := json.Marshal(Cep{logradouro, bairro, localidade, uf})
			fmt.Fprint(rw, string(jsonCep))
		})
	} else {

		rw.WriteHeader(404)
		jsonResponse, _ := json.Marshal(Response{"Cep n&#227;o encontrado!"})
		fmt.Fprint(rw, string(jsonResponse))
	}

}

func main() {

	http.HandleFunc("/", buscaHandler)
	http.ListenAndServe(":8080", nil)

}

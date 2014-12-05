package controllers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/martini-contrib/render"
	"log"
)

const (
	URL_CPF = "https://contas.tcu.gov.br/certidao/Web/Certidao/home.faces?formConsultarCertidao=formConsultarCertidao&formConsultarCertidao%3AtxtCpf=067.624.704-08&formConsultarCertidao%3AbtnEmitirCertidao=Emitir+certid%E3o&formConsultarCertidao%3AqtdAcessos=2608308&javax.faces.ViewState=j_id3"
)

type CpfController struct{}

func (c *CpfController) Index(r render.Render) {
	r.HTML(200, "cpf", "Index")
}

func (c *CpfController) BuscarCpf(r render.Render) {

	doc, err := goquery.NewDocument(URL_CPF)

	if err != nil {
		log.Fatal(err)
	}

	goquery.
		fmt.Println(doc.Find("body").Find("p").Find("font").Find("b").Eq(8).Text())
}

package webdev

import (
	"log"
	"os"
	"text/template"
)

func ExecuteTpl() error {

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("Error parsing the template")
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
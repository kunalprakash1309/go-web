package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	// ParseFiles takes file of any extension and give 
	// pointer to container which contain all template (in this case only one template)

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
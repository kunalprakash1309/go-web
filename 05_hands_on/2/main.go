// Create a data structure to pass to a template which:-
//     - contains information about California hotels including Name, Address, City, Zip, Region
//     - region can be: Southern, Central, Northern
//     - can hold an unlimited number of hotels

package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name string
	Address string
	City string
	Zip int
	Region string
}

func main(){
	tpl, err := template.ParseFiles("tpl.gohtml");
	if err != nil {
		log.Fatalln(err)
	}

	hotels := []hotel{
		hotel{
			Name: "sdfkj",
			Address: "sfjlsfj",
			City: "fslfkj",
			Zip: 343,
			Region: "North",
		},
		hotel{
			Name: "fsdkflj",
			Address: "sfjl",
			City: "sflkj",
			Zip: 545,
			Region: "south",
		},
	}

	err = tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
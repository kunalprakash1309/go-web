package main

import (
	"log"
	"os"
	"encoding/csv"
	"time"
	"fmt"
	"strconv"
	"text/template"
)

type Record struct{
	Date time.Time
	Data float64
}

func main(){
	src, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln("error opening csv file", err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("error in reading csv file", err)
	}

	fmt.Println(len(rows))

	records := make([]Record, len(rows))

	for i, v := range rows{
		if i == 0{
			continue
		}

		date, _ := time.Parse("2006-01-02", v[0])
		open, _ := strconv.ParseFloat(v[1], 1)
		records = append(records, Record{
			Date: date,
			Data: open,
		})
	}

	fmt.Println(records)

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, records)
	if err != nil {
		log.Fatalln(err)
	}
	
}
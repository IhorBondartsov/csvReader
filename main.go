package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/IhorBondartsov/csvReader/cfg"
	"github.com/IhorBondartsov/csvReader/client"
	"github.com/IhorBondartsov/csvReader/entity"
	"github.com/IhorBondartsov/csvReader/parsecsv"
)

var (
	path   string
	url    string
	method string
)

func init() {
	flag.StringVar(&path, "path", cfg.Path, "path where keeps data")
	flag.StringVar(&url, "url", cfg.URL, "url where saves data")
	flag.StringVar(&method, "method", cfg.URL, "method which will be call")
}

func main() {
	result := make(chan entity.PersonData)
	reader := parsecsv.NewReader(path, parsecsv.NewParser(), result)
	go reader.StartRead()

	cl, err := client.NewConnection(cfg.URL)
	if err != nil {
		log.Fatal("Cant create connection")
	}

	for i := range result {
		err := cl.SendPersonData(i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

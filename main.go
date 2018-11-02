package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/IhorBondartsov/csvReader/cfg"
	"github.com/IhorBondartsov/csvReader/entity"
	"github.com/IhorBondartsov/csvReader/parsecsv"
	"net/http"
	"strings"
)

var (
	path string
	url string
	method string
)

func init() {
	flag.StringVar(&path, "path", cfg.Path, "path where keeps data")
	flag.StringVar(&url, "url", cfg.URL, "url where saves data")
	flag.StringVar(&method, "method", cfg.URL, "method which will be call")
}

func main() {
	result := make(chan entity.PersonData)
	reader := parsecsv.NewReader(path, parsecsv.NewParser(),result)
	go reader.StartRead()

	for i := range result {
		SendRequest(i)
	}
}

func SendRequest(pd entity.PersonData){
	data, err := json.Marshal(map[string]interface{}{
		"method": cfg.Method,
		"id":     1,
		"params": pd,
	})

	if err != nil {
		fmt.Printf("Marshal: %v", err)
		return
	}

	resp, err := http.Post(url,
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		fmt.Printf("Post: %v", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}

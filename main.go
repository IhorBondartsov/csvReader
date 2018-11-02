package main

import (
	"flag"
	"fmt"
	"github.com/IhorBondartsov/csvReader/cfg"
	"github.com/IhorBondartsov/csvReader/entity"
	"github.com/IhorBondartsov/csvReader/parsecsv"
	"github.com/ybbus/jsonrpc"
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

	rpcClient := jsonrpc.NewClient("http://127.0.0.1:1812/rpc")

	for i := range result {
		res, err := rpcClient.Call("API.Save", i)
		if err != nil {
			fmt.Println(err)
		}
		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
}

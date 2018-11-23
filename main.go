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
	workers int
	bufferSize int
)

func init() {
	flag.StringVar(&path, "path", cfg.FilePath, "path where keeps data")
	flag.StringVar(&url, "url", cfg.URL, "url where saves data")
	flag.IntVar(&workers, "w", cfg.WorkersForParsing, "count workers which will be started for parsing")
	flag.IntVar(&bufferSize, "s", cfg.SizeChannelBufferForParser, "channel buffer size for worker")
}

func main() {
	fmt.Println(greeating)
	result := make(chan entity.PersonData)
	cfgP := parsecsv.ReaderCfg{
		BufferSize: cfg.SizeChannelBufferForParser,
		Path:path,
		Parser:parsecsv.NewParser(),
		Result:result,
		Workers: bufferSize,
	}
	reader, err := parsecsv.NewReader(cfgP)
	if err != nil {
		fmt.Println(err)
		return
	}
	go reader.StartRead()

	cl, err := client.NewConnection(url)
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


var greeating = `
HELLO! MY DEAR FRIEND!!
___________________________________
 ∧__∧
/ . .\
( >ω<)
(っ▄︻▇〓▄︻┻┳═*  
(     )    /\💥
...................................
Start work ...`
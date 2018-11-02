package parsecsv

import (
	"bufio"
	"fmt"
	"github.com/IhorBondartsov/csvReader/entity"
	"os"
	"sync"
)

type Reader interface {
	StartRead()
}

func NewReader(path string, parser Parser, result chan entity.PersonData) Reader {
	return &reader{
		result: result,
		path:   path,
		parser:parser,
	}
}

type reader struct {
	result chan entity.PersonData
	path   string
	parser Parser
}

func (r *reader) StartRead() {
	wg := sync.WaitGroup{}
	file, err := os.Open(r.path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		x := scanner.Text()
		go func() {
			defer wg.Done()
			d, e := r.parser.Parse(x)
			if e != nil {
				fmt.Println(e.Error())
			}
			r.result <- d
		}()
	}
	wg.Wait()

	close(r.result)
	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}


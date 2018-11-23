package parsecsv

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/IhorBondartsov/csvReader/entity"
	"os"
	"sync"
)

type Reader interface {
	StartRead()
}

type ReaderCfg struct {
	Result     chan entity.PersonData
	Path       string
	Parser     Parser
	Workers    int
	BufferSize int
}

func NewReader(cfg ReaderCfg) (Reader, error) {
	if cfg.Workers == 0 {
		fmt.Println("Count workers equal 0, no sence to create struct")
		return nil, errors.New("WORKERS EQUAL 0")
	}
	return &reader{
		result:     cfg.Result,
		path:       cfg.Path,
		parser:     cfg.Parser,
		workerResp: make(chan string, cfg.BufferSize),
	}, nil
}

type reader struct {
	result       chan entity.PersonData
	path         string
	parser       Parser
	countWorkers int
	workerResp   chan string
}

func (r *reader) startWorkers(wg *sync.WaitGroup) {
	for i := 0; i < r.countWorkers; i++ {
		go r.worker(wg)
	}
}

func (r *reader) worker(wg *sync.WaitGroup) {
	for msg := range r.workerResp {
		d, e := r.parser.Parse(msg)
		if e != nil {
			fmt.Println(e.Error())
		}
		r.result <- d
		wg.Done()
	}
}

func (r *reader) StartRead() {
	wg := sync.WaitGroup{}
	r.startWorkers(&wg)
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
		r.workerResp <- x
	}
	wg.Wait()

	close(r.result)
	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}

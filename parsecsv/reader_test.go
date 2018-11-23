package parsecsv

import "testing"

var testPath ="./data/test.csv"

func TestNewReader(t *testing.T) {
	result := make(chan entity.PersonData)
	cfgR := ReaderCfg{
		Workers: 1,
		Result: result,

	}
}
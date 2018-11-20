package parsecsv

import "testing"

// go test -run=parser_test.go -bench=. -benchmem

var testData = []string{
	"1,Kirk,ornare@sedtortor.net,(013890) 37420",
	"2,Cain,volutpat@semmollisdui.com,(016977) 2245",
	"3,Geoffrey,vitae@consectetuermaurisid.co.uk,0800 1111",
	"4,Walter,odio.a.purus@sit.edu,(0161) 328 6656",
	"5,Armand,Cras.vulputate@metusvitae.co.uk,0836 796 0064",
}

func BenchmarkParse(b *testing.B) {
	p := NewParser()
	testCSVStr := "3,Geoffrey,vitae@consectetuermaurisid.co.uk,0800 1111"
	for i := 0; i < b.N; i++ {
		p.Parse(testCSVStr)
	}
}

package parsecsv

import (
	"github.com/IhorBondartsov/csvReader/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

// go test -run=parser_test.go -bench=. -benchmem


// This testData have interconnection with testPerson see resultRelation
var testData = []string{
	"1,Kirk,ornare@sedtortor.net,(013890) 37420",
	"2,Cain,volutpat@semmollisdui.com,(016977) 2245",
	"3,Geoffrey,vitae@consectetuermaurisid.co.uk,0800 1111",
	"4,Walter,odio.a.purus@sit.edu,(0161) 328 6656",
	"5,Armand,Cras.vulputate@metusvitae.co.uk,0836 796 0064",
}

// This testPerson have interconnection with testData see resultRelation
var testPerson = []entity.PersonData{
	entity.PersonData{
		Id:1,
		Name: "Kirk",
		Email: "ornare@sedtortor.net",
		MobileNumber: "01389037420",
	},
	entity.PersonData{
		Id:2,
		Name: "Cain",
		Email: "volutpat@semmollisdui.com",
		MobileNumber: "0169772245",
	},
	entity.PersonData{
		Id:3,
		Name: "Geoffrey",
		Email: "vitae@consectetuermaurisid.co.uk",
		MobileNumber: "08001111",
	},
	entity.PersonData{
		Id:4,
		Name: "Walter",
		Email: "odio.a.purus@sit.edu",
		MobileNumber: "01613286656",
	},
	entity.PersonData{
		Id:5,
		Name: "Armand",
		Email: "Cras.vulputate@metusvitae.co.uk",
		MobileNumber: "08367960064",
	},
}

var resultRelation = map[string]entity.PersonData{
	testData[0]: testPerson[0],
	testData[1]: testPerson[1],
	testData[2]: testPerson[2],
	testData[3]: testPerson[3],
	testData[4]: testPerson[4],
}

func BenchmarkParse(b *testing.B) {
	p := NewParser()
	testCSVStr := "3,Geoffrey,vitae@consectetuermaurisid.co.uk,0800 1111"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(testCSVStr)
	}
}


func TestParse(t *testing.T) {
	a := assert.New(t)
	p := NewParser()

	for k, v := range resultRelation{
		res, err := p.Parse(k)
		a.NoError(err)
		a.Equal(v, res)
	}
}

func TestParseWrongData(t *testing.T) {
	a := assert.New(t)
	p := NewParser()
	wrongData := map[string] error{
		"WRONG ID, Geoffrey,vitae@consectetuermaurisid.co.uk,0800 1111": entity.ErrInvalidID.Error(),
		"1, Geoffrey,vitae@consectetuermaurisid.co.uk,ddddddd": entity.ErrInvalidNumber.Error(),
		"1, Geoffrey": entity.ErrDataMalformed.Error(),
	}

	for k, v := range wrongData{
		res, err := p.Parse(k)
		a.Error(v, err)
		a.Equal(entity.PersonData{}, res)
	}
}
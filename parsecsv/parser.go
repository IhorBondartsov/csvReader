package parsecsv

import (
	"github.com/IhorBondartsov/csvReader/entity"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"unicode"
)

const countFieldInStruct = 4

type Parser interface {
	Parse(str string) (data entity.PersonData, err error)
}

func NewParser() Parser{
	return &MyCustomParser{}
}

type MyCustomParser struct{}


func (r *MyCustomParser) Parse(str string) (data entity.PersonData, err error) {

	arr := strings.Split(str, ",")
	if len(arr) != countFieldInStruct  {
		err =  errors.New("Invalid CSV")
		return
	}
	data.Id, err = strconv.Atoi(arr[0])
	if err != nil{
		return
	}
	data.Email = arr[1]
	data.Name = arr[2]

	data.MobileNumber = strings.TrimFunc(arr[3], func(r rune) bool {
		return !unicode.IsNumber(r)
	})

	return
}


package parsecsv

import (
	"fmt"
	"github.com/IhorBondartsov/csvReader/entity"
	"regexp"
	"strconv"
	"strings"
)

const countFieldInStruct = 4

type Parser interface {
	Parse(str string) (data entity.PersonData, err error)
}

func NewParser() Parser {
	return &MyCustomParser{
		reqForNumber:regexp.MustCompile("[0-9]+"),
	}
}

type MyCustomParser struct{
	reqForNumber *regexp.Regexp
}

func (r *MyCustomParser) Parse(str string) (data entity.PersonData, err error) {
	arr := strings.Split(str, ",")
	if len(arr) != countFieldInStruct {
		return entity.PersonData{}, entity.ErrDataMalformed.Error()
	}
	data.Id, err = strconv.Atoi(arr[0])
	if err != nil {
		return entity.PersonData{}, entity.ErrInvalidID.Error()
	}
	data.Name = arr[1]
	data.Email = arr[2]

	result := r.reqForNumber.FindAllString(arr[3], -1)
	fmt.Println(len(result))
	if len(result)  == 0 {
		return entity.PersonData{}, entity.ErrInvalidNumber.Error()
	}
	data.MobileNumber = strings.Join(result[:],"")
	return
}

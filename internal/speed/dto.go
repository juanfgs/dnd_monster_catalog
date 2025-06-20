package speed 

import (
	"reflect"
	"strings"
	"strconv"
	"log"
)
type SpeedDTO struct {
	Walk string 
	Swim string 
	Fly string 
	Burrow string
	Climb string
}

type ValueUnit struct {
	Value int64
	Unit string
}

func BuildModels(s SpeedDTO) []Speed {
	speeds := make([]Speed, 0, 0)
	fields := reflect.VisibleFields(reflect.TypeOf(s))
	for _, field := range(fields) {
		value := reflect.ValueOf(s).FieldByName(field.Name).String()
		if field.Name != "" && value != "" {
			valueUnit, err := extractValue(value)
			if err != nil {
				log.Fatal(err)
			}
			speeds = append(speeds, Speed{
				Type: field.Name,
				Value: valueUnit.Value, 
				Unit: valueUnit.Unit, 
			})
			log.Println(speeds)
		}
	}
	return speeds
}

func extractValue(s string) (*ValueUnit, error) {
	fields := strings.Split(s, " ")
	value, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		return nil, err
	}
	switch fields[1] {
	default: 
		return &ValueUnit{
			Value: value ,
			Unit: "feet",
		}, nil
	}
	
	
} 

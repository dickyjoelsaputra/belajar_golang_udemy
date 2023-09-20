package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10" min:"5"`
}

func IsValid(data interface{})bool{
	t := reflect.TypeOf(data)
	
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}

	return true
}

func main() {
	sample := Sample{"sample"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")

	fmt.Println(structField)
	fmt.Println(required)
}
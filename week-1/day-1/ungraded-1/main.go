package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type User struct {
	Name  string `required:"true" minLen:"3" maxLen:"20"`
	Crime string `required:"true" minLen:"3" maxLen:"100"`
	Age   int    `required:"true" min:"21" max:"80"`
	Email string `required:"true" minLen:"10" maxLen:"100"`
}

func main() {
	vilain1 := User{
		Name:  "Joni",
		Crime: "Pencurian dana BTS 4G",
		Age:   60,
		Email: "joni@mail.com",
	}

	if err := vilain1.ValidateStruct(); err != nil {
		fmt.Println(err)
	}
}

func (u User) ValidateStruct() error {
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		if field.Tag.Get("required") == "true" && value == "" {
			return fmt.Errorf("%s is required", field.Name)
		}

		if field.Type.String() == "string" {
			value := value.(string)

			emailRegex, _ := regexp.Compile(`^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$`)
			if field.Name == "Email" && !emailRegex.MatchString(value) {
				return fmt.Errorf("%v does not comply with email format", value)
			}

			valueLen := len(value)
			minLen, _ := strconv.Atoi(field.Tag.Get("minLen")) 
			maxLen, _ := strconv.Atoi(field.Tag.Get("maxLen")) 

			if valueLen < minLen || valueLen > maxLen {
				return fmt.Errorf("%v's length exceeded minimum / maximum length", value)
			}

		} else if field.Type.String() == "int" {
			value := value.(int)
			
			min, _ := strconv.Atoi(field.Tag.Get("min")) 
			max, _ := strconv.Atoi(field.Tag.Get("max")) 

			if value < min || value > max {
				return fmt.Errorf("%v exceeded minimum / maximum value", value)
			}
		}
	}
	return nil
}

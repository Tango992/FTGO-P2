package main

import (
	"fmt"
	"reflect"
)

type Users struct {
	Name      string `required:"true"`
	Username  string `required:"true"`
	Password  string `required:"true"`
	Level     string `required:"true"`
	CreatedAt string `required:"true"`
	UpdateAt  string `required:"true"`
}

func ValidateStruct(s interface{}) error {
	t := reflect.TypeOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(s).Field(i).Interface()
			if value == "" {
				return fmt.Errorf("%s is required", field.Name)
			}
		}
	}
	return nil
}

func main() {
	// var number float64 = 23.42
	// var reflectValue = reflect.ValueOf(number)

	// fmt.Println("tipe variable :", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Float64 {
	// 	fmt.Println("nilai float :", reflectValue.Float())
	// } else if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nili int :", reflectValue.Int())
	// }

	newUser := Users{
		Name: "Rizky",
		Username: "rizky01",
		Password: "123456",
		Level: "asd",
		CreatedAt: "2021-01-01",
		UpdateAt: "2021-01-01",
	}

	// userValue := reflect.ValueOf(newUser)
	// fmt.Println(userValue)

	// userType := reflect.TypeOf(newUser)
	// fmt.Println(userType)

	// userField := userType.NumField()
	// fmt.Println(userField)

	// fmt.Println(userType.Field(2))
	// fmt.Println(userType.Field(2).Name)

	if err := ValidateStruct(newUser); err != nil {
		fmt.Println(err.Error())
	}
}

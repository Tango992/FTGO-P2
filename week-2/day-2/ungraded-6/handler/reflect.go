package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"slices"
	"ungraded-6/entity"
)

func ValidateStruct(u any) *entity.Response {
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := 0;i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if field.Tag.Get("required") == "true" && value.IsZero() {
			return &entity.Response{
				Code: http.StatusBadRequest,
				Message: fmt.Sprintf("%v is required", field.Name),
				Data: nil,
			}
		} else if field.Tag.Get("required") == "" {
			continue
		}

		if field.Type.String() == "string" {
			value := value.Interface().(string)

			emailRegex, _ := regexp.Compile(`^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$`)
			if field.Name == "Email" && !emailRegex.MatchString(value) {
				return &entity.Response{
					Code: http.StatusBadRequest,
					Message: "Invalid email format",
					Data: nil,
				}
			}

			roles := []string{"admin", "superadmin"}
			if field.Name == "Role" && !slices.Contains(roles, value){
				return &entity.Response{
					Code: http.StatusBadRequest,
					Message: "Invalid role",
					Data: nil,
				}
			}

			valueLen := len(value)
			minLenStr := field.Tag.Get("minLen")
			maxLenStr := field.Tag.Get("maxLen")

			if minLenStr == "" {
				minLenStr = "0"
			}

			if maxLenStr == "" {
				maxLenStr = "255"
			}

			minLen, _ := strconv.Atoi(minLenStr) 
			maxLen, _ := strconv.Atoi(maxLenStr) 

			if valueLen < minLen || valueLen > maxLen {
				return &entity.Response{
					Code: http.StatusBadRequest,
					Message: fmt.Sprintf("%v's length exceeded minimum / maximum length", field.Name),
					Data: nil,
				}
			}

		} else if field.Type.String() == "int" {
			value := value.Interface().(int)

			min, _ := strconv.Atoi(field.Tag.Get("min")) 
			max, _ := strconv.Atoi(field.Tag.Get("max")) 

			if value < min || value > max {
				return &entity.Response{
					Code: http.StatusBadRequest,
					Message: fmt.Sprintf("%v's length exceeded minimum / maximum value", field.Name),
					Data: nil,
				}
			}
		}
	}
	return nil
}
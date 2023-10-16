package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type Hero struct {
	Name     string `required:"true" minLen:"3" maxLen:"16"`
	Email    string `required:"true" minLen:"3" maxLen:"50"`
	Strength int    `required:"true" min:"100" max:"250"`
	Health   int    `required:"true" min:"100" max:"250"`
}

func IsRequired(hero interface{}) string {
	t := reflect.TypeOf(hero)
	v := reflect.ValueOf(hero)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {

			if v.Field(i).Interface() == "" {
				return fmt.Sprintf("%s is required", field.Name)
			}
		}

		// string minLen and maxLen Check
		minLen := field.Tag.Get("minLen")
		maxLen := field.Tag.Get("maxLen")

		if minLen != "" && maxLen != "" {
			// Convertion from string to int
			min, err := strconv.Atoi(minLen)
			if err != nil {
				return fmt.Sprintf("'%v', not valid min length.", min)
			}
			max, err := strconv.Atoi(maxLen)
			if err != nil {
				return fmt.Sprintf("'%v', not valid min length.", max)
			}

			val := v.Field(i).Interface()

			if len(val.(string)) < min {
				return fmt.Sprintf("%v length must be greater than %v", val, min)
			}

			if len(val.(string)) > max {
				return fmt.Sprintf("%v length must be less than %v", val, max)
			}
		}

		// integer min and max value check
		min := field.Tag.Get("min")
		max := field.Tag.Get("max")

		if min != "" && max != "" {
			// Convertion from string to int
			min, err := strconv.Atoi(min)
			if err != nil {
				return fmt.Sprintf("'%v', not valid min length.", min)
			}
			max, err := strconv.Atoi(max)
			if err != nil {
				return fmt.Sprintf("'%v', not valid min length.", max)
			}

			val := v.Field(i).Interface()

			if val.(int) < min {
				return fmt.Sprintf("Value of %v must be greater than %v", val, min)
			}

			if val.(int) > max {
				return fmt.Sprintf("Value of %v must be less than %v", val, max)
			}
		}
	}

	// Email Check
	emailVal := v.FieldByName("Email")
	result := EmailValidation(emailVal.String())
	if !result {
		return fmt.Sprintf("%v -> Email not valid.", emailVal)
	}

	return fmt.Sprintf("%s, You are ready to go.", v.FieldByName("Name"))
}

func EmailValidation(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)
}

func main() {
	hulk := Hero{
		Name:     "Hulk",
		Email:    "hulk@avengers.com",
		Strength: 200,
		Health:   140,
	}

	fmt.Println(IsRequired(hulk))
}

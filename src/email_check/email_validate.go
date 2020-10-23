package main

import (
	"fmt"
	"regexp"
	"strings"
)

func email_validate(email string) {
	if strings.Count(email, "@") == 1 {
		email_substrings := strings.Split(email, "@")	
		if len(email_substrings[0]) > 0 && len(email_substrings[0]) < 128 && len(email_substrings[1]) > 3 && len(email_substrings[1]) < 256 {
			email_pattern := `[_0-9a-z\-]*\.?[_0-9a-z\-]*("[_0-9a-z!:\-]*\.?[_0-9a-z!:\-]*")*[_0-9a-z\-]*\.?[_0-9a-z\-]*@(((([_0-9a-z]{1,})|([_0-9a-z]{1}[_0-9a-z-]*[_0-9a-z]{1}))\.){1,}[0-9a-z]{2,})`
			matched, _ := regexp.MatchString(email_pattern, email)
			if matched {
				fmt.Printf("email корректный")
			} else {
				fmt.Printf("email не соответсвует шаблону")
			}
		} else {
			fmt.Printf("длина email не соответсвует условиям")
		}
	} else {
		fmt.Printf("email содержит более одного символа @")
	}
}		
		
		

func main() {
	var email string
	fmt.Printf("Insert email for validate: ")
	fmt.Scanf("%s\n", &email)
	email_validate(email)
}
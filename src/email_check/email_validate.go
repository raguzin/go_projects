package main

import (
	"fmt"
	"regexp"
	"strings"
)

func email_validate(email string) {
	if strings.Count(email, "@") == 1 {	// e-mail состоит из имени и доменной части, эти части разделяются символом "@"
		email_substrings := strings.Split(email, "@")	
		domain_substring := email_substrings[1]
		domain_substring_pattern := `^(\w+(-*)\w+\.)+\w{2,}$`
		if len(domain_substring) > 3 && len(domain_substring) < 256 {
			matched, err := regexp.MatchString(domain_substring_pattern, domain_substring)
			if matched == false {
				fmt.Printf("доменная часть не соответствует условию")
			}
			//fmt.Println(matched)
			//fmt.Println(err)
		} else {
			fmt.Printf("доменная часть короче 3 или длинее 256 символов")
		}

		name_substring := email_substrings[0]
		name_substring_pattern := `\w*(-*)(\.)?"{2}`
		if len(name_substring) > 128 {
			matched, err := regexp.MatchString(name_substring_pattern, name_substring)
			//if matched == false {
			//	fmt.Printf("доменная часть не соответствует условию")
			//}
			fmt.Println(matched)
			fmt.Println(err)
		} else {
			fmt.Printf("именная часть длинее 128 символов")
		}
	}
}
		
		
		

func main() {
	var email string
	fmt.Printf("Insert email for validate: ")
	fmt.Scanf("%s\n", &email)
	email_validate(email)

//	fmt.Printf("You email - %s\n", email)

}
package main

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"os"
	"testos/another"
	"net/url"
)

//К приложению из практической части предыдущего урока добавьте возможность читать данные из файлов.
//Конфигурация может быть задана в форматах yaml или json. Также по желанию вы можете добавить и другие форматы.
//Помимо чтения конфигурации приложение также должно валидировать её - например,
//все URL’ы должны соответствовать ожидаемым форматам.

func main() {

	data, err := os.ReadFile("data.json")

	if err != nil {
		panic(err)
	}

	var dt another.Config

	if err = json.Unmarshal(data, &dt); err != nil {
		panic(err)
	}

	//	fmt.Printf("%5s valid: %t\n", dt.DBEmail, valid(dt.DBEmail))
	if valid(dt.DBEmail) != true {
		fmt.Println("Некорректный email в профиле")
		panic(err)
	}
	u, err := url.ParseRequestURI(dt.DBUrl)	
  	  if err != nil {
		fmt.Println("Url не авктуален и отcутствуют данные", u)
   panic(err)
}

	fmt.Printf("%+v\n", dt)
	//

	s := another.ReadConfig()

	fmt.Println(s)
}

func valid(DBEmail string) bool {
	_, err := mail.ParseAddress(DBEmail)
	return err == nil
}

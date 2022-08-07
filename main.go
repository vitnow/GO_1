package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testos/another"
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
	fmt.Printf("%+v\n", dt)
	//

	s := another.ReadConfig()

	fmt.Println(s)
}

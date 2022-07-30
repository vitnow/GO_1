package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testos/another"
)

type Config struct {
	dbConfig
	poolConfig
}

type dbConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

type poolConfig struct {
	PoolMaxConns    int
	PoolMinConns    int
	GoroutinesCount int
}

//К приложению из практической части предыдущего урока добавьте возможность читать данные из файлов.
//Конфигурация может быть задана в форматах yaml или json. Также по желанию вы можете добавить и другие форматы.
//Помимо чтения конфигурации приложение также должно валидировать её - например,
//все URL’ы должны соответствовать ожидаемым форматам.

func main() {
	var dt Config
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &dt); err != nil {
		panic(err)
	}
	fmt.Println(dt)

	s := another.ReadConfig()

	fmt.Println(s)
}

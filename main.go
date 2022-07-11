package main

// 1. С какими интерфейсами мы уже сталкивались в предыдущих уроках?
// Обратите внимание на уроки, в которых мы читали из стандартного ввода и писали в стандартный вывод.

// 2. Посмотрите примеры кода в своём портфолио.
// Везде ли ошибки обрабатываются грамотно? Хотите ли вы переписать какие-либо функции?

// 3. Тест

import "fmt"

type Lenguager interface {
	Hello() string
}
type Russian struct{}

func (Russian) Hello() string {
	return "Привет"
}

type English struct{}

func (English) Hello() string {
	return "Hello"
}

type Franch struct{}

func (Franch) Hello() string {
	return "Bonjour"
}

func sayHello(hello Lenguager) {
	fmt.Println(hello.Hello())
}

func main() {
	russianLang := Russian{}
	englishLang := English{}
	franchLang := Franch{}

	sayHello(russianLang)
	sayHello(englishLang)
	sayHello(franchLang)
}

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

	//sayHello(russianLang)
	//sayHello(englishLang)
	//sayHello(franchLang)

	sl := make([]Lenguager, 0, 3)
	sl = append(sl, russianLang, englishLang, franchLang)

	for _, Lenguager := range sl {
		fmt.Println(Lenguager.Hello())
	}

}
# GO_1

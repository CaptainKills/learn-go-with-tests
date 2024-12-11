package helloworld

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	prefixEnglish = "Hello "
	prefixSpanish = "Hola "
	prefixFrench  = "Bonjour "

	postfix = "!"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(lang) + name + postfix
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = prefixSpanish
	case french:
		prefix = prefixFrench
	default:
		prefix = prefixEnglish
	}

	return
}

func main() {
	fmt.Println(Hello("Danick", "English"))
}

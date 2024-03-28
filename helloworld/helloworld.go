package helloworld

const (
	spanish    = "Spanish"
	french     = "French"
	portuguese = "Portuguese"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix

	case french:
		prefix = frenchHelloPrefix

	case portuguese:
		prefix = portugueseHelloPrefix

	default:
		prefix = englishHelloPrefix
	}

	return
}

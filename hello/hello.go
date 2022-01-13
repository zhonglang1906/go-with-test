package hello

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	//if language == spanish {
	//	return spanishHelloPrefix + name
	//}
	//
	//if language == french {
	//	return frenchHelloPrefix + name
	//}
	////使用switch代替多个if
	//prefix := englishHelloPrefix
	//switch language {
	//case french:
	//	prefix = frenchHelloPrefix
	//case spanish:
	//	prefix = spanishHelloPrefix
	//}
	//将长代码封装到另一个函数。

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

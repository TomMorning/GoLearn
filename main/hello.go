package main

import "fmt"

const englishPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Hddfla, "

// 版本一
//func Hello(name string) string {
//	return englishPrefix + name
//}

// 版本二
//func Hello(name string, language string)  string{
//	if name == "" {
//		name = "World"
//	}
//	prefix := englishPrefix
//	switch language {
//	case french:
//		prefix = frenchHelloPrefix
//	case spanish:
//		prefix = spanishHelloPrefix
//	}
//	return prefix + name
//}

// 版本三
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", french))
}

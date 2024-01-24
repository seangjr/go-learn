package hello

import "fmt"

var prefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
		case "Spanish":
			prefix = "Hola, "
		case "French":
			prefix = "Bonjour, "
		case "Chinese":
			prefix = "你好, "
		case "Japanese":
			prefix = "こんにちは, "
		case "Korean":
			prefix = "안녕하세요, "
		case "Vietnamese":
			prefix = "Xin chào, "
	}
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case "Spanish":
			prefix = "Hola, "
		case "French":
			prefix = "Bonjour, "
		case "Chinese":
			prefix = "你好, "
		case "Japanese":
			prefix = "こんにちは, "
		case "Korean":
			prefix = "안녕하세요, "
		case "Vietnamese":
			prefix = "Xin chào, "
		default:
			prefix = "Hello, "
	}
	return
}

func main()  {
	fmt.Println(Hello("Sean", "English"))
}
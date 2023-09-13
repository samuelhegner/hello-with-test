package main

import "fmt"

func main() {
	fmt.Println(Hello("", "french"))
}

type greetingPair struct {
	greeting string
	world    string
}

var localGreetings = map[string]greetingPair{
	"spanish": {
		greeting: "Hola ",
		world:    "Mundo",
	},
	"english": {
		greeting: "Hello ",
		world:    "World",
	},
	"french": {
		greeting: "Bonjour ",
		world:    "le monde",
	},
}

func Hello(name string, lang string) string {
	if _, ok := localGreetings[lang]; !ok {
		lang = "english"
	}

	g := localGreetings[lang].greeting
	w := localGreetings[lang].world

	if name == "" {
		return g + w + "!"
	}

	return g + name + "!"
}

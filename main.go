package main

import (
	"fmt"
)

func main() {
	dictionaryApi := Init("API")

	buffer, _ := dictionaryApi.Lookup("Привет", "ru-en", "", FLAGS_FAMILY)
	fmt.Println(buffer)
}

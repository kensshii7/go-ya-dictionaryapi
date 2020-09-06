package main

import (
	"fmt"
)

func main() {
	dictionaryApi := Init("dict.1.1.20200906T162446Z.618894a4672590dc.b0914b4496701dec37a2a9338404638b49f56026")

	buffer, _ := dictionaryApi.Lookup("Привет", "ru-en", "", FLAGS_FAMILY)
	fmt.Println(buffer)
}

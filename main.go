package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	caser := cases.Title(language.English)
	fmt.Println(caser.String("live streaming service"))
}

package main

import(
	"fmt"
	"strings"
)

func main() {
	mainStrings()
}

func mainStrings() {
	s := "hello string"

	fmt.Println("Contains 'string': ", strings.Contains(s, "string"))
	fmt.Println("Contains 'hell': ", strings.Contains(s, "hell"))
	fmt.Println("Count 'l': ", strings.Count(s, "l"))
	fmt.Println("Index 'string': ", strings.Index(s, "string"))
	fmt.Println("Repeat 'ha': ", strings.Repeat("ha", 5))
	fmt.Println("Replace 'o' with 'x': ", strings.Replace(s, "o", "x", 1))
	fmt.Println("ReplaceAll 'o' with 'x': ", strings.ReplaceAll(s, "o", "x"))
	fmt.Println("Split 'o': ", strings.Split(s, "o"))
	fmt.Println("Title 'hello string': ", strings.Title(s))
	fmt.Println("ToLower 'HELLO STRING': ", strings.ToLower(s))
	fmt.Println("ToUpper 'hello string': ", strings.ToUpper(s))
	fmt.Println("HasPrefix 'hello': ", strings.HasPrefix(s, "hello"))
	fmt.Println("HasSuffix 'string': ", strings.HasSuffix(s, "string"))

	s2 := "      password          "
	fmt.Println("TrimSpace '      password          ': ", "'"+strings.TrimSpace(s2)+"'")
}
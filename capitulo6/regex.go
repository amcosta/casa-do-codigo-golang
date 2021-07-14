package capitulo6

import (
	"fmt"
	"regexp"
	"strings"
)

func FuncaoAnonima() {
	expr := regexp.MustCompile("\\b\\w")
	transformadora := func(s string) string {
		fmt.Println(s)
		return strings.ToUpper(s)
	}

	texto := "antonio carlos jobim"
	fmt.Println(transformadora(texto))
	fmt.Println(expr.ReplaceAllStringFunc(texto, transformadora))
}

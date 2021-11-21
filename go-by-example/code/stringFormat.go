package gobyex

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func StringFormatting() {

	p := point{1, 2}
	fmt.Printf("s1 %v\n", p)
	fmt.Printf("s2 %+v\n", p)
	fmt.Printf("s3 %#v\n", p)

	fmt.Printf("type %T\n", p)
	fmt.Printf("bool %t\n", true)

	fmt.Printf("char int %c\n", 122)

	fmt.Printf("f1 %e\n", 123400000.0)
	fmt.Printf("f2 %E\n", 123400000.0)

	fmt.Printf("str1 %s\n", "\"hello\"")
	fmt.Printf("str2 %q\n", "\"hello\"")

	fmt.Printf("pointer1 %p\n", &p)

	fmt.Printf("width1 |%6d|%6d|\n", 12, 123)
	fmt.Printf("width2 |%6.2f|%6.2f|\n", 1.2, 23.34)
	fmt.Printf("width3 left justify |%-6.2f|%-6.2f|\n", 1.2, 23.34)

	fmt.Printf("str width |%6s|%6s|\n", "foo", "b")

	s := fmt.Sprintf("sprinf :%d %f %s", 1, 2.3, "ad")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "error :%s\n", "os")

}

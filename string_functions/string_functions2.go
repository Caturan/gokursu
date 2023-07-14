package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	name := "Cevdet"
	fmt.Println(s.HasPrefix(name, "Cevd"))
	fmt.Println(s.HasSuffix(name, "det"))

	letters := []string{"H", "a", "z", "a", "l"}
	name1 := (s.Join(letters, ""))
	fmt.Println(name1)
	letters1 := []string{"L", "O", "V", "E"}
	result := (s.Join(letters1, "*"))
	fmt.Println(result)

	result1 := (s.ReplaceAll(result, "*", ""))
	fmt.Println(result1)
	fmt.Println(s.Replace(result, "*", "-", 1))

	fmt.Println(s.Split(name, ""))

	fmt.Println(s.Repeat(result1, 3))
}

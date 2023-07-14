package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	name := "Aytac"
	t := s.Count(name, "t")
	fmt.Println("Count of t in string of 'Aytac' :", t)
	T := s.Count(name, "T")
	fmt.Println("Count of T in strings of 'Aytac' :", T)

	y := s.Contains(name, "y")
	fmt.Println("Is it that s contains 'y' ? ", y)
	Y := s.Contains(name, "Y")
	fmt.Println("Is it that s contains 'Y' ? ", Y)

	// Golang is a key sensitive language

	name1 := "Cemile"
	e := s.Index(name1, "e")
	fmt.Println("First index of 'e' :", e)
	b := s.Index(name1, "b")
	fmt.Println(b) // If char not in this string program return -1
	if b == -1 {
		fmt.Println("Do not have b")
	} else {
		fmt.Println("Have b")
	}

	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name1))
}

package main

import "github.com/wild-mouse/go-playground/pkg/rest"

func main()  {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}

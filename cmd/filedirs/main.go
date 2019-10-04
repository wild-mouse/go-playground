package main

import (
	"github.com/wild-mouse/go-playground/pkg/filedirs"
)

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CapitalizerExample(); err != nil  {
		panic(err)
	}
}

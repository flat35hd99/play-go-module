package main

import (
	"fmt"
	"os"
	pgm "play-go-module"
)

func main() {
	u := pgm.User{
		ID:   1,
		Name: "Flat",
	}
	fmt.Fprintf(os.Stdout, "Hello, %s", u.Name)
}

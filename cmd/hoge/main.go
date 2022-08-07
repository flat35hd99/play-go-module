package main

import (
	"fmt"
	"os"
	xxx "play-go-module"
)

func main() {
	u := xxx.User{
		ID:   1,
		Name: "Flat",
	}
	fmt.Fprintf(os.Stdout, "Hello, %s", u.Name)
}

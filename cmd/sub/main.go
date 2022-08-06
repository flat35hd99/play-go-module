package main

import (
	"fmt"
	"os"
)

func main() {
	u := User{
		ID:   1,
		Name: "Flat",
	}
	fmt.Fprintf(os.Stdout, "Hello, %s", u.Name)
}

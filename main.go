package main

import (
	"fmt"

	"earthly/hello"

	"github.com/google/uuid"
)

func main() {
	uuid.New()
	fmt.Println(hello.Hello("rei"))
}

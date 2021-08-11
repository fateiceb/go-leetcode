package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	sex  uint8
}

func main() {
	fmt.Printf("%#v", reflect.Type)
}

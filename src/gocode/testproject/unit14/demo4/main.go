package main

import (
	"fmt"
	"reflect"
)

func main() {

	num := 10

	reflectSetInt(&num, 20)

	fmt.Println(num)

}

func reflectSetInt(x interface{}, value int) {
	(reflect.ValueOf(x)).Elem().SetInt(int64(value))
}

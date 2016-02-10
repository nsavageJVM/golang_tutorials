package main

import (
	"fmt"
	"container/list"

	"github.com/nsavageJVM/rest"

)

type info struct {
	result string
}

// return both result and error  https://gobyexample.com/multiple-return-values
func infoRepo (a int) (string, error) {
	return "this is the intro package", nil
}

func main() {
	// variable declaration and initialisation https://golang.org/ref/spec#Variable_declarations
	var introMssg string = "Hello from Go toolchain"
	// string formatting
	fmt.Println("Go reports: %+v\n ", introMssg)
	// avoid this
	fmt.Println("Go reports: " + introMssg)
	// short variable declaration and initialisation  https://golang.org/ref/spec#Short_variable_declarations work with functions and errors
	mssg, err := infoRepo( 2 )
	// handling errors  https://golang.org/pkg/errors/
	if err != nil {
		fmt.Println("Go reports error condition: %+v\n ",err)
	}
	fmt.Println("Go reports error condition: %+v\n ", mssg)

	// go arrays  https://tour.golang.org/moretypes/6  https://golang.org/ref/spec#Numeric_types
	xs := []float64 {98, 93, 77, 82, 83}
	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total )

	// working with struct, Structs provide a flexible way of defining composite types
	//https://gobyexample.com/structs
	var s = info{}
	// generate a pointer to s    https://golang.org/ref/spec#Address_operators
	sp := &s
	// now we point to the memory location of s we can set its fields
	// https://tour.golang.org/moretypes/1
	sp.result = "set a struct pointer value"
	fmt.Println("Go reports: %+v\n ", sp.result )
	l := list.New()
	l.PushBack(4)
	l.PushFront(1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	rest.TestScope(" testing the scope")


}


package main

/*
#cgo LDFLAGS: -L ./ -ladd
#include "Add.h"
*/
import "C"
import "fmt"

func main(){
	fmt.Println("this is test")
	fmt.Println(C.add(C.int(1),C.int(2)));
}

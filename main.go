package main

import "fmt"

/*
#include "test/test.h"
*/
import "C"

func main() {
	fmt.Println(C.x)
}

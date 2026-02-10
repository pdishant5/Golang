package main

import (
	"fmt"
	"os"
)

func Abc() {
	fmt.Println("Hello!")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		fmt.Println("Hello!!")
		os.Exit(1)
	}()
	panic("panic occured!")
}

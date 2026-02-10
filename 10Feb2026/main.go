package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func function() (x int) {
	// a := 5
	// x := &a
	x = 5
	defer func() {
		x++
	}()
	return
}

func foo2(a int) error {
	if a == 1 {
		return errors.New("Something went wrong!")
	}
	return nil
}

func foo() error {
	if err := foo2(2); err != nil {
		return err
	}
	return nil
}

func main() {
	// fmt.Println(function())

	// err := foo()
	// fmt.Println("Error:", err)

	// // defer fmt.Println("Hello!")
	// // select {} // fatal error

	// slice1 := make([]int, 3)
	// slice2 := new([]int)

	// fmt.Println("Length:", len(slice1), "; Capacity:", cap(slice1))
	// fmt.Println(slice1)
	// fmt.Println("Length:", len(*slice2), "; Capacity:", cap(*slice2))
	// fmt.Println(slice2)
	// fmt.Println(*slice2 == nil)

	// *slice2 = append(*slice2, 1, 2, 3)

	// fmt.Println("Length:", len(*slice2), "; Capacity:", cap(*slice2))
	// fmt.Println(slice2)

	// x := new(int)
	// fmt.Println(x)

	// m1 := make(map[int]int)
	// m2 := new(map[int]int)

	// fmt.Println("Size:", len(m1))
	// fmt.Println(m1 == nil)
	// fmt.Println("Size:", len(*m2))
	// fmt.Println(*m2 == nil)

	// var mu sync.Mutex

	// go func1(mu)

	// go func2(mu)

	// select {}
	// time.Sleep(6 * time.Second)
	// fmt.Println("Inside  main")

	// slice := make([]int, 5)
	// slice[0], slice[1], slice[2] = 1, 2, 3
	// s := slice

	// fmt.Println("Length:", len(slice), "; Capacity:", cap(slice))
	// fmt.Println(slice)

	// slice = slice[:2:2]
	// fmt.Println("Length:", len(slice), "; Capacity:", cap(slice))
	// fmt.Println("Length:", len(s), "; Capacity:", cap(s))
	// fmt.Println(slice)
	// slice = slice[:2]
	// fmt.Println(s)

	slice := make([]int, 5, 10)
	fmt.Println("Length:", len(slice), "; Capacity:", cap(slice))
	fmt.Println(slice)
	slice = slice[:5:5]
	fmt.Println("Length:", len(slice), "; Capacity:", cap(slice))
	fmt.Println(slice)

	slice2 := slice[:3]
	fmt.Println("Length:", len(slice2), "; Capacity:", cap(slice2))
	fmt.Println(slice2)

	slice2 = slice2[:5]
	fmt.Println("Length:", len(slice2), "; Capacity:", cap(slice2))
	fmt.Println(slice2)

}

func func1(mu *sync.Mutex) {
	fmt.Println("Mutex 1 starting!")

	mu.Lock()
	fmt.Println("Mutex locked!")
	// go func2(mu)
	time.Sleep(5 * time.Second)
	fmt.Println("Waken up!")

	mu.Unlock()
	fmt.Println("Mutex unlocked!")

}

func func2(mu *sync.Mutex) {
	// time.Sleep(time.Second)

	fmt.Println("Hello!")
	mu.Unlock()
	fmt.Println("Mutex 2 unlocked!")
}

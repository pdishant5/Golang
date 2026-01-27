package main

import "fmt"

func main() {
	fmt.Println(add(5, 10))                       // Output: 15
	fmt.Println(add(5.5, 10.3))                   // Output: 15.8
	fmt.Println(add("Hello, ", "Go"))             // Output: Hello, Go
	fmt.Println(add(5, "Go"))                     // Output: <nil>
	fmt.Println("\n", addWithGenerics(5, 10))     // Output: 15
	fmt.Println(addWithGenerics(5.5, 10.3))       // Output: 15.8
	fmt.Println(addWithGenerics("Hello, ", "Go")) // Output: Hello, Go
	// addWithGenerics(5, "Go") 				  // compile-time error
}

func addWithGenerics[T int | float64 | string](a, b T) T {
	return a + b
}

func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aStr, aIsStr := a.(string)
	bStr, bIsStr := b.(string)

	if aIsStr && bIsStr {
		return aStr + bStr
	}

	return nil
}

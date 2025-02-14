package main

import "fmt"

// Generic Function
func printType[T any](value T) {
	fmt.Println("Value:", value)
}

// Generic Struct
type Box[T any] struct {
	value T
}

// Generic Method
func (b *Box[T]) GetValue() T {
	return b.value
}

// Generic Constraint: Only Numeric Types
type Number interface {
	int | float64
}

func add[T Number](a, b T) T {
	return a + b
}

func main() {
	// Using Generic Function
	printType(42)
	printType("Hello, Generics!")

	// Using Generic Struct
	intBox := Box[int]{value: 100}
	strBox := Box[string]{value: "GoLang"}
	fmt.Println("Box Value (int):", intBox.GetValue())
	fmt.Println("Box Value (string):", strBox.GetValue())

	// Using Generic Function with Constraint
	fmt.Println("Sum of Integers:", add(10, 20))
	fmt.Println("Sum of Floats:", add(3.5, 2.5))
}

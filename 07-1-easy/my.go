package main

//https://semaphoreci.com/blog/testify-go
import "fmt"

// MyFunction returns the sum of two integers.
func MyFunction(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("both arguments must be non-negative")
	}
	return a + b, nil
}

type Calc struct{}

// CalculateArea
func (cla Calc) CalculateArea(width int, height int) int {
	return width * height
}

func main() {
	fmt.Println(MyFunction(23, 34))
	fmt.Println()
}

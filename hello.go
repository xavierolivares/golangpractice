package main

import (
	"fmt"
	// "errors"
	// "math"
)

type person struct {
	name string
	age int
}

func main() {

	i := 7
	inc(&i)

	fmt.Println(&i)

	// p := person{name: "Xavier", age: 6}
	// fmt.Println(p.age)
	// m := make(map[string]string)
	// m["a"] = "alpha"
	// m["b"] = "beta"

	// for key, value := range m {
	// 	fmt.Println("index:", key, "value:", value)
	// }
	// arr := []string{"a", "b", "c"}

	// for index, value := range arr {
	// 	fmt.Println("index:", index, "value:", value)
	// }
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// // fmt.Println("Hello, World")
	// vertices := make(map[string]int)

	// vertices["triangle"] =2
	// vertices["square"]= 3
	// vertices["dodecagon"] = 12
	// // a := []int{5,4,3,2,1}
	// // a[2] = 8
	// // a = append(a,13)
	// delete(vertices,"square")
	// fmt.Println(vertices)

// 	result := sum(2,3)
// 	fmt.Println(result)
// }

// func sum (x int, y int) int {
// 	return x + y

// 	result, err := sqrt(-16)

// 	if err != nil {
// 		fmt.Println(result)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("Undefined for negative numbers")
// 	}
// 	return math.Sqrt(x), nil
}

func inc(x *int) {
	*x++
}
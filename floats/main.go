package main

import (
	"fmt"
	"math"
)

func main() {
	var numFloat32 float32

	numFloat32 = math.MaxFloat32
	fmt.Printf("numFloat32: %v\n", numFloat32)

	var numFloat64 float64

	numFloat64 = math.MaxFloat64
	fmt.Printf("numFloat32: %v\n", numFloat64)
}

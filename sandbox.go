package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// pre-initialized
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// numeric constants
const (
	BigNumber   = 1 << 100        // 1 is shifted 1-bit left 100 times, creating a huge number
	SmallNumber = BigNumber >> 99 // A huge number is shifted 1-bit right 99 times, creating 1<<1 or simply put 2
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value: %v\n", z, z)

	//zero-initialized
	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Printf("%v, %v, %v, %v", i, f, b, s)
	fmt.Println()

	//type conversion
	var (
		int1, int2         = 3, 4
		float      float64 = math.Sqrt(float64(int1*int1 + int2*int2))
		unsigned   uint    = uint(float)
	)

	fmt.Println(int1, int2, unsigned)
	fmt.Println()

	//type inference
	v := 42.042 + 0.5i //change this value
	fmt.Printf("Variable v is of type %T\n", v)
	fmt.Println()

	//constants
	const Pi = 3.14
	const World = "世界"
	const Truth = true

	fmt.Println("Hello", World)
	fmt.Println(Pi, "zda")
	fmt.Println(Truth, "ili ne", Truth)

	//numeric constants declaration referenced at the top
	fmt.Println(needInt(SmallNumber))
	//fmt.Println(needInt(BigNumber)) -- NumericOverflow
	fmt.Println(needFloat(SmallNumber))
	fmt.Println(needFloat(BigNumber))
}

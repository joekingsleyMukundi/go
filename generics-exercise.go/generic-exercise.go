package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Distance int32
type Velocity float64

type Number interface {
	constraints.Float | constraints.Integer
}

func clamp[T Number](value, min, max T) T {
	if value > max {
		return max
	} else if value < min {
		return min
	} else {
		return value
	}
}

func testClampInt8() {
	var (
		min int8 = -10
		max int8 = 10
	)
	if clamp(-30, min, max) != min {
		panic("clamp failed for int8")
	}
}

func testClampUint32() {
	var (
		min uint32 = 0
		max uint32 = 10
	)
	if clamp(20, min, max) != max {
		panic("clamp failed for uint32")
	}
}

func testClampFloat32() {
	var (
		min float32 = 5.5
		max float32 = 9.9
	)
	if clamp(0, min, max) != min {
		panic("clamp failed for float32")
	}
}

func testClampDistance() {
	var (
		min Distance = 0
		max Distance = 100
	)
	if clamp(Distance(99), min, max) != 99 {
		panic("clamp failed for Distance")
	}
}

func testClampVelocity() {
	var (
		min Velocity = 0.0
		max Velocity = 99.9
	)
	if clamp(Velocity(100), min, max) != max {
		panic("clamp failed for Velocity")
	}
}

func main() {
	testClampInt8()
	testClampUint32()
	testClampFloat32()
	testClampDistance()
	testClampVelocity()
	fmt.Println("Exercise complete!")
}

package hand

import (
	"fmt"
	"log"
	"math"
)

func Safe(err error) bool {
	if err != nil {
		fmt.Errorf("core.Safe")
		panic(err.Error())
		return false
	}
	return true
}

func FiniteFloat(value float64) float64  {
	if (math.IsNaN(value)) {
		fmt.Errorf("FiniteFloat in NaN")
		return -1
	}
	if math.IsInf(value, 1) {
		return 9007199254740991
	}
	if math.IsInf(value, -1) {
		return -9007199254740991
	}
	return value
}

func SafeFloat64(value float64, err error) float64 {
	if err != nil {
		log.Println("core.SafeFloat64")
		log.Fatal(err.Error())
		return FiniteFloat(value)
	}
	return FiniteFloat(value)
}
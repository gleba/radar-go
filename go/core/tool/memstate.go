package tool

import (
	"fmt"
	"runtime"
)

func MemState() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats

	fmt.Print(":::::::::::::::::  ")
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Println("")
}

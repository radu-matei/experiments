package main

import (
	"C"
	"fmt"

	"github.com/engineerd/wasm-to-oci/pkg/oci"
)

//export Pull
func Pull(ref, outFile string) int64 {
	err := oci.Pull(ref, outFile)
	if err != nil {
		fmt.Printf("cannot pull module: %v", err)
		return 1
	}

	return 0
}

//export Sum
func Sum(s []int, c chan int) (chan int, error) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c

	return c, nil
}

func main() {}

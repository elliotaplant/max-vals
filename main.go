package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Range uint8:  0 - %d\n", ^uint8(0))                       // 0 - 255
	fmt.Printf("Range uint16: 0 - %d\n", ^uint16(0))                      // 0 - 65535
	fmt.Printf("Range uint32: 0 - %d\n", ^uint32(0))                      // 0 - 4294967295
	fmt.Printf("Range uint64: 0 - %d\n", ^uint64(0))                      // 0 - 18446744073709551615
	fmt.Printf("Range int8:   %d - %d\n", ^int8(0)<<(8-1), 1<<(8-1)-1)    // -128 - 127
	fmt.Printf("Range int16:  %d - %d\n", ^int16(0)<<(16-1), 1<<(16-1)-1) // -32768 - 32767
	fmt.Printf("Range int32:  %d - %d\n", ^int32(0)<<(32-1), 1<<(32-1)-1) // -2147483648 - 2147483647
	fmt.Printf("Range int64:  %d - %d\n", ^int64(0)<<(64-1), 1<<(64-1)-1) // -9223372036854775808 - 9223372036854775807
}

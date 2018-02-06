// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	emptySlice := make([]byte, 99999)
// 	filePointer, _ := os.Open(os.Args[1])

// 	filePointer.Read(emptySlice)
// 	fmt.Println(string(emptySlice[:]))
// }

package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
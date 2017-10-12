package main

import "os"
import "bytes"

func main() {
	var buffer bytes.Buffer
	const BFSIZE = 1204
	// allocate buffer of size BFSIZE
	for i := 0; i < BFSIZE; i++ {
		buffer.WriteString("y\n")
	}
	for true {
		os.Stdout.Write(buffer.Bytes())
	}
}
